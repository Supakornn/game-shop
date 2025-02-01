package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/entities"
	"gorm.io/gorm"

	_inventoryException "github.com/supakornn/game-shop/pkg/inventory/exception"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db: db, logger: logger}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("filling inventory failed: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities, err := r.findPlayerItemByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true
		if err := conn.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("removing player item failed: %s", err.Error())
			return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(
		&entities.Inventory{},
	).Where("player_id = ? AND item_id = ? AND is_deleted = ?",
		playerID, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Errorf("counting player item failed: %s", err.Error())
		return 0
	}

	return count
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? AND is_deleted = ?",
		playerID, false,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("listing inventory failed: %s", err.Error())
		return nil, &_inventoryException.PlayerItemFinding{PlayerID: playerID}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) findPlayerItemByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? AND item_id = ? AND is_deleted = ?",
		playerID, itemID, false,
	).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("finding player item by id failed: %s", err.Error())
		return nil, &_inventoryException.PlayerItemRemoving{ItemID: itemID}
	}

	return inventoryEntities, nil
}
