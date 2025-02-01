package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/databases"
)

type inventoryRepositoryImpl struct {
	db     *databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db *databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db: db, logger: logger}
}
