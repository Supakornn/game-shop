package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/pkg/custom"
	_inventoryService "github.com/supakornn/game-shop/pkg/inventory/service"
	"github.com/supakornn/game-shop/pkg/validation"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryControllerImpl(inventoryService _inventoryService.InventoryService, logger echo.Logger) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		c.logger.Errorf("Error getting playerID: %v", err)
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)
}
