package service

import (
	_PlayerCoinRepository "github.com/supakornn/game-shop/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _PlayerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository _PlayerCoinRepository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{
		playerCoinRepository: playerCoinRepository,
	}
}
