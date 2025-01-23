package server

import (
	_itemManagingController "github.com/supakornn/game-shop/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/supakornn/game-shop/pkg/itemManaging/repository"
	_itemManagingService "github.com/supakornn/game-shop/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("/", itemManagingController.Creating)
}
