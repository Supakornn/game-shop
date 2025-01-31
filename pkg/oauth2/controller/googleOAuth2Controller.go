package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/avast/retry-go"
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/config"
	"github.com/supakornn/game-shop/pkg/custom"
	_oauth2Exception "github.com/supakornn/game-shop/pkg/oauth2/exception"
	_oauth2Service "github.com/supakornn/game-shop/pkg/oauth2/service"
	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"

	_adminModel "github.com/supakornn/game-shop/pkg/admin/model"
	_oauth2Model "github.com/supakornn/game-shop/pkg/oauth2/model"
	_playerModel "github.com/supakornn/game-shop/pkg/player/model"
)

type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.Oauth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookie  = "act"
	refreshTokenCookie = "rft"
	stateCookieName    = "state"

	letter = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(oauth2Service _oauth2Service.Oauth2Service, oauth2Conf *config.OAuth2, logger echo.Logger) Oauth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		logger:        logger,
	}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.EndPoints.AuthUrl,
			TokenURL:      oauth2Conf.EndPoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.EndPoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.EndPoints.AuthUrl,
			TokenURL:      oauth2Conf.EndPoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.EndPoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	state := c.randomState()
	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	state := c.randomState()
	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3)); err != nil {
		c.logger.Errorf("Failed to get token: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	token, err := playerGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to get token: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	client := playerGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	playerCreatingReq := &_playerModel.PlayerCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingReq); err != nil {
		c.logger.Errorf("Failed to create player account: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, accessTokenCookie, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookie, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3)); err != nil {
		c.logger.Errorf("Failed to get token: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	token, err := adminGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to get token: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	client := adminGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	adminCreatingReq := &_adminModel.AdminCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(adminCreatingReq); err != nil {
		c.logger.Errorf("Failed to create admin account: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, accessTokenCookie, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookie, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})

}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
	accessToken, err := pctx.Cookie(accessTokenCookie)
	if err != nil {
		c.logger.Errorf("Failed to get access token cookie: %v", err)
		return custom.Error(pctx, http.StatusBadRequest, &_oauth2Exception.Logout{})
	}

	if err := c.revokeToken(accessToken.Value); err != nil {
		c.logger.Errorf("Failed to revoke token: %v", err)
		return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.Logout{})
	}

	c.removeSameSiteCookie(pctx, accessTokenCookie)
	c.removeSameSiteCookie(pctx, refreshTokenCookie)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LogoutResponse{Message: "Logout Success"})
}

func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) callbackValidating(pctx echo.Context) error {
	state := pctx.QueryParam("state")

	stateCookie, err := pctx.Cookie(stateCookieName)
	if err != nil {
		c.logger.Errorf("Failed to get state cookie: %v", err)
		return &_oauth2Exception.Unauthorized{}
	}

	if state != stateCookie.Value {
		c.logger.Errorf("State is not match")
		return &_oauth2Exception.Unauthorized{}
	}

	c.removeCookie(pctx, stateCookieName)

	return nil
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.UserInfo, error) {
	resp, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Errorf("Error getting user info: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	userInfoInBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorf("Error reading user info: %v", err)
		return nil, err
	}

	userInfo := new(_oauth2Model.UserInfo)
	if err := json.Unmarshal(userInfoInBytes, userInfo); err != nil {
		c.logger.Errorf("Error unmarshalling user info: %v", err)
		return nil, err
	}

	return userInfo, nil
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeUrl := fmt.Sprintf("%s?token=%s", c.oauth2Conf.RevokeUrl, accessToken)

	resp, err := http.Post(revokeUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		c.logger.Errorf("Error revoking token: %v", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}
