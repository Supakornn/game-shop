package service

import (
	_inventoryModel "github.com/supakornn/game-shop/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
