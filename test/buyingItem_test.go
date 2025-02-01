package test

import (
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/supakornn/game-shop/entities"
	_inventoryRepository "github.com/supakornn/game-shop/pkg/inventory/repository"
	_itemShopException "github.com/supakornn/game-shop/pkg/itemShop/exception"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
	_itemShopService "github.com/supakornn/game-shop/pkg/itemShop/service"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
	_playerCoinRepository "github.com/supakornn/game-shop/pkg/playerCoin/repository"
	"gorm.io/gorm"
)

func TestBuyingItem_PlayerCoinConversion(t *testing.T) {
	now := time.Now()
	playerCoin := &entities.PlayerCoin{
		ID:        123,
		PlayerID:  "player1",
		Amount:    1000,
		CreatedAt: now,
	}

	model := playerCoin.ToPlayerCoinModel()

	assert.Equal(t, playerCoin.ID, model.ID)
	assert.Equal(t, playerCoin.PlayerID, model.PlayerID)
	assert.Equal(t, playerCoin.Amount, model.Amount)
	assert.True(t, model.CreatedAt.Equal(playerCoin.CreatedAt))
}

func TestItemShopService_Buying(t *testing.T) {
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
		buyingReq     *_itemShopModel.BuyingReq
		mockSetup     func()
		expectedError error
	}{
		{
			name: "should buy item successfully",
			buyingReq: &_itemShopModel.BuyingReq{
				PlayerID: "player1",
				ItemID:   1,
				Quantity: 2,
			},
			mockSetup: func() {
				mockItemShopRepo.On("FindByID", uint64(1)).Return(&entities.Item{
					ID:          1,
					Name:        "Test Item",
					Description: "Test Description",
					Price:       100,
					Picture:     "test.jpg",
				}, nil)

				mockPlayerCoinRepo.On("Showing", "player1").Return(&_playerCoinModel.PlayerCoinShowing{
					PlayerID: "player1",
					Coin:     1000,
				}, nil)

				mockItemShopRepo.On("TransactionBegin").Return(&gorm.DB{})
				mockItemShopRepo.On("TransactionCommit", mock.Anything).Return(nil)

				mockItemShopRepo.On("PurchaseHistory", mock.Anything, mock.Anything).Return(&entities.PurchaseHistory{
					ID: 1,
				}, nil)

				mockPlayerCoinRepo.On("CoinAdding", mock.Anything, mock.Anything).Return(&entities.PlayerCoin{
					ID:       1,
					PlayerID: "player1",
					Amount:   800,
				}, nil)

				mockInventoryRepo.On("Filling", mock.Anything, "player1", uint64(1), 2).Return([]*entities.Inventory{}, nil)
			},
			expectedError: nil,
		},
		{
			name: "should fail when not enough coins",
			buyingReq: &_itemShopModel.BuyingReq{
				PlayerID: "player1",
				ItemID:   1,
				Quantity: 10,
			},
			mockSetup: func() {
				mockItemShopRepo.On("FindByID", uint64(1)).Return(&entities.Item{
					ID:          1,
					Name:        "Test Item",
					Description: "Test Description",
					Price:       10000,
					Picture:     "test.jpg",
				}, nil)

				mockPlayerCoinRepo.On("Showing", "player1").Return(&_playerCoinModel.PlayerCoinShowing{
					PlayerID: "player1",
					Coin:     500,
				}, nil)

				mockItemShopRepo.On("TransactionBegin").Return(&gorm.DB{}).Once()

				mockInventoryRepo.On("Filling", mock.Anything, "player1", uint64(1), 10).Return([]*entities.Inventory{}, &_itemShopException.CoinNotEnough{})

				mockItemShopRepo.On("TransactionRollback", mock.Anything).Return(nil).Once()
			},
			expectedError: &_itemShopException.CoinNotEnough{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.Buying(tt.buyingReq)

			if tt.expectedError != nil {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.expectedError.Error(), err.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
