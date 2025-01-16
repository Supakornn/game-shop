package repository

import "github.com/supakorn/game-shop/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
