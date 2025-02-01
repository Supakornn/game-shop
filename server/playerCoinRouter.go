package server

import (
	_playerCoinController "github.com/supakornn/game-shop/pkg/playerCoin/controller"
	_PlayerCoinRepository "github.com/supakornn/game-shop/pkg/playerCoin/repository"
	_playerCoinService "github.com/supakornn/game-shop/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _PlayerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
}
