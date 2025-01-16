package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakorn/game-shop/entities"
	_itemShopException "github.com/supakorn/game-shop/pkg/itemShop/exception"
	_itemShopModel "github.com/supakorn/game-shop/pkg/itemShop/model"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{})

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	if err := query.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list item: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
