package repository

import (
	"github.com/supakornn/game-shop/entities"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDs []uint64) ([]*entities.Item, error)
	PurchaseHistory(tx *gorm.DB, purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
}
