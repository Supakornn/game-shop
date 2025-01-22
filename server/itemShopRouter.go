package server

import (
	_itemShopController "github.com/supakornn/game-shop/pkg/itemShop/controller"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
	_itemShopService "github.com/supakornn/game-shop/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopController(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
