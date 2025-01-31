package server

import (
	_adminRepository "github.com/supakornn/game-shop/pkg/admin/repository"
	_oauth2Controller "github.com/supakornn/game-shop/pkg/oauth2/controller"
	_oauth2Service "github.com/supakornn/game-shop/pkg/oauth2/service"
	_playerRepository "github.com/supakornn/game-shop/pkg/player/repository"
)

func (s *echoServer) initOAuth2Router() {
	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _playerRepository.NewPlayerRepository(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepository(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(oauth2Service, s.conf.Oauth2, s.app.Logger)

	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/player/login/callback", oauth2Controller.PlayerLoginCallback)
	router.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	router.DELETE("/logout", oauth2Controller.Logout)
}
