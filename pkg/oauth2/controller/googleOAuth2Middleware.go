package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/pkg/custom"
	_oauth2Exception "github.com/supakornn/game-shop/pkg/oauth2/exception"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()
	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.playerTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsPlayerExist(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()
	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.adminTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsAdminExist(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookie, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookie, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookie, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookie, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(accessTokenCookie)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	refreshToken, err := pctx.Cookie(refreshTokenCookie)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil

}
