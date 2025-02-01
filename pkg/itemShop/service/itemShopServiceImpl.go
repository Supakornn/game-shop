package service

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/entities"
	_inventoryRepository "github.com/supakornn/game-shop/pkg/inventory/repository"
	_itemShopException "github.com/supakornn/game-shop/pkg/itemShop/exception"
	_itemShopModel "github.com/supakornn/game-shop/pkg/itemShop/model"
	_itemShopRepository "github.com/supakornn/game-shop/pkg/itemShop/repository"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
	_PlayerCoinRepository "github.com/supakornn/game-shop/pkg/playerCoin/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _PlayerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerCoinRepository _PlayerCoinRepository.PlayerCoinRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	logger echo.Logger,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository:   itemShopRepository,
		playerCoinRepository: playerCoinRepository,
		inventoryRepository:  inventoryRepository,
		logger:               logger,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCounting, itemFilter.Size)

	return s.toItemResultRes(itemList, itemFilter.Page, totalPage), nil
}

func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		s.logger.Error("item not found")
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), buyingReq.Quantity)
	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()

	purchaseRecording, err := s.itemShopRepository.PurchaseHistory(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		IsBuying:        true,
		Quantity:        buyingReq.Quantity,
	})
	if err != nil {
		s.logger.Errorf("purchase recording failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("purchase recording success %v", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		s.logger.Errorf("deducting coin failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("deducting coin success %v", playerCoin.Amount)

	inventoryEntities, err := s.inventoryRepository.Filling(tx, buyingReq.PlayerID, buyingReq.ItemID, int(buyingReq.Quantity))
	if err != nil {
		s.logger.Errorf("filling inventory failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("filling inventory success %v", len(inventoryEntities))

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(sellingReq.ItemID)
	if err != nil {
		s.logger.Error("item not found")
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), sellingReq.Quantity)
	totalPrice = totalPrice / 2

	if err := s.playerItemChecking(sellingReq.PlayerID, sellingReq.ItemID, sellingReq.Quantity); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()

	purchaseRecording, err := s.itemShopRepository.PurchaseHistory(tx, &entities.PurchaseHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          sellingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		IsBuying:        false,
		Quantity:        sellingReq.Quantity,
	})
	if err != nil {
		s.logger.Errorf("purchase recording failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("purchase recording success %v", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		s.logger.Errorf("deducting coin failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("deducting coin success %v", playerCoin.Amount)

	if err := s.inventoryRepository.Removing(tx, sellingReq.PlayerID, sellingReq.ItemID, int(sellingReq.Quantity)); err != nil {
		s.logger.Errorf("removing inventory failed: %v", err)
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("removed inventory success %v", sellingReq.Quantity)

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItem int64, Size int64) int64 {
	totalPage := totalItem / Size
	if totalItem%Size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultRes(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Error("coin not enough")
		return &_itemShopException.CoinNotEnough{}
	}

	return nil
}

func (s *itemShopServiceImpl) playerItemChecking(playerID string, itemID uint64, qty uint) error {
	itemCounting := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(itemCounting) < int(qty) {
		s.logger.Error("item not enough")
		return &_itemShopException.ItemQuantityNotEnough{ItemID: itemID}
	}

	return nil
}
