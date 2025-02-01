package repository

import (
	"github.com/supakornn/game-shop/entities"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
