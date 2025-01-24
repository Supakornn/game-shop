package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/entities"
	_itemManagingException "github.com/supakornn/game-shop/pkg/itemManaging/exception"
	_itemManagingModel "github.com/supakornn/game-shop/pkg/itemManaging/model"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("creating item failed: %v", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) error {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("editing item failed: %v", err.Error())
		return &_itemManagingException.ItemEditing{ItemID: itemID}
	}

	return nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("archiving item failed: %v", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
