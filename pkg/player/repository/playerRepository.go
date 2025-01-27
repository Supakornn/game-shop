package repository

import "github.com/supakornn/game-shop/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
