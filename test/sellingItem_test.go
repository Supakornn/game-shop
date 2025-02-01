package test

import (
	"errors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/supakornn/game-shop/entities"
	_inventoryRepository "github.com/supakornn/game-shop/pkg/inventory/repository"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
	_itemShopService "github.com/supakornn/game-shop/pkg/itemShop/service"
	_playerCoinRepository "github.com/supakornn/game-shop/pkg/playerCoin/repository"
	"gorm.io/gorm"
)

func TestItemShopService_Selling(t *testing.T) {
	mockItemShopRepo := new(_itemShopRepository.ItemShopRepositoryMock)
	mockPlayerCoinRepo := new(_playerCoinRepository.PlayerCoinRepositoryMock)
	mockInventoryRepo := new(_inventoryRepository.InventoryRepositoryMock)
	logger := echo.New().Logger

	service := _itemShopService.NewItemShopServiceImpl(
		mockItemShopRepo,
		mockPlayerCoinRepo,
		mockInventoryRepo,
		logger,
	)

	tests := []struct {
		name          string
		sellingReq    *_itemShopModel.SellingReq
		mockSetup     func()
		expectedError error
	}{
		{
			name: "should sell item successfully",
			sellingReq: &_itemShopModel.SellingReq{
				PlayerID: "player1",
				ItemID:   1,
				Quantity: 1,
			},
			mockSetup: func() {

				mockItemShopRepo.On("FindByID", uint64(1)).Return(&entities.Item{
					ID:          1,
					Name:        "Test Item",
					Description: "Test Description",
					Price:       100,
					Picture:     "test.jpg",
				}, nil)

				mockInventoryRepo.On("PlayerItemCounting", "player1", uint64(1)).Return(int64(2))

				mockItemShopRepo.On("TransactionBegin").Return(&gorm.DB{})
				mockItemShopRepo.On("TransactionCommit", mock.Anything).Return(nil)

				mockItemShopRepo.On("PurchaseHistory", mock.Anything, mock.Anything).Return(&entities.PurchaseHistory{
					ID: 1,
				}, nil)

				mockPlayerCoinRepo.On("CoinAdding", mock.Anything, mock.Anything).Return(&entities.PlayerCoin{
					ID:       1,
					PlayerID: "player1",
					Amount:   50,
				}, nil)

				mockInventoryRepo.On("Removing", mock.Anything, "player1", uint64(1), 1).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "should fail when not enough items to sell",
			sellingReq: &_itemShopModel.SellingReq{
				PlayerID: "player1",
				ItemID:   1,
				Quantity: 5,
			},
			mockSetup: func() {
				mockItemShopRepo.On("FindByID", uint64(1)).Return(&entities.Item{
					ID:    1,
					Price: 100,
				}, nil)

				mockInventoryRepo.On("PlayerItemCounting", "player1", uint64(1)).Return(int64(2))
			},
			expectedError: errors.New("not enough items"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.Selling(tt.sellingReq)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
