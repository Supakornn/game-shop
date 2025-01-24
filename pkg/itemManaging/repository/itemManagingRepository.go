package repository

import (
	"github.com/supakornn/game-shop/entities"
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) error
	Archiving(itemID uint64) error
}
