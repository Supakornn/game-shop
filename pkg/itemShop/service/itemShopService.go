package service

import (
	_itemShopModel "github.com/supakorn/game-shop/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
