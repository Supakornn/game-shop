package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/supakornn/game-shop/entities"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepositoryMock struct {
	mock.Mock
}

func (m *PlayerCoinRepositoryMock) CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	args := m.Called(tx, playerCoinEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *PlayerCoinRepositoryMock) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	args := m.Called(playerID)
	return args.Get(0).(*_playerCoinModel.PlayerCoinShowing), args.Error(1)
}
