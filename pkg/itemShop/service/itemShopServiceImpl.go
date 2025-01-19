package service

import (
	"github.com/supakorn/game-shop/entities"
	_itemShopModel "github.com/supakorn/game-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/supakorn/game-shop/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository: itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCounting, itemFilter.Size)

	return s.toItemResultRes(itemList, itemFilter.Page, totalPage), nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItem int64, Size int64) int64 {
	totalPage := totalItem / Size
	if totalItem%Size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultRes(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
