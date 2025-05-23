package service

import (
	"github.com/supakornn/game-shop/entities"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
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

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoinResult, err := s.playerCoinRepository.CoinAdding(nil, playerCoinEntity)
	if err != nil {
		return nil, err
	}

	return playerCoinResult.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}
	return playerCoinShowing
}
