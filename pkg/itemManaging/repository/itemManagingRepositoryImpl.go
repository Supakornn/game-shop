package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/entities"
	_itemManagingException "github.com/supakornn/game-shop/pkg/itemManaging/exception"
	"gorm.io/gorm"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("creating item failed: %v", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}
