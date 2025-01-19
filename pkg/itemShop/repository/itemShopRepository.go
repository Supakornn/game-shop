package repository

import (
	"github.com/supakorn/game-shop/entities"
	_itemShopModel "github.com/supakorn/game-shop/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
}
