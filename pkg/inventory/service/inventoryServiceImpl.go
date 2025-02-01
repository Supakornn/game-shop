package service

import (
	_inventoryRepository "github.com/supakornn/game-shop/pkg/inventory/repository"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(inventoryRepository _inventoryRepository.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		itemShopRepository:  itemShopRepository,
	}
}
