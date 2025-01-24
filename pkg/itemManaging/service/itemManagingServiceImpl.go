package service

import (
	"github.com/supakornn/game-shop/entities"
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
	_itemManagingRepository "github.com/supakornn/game-shop/pkg/itemManaging/repository"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository, itemShopRepository _itemShopRepository.ItemShopRepository) ItemManagingService {
	return &itemManagingServiceImpl{
		itemManagingRepository: itemManagingRepository,
		itemShopRepository:     itemShopRepository,
	}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Price:       itemCreatingReq.Price,
		Picture:     itemCreatingReq.Picture,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	if err := s.itemManagingRepository.Editing(itemID, itemEditingReq); err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepository.FindByID(itemID)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}
