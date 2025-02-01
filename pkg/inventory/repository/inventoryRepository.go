package repository

import (
	"github.com/supakornn/game-shop/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
