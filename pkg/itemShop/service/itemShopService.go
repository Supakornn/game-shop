package service

import (
	_itemShopModel "github.com/supakorn/game-shop/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item, error)
}
