package service

import (
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
}
