package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/entities"
	_playerCoinException "github.com/supakornn/game-shop/pkg/playerCoin/exception"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	playerCoin := new(entities.PlayerCoin)

	if err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Errorf("adding coin failed: %s", err.Error())
		return nil, &_playerCoinException.CoinAdding{}
	}

	return playerCoin, nil
}

func (r *playerCoinRepositoryImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoinShowing := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(
		&entities.PlayerCoin{},
	).Where(
		"player_id = ?",
		playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group(
		"player_id",
	).Scan(playerCoinShowing).Error; err != nil {
		r.logger.Errorf("showing coin failed: %s", err.Error())
		return nil, &_playerCoinException.PlayerCoinShowing{}
	}

	return playerCoinShowing, nil
}
