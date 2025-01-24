package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/entities"
	_itemManagingException "github.com/supakornn/game-shop/pkg/itemManaging/exception"
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
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

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) error {
	if err := r.db.Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("editing item failed: %v", err.Error())
		return &_itemManagingException.ItemEditing{}
	}

	return nil
}
