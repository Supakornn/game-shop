package controller

import (
	_inventoryService "github.com/supakornn/game-shop/pkg/inventory/service"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
}

func NewInventoryController(inventoryService _inventoryService.InventoryService) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
	}
}
