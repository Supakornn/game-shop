package service

import (
	"github.com/supakornn/game-shop/entities"
	_inventoryModel "github.com/supakornn/game-shop/pkg/inventory/model"
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

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventoryEntities, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantity := s.getUniqueItemWithQuantity(inventoryEntities)

	return s.inventoryListingResult(uniqueItemWithQuantity), nil

}

func (s *inventoryServiceImpl) getUniqueItemWithQuantity(inventoryEntities []*entities.Inventory) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) inventoryListingResult(uniqueItemWithQuantity []_inventoryModel.ItemQuantityCounting) []*_inventoryModel.Inventory {
	uniqueItemIDList := s.getItemByID(uniqueItemWithQuantity)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	result := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantity)

	for _, itemEntity := range itemEntities {
		result = append(result, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}

	return result
}

func (s *inventoryServiceImpl) getItemByID(uniqueItemWithQuantity []_inventoryModel.ItemQuantityCounting) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantity {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(uniqueItemWithQuantity []_inventoryModel.ItemQuantityCounting) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantity {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
