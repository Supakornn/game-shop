package service

import (
	"github.com/supakornn/game-shop/entities"
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
	_itemManagingRepository "github.com/supakornn/game-shop/pkg/itemManaging/repository"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository: itemManagingRepository}
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
