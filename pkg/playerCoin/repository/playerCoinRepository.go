package repository

import (
	"github.com/supakornn/game-shop/entities"
	_playerCoinModel "github.com/supakornn/game-shop/pkg/playerCoin/model"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
