package server

import (
	_inventoryController "github.com/supakornn/game-shop/pkg/inventory/controller"
	_inventoryRepository "github.com/supakornn/game-shop/pkg/inventory/repository"
	_inventoryService "github.com/supakornn/game-shop/pkg/inventory/service"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	inventoryService := _inventoryService.NewInventoryServiceImpl(inventoryRepository, itemShopRepository)
	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing)
}
