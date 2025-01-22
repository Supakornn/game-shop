package repository

import (
	"github.com/supakornn/game-shop/entities"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
}
