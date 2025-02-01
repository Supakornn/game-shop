package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/supakornn/game-shop/entities"
	"gorm.io/gorm"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	args := m.Called(tx, playerID, itemID, qty)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {
	args := m.Called(tx, playerID, itemID, limit)
	return args.Error(0)
}

func (m *InventoryRepositoryMock) PlayerItemCounting(playerID string, itemID uint64) int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(int64)
}

func (m *InventoryRepositoryMock) Listing(playerID string) ([]*entities.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}
