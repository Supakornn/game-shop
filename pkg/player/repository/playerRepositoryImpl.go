package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/entities"
	_playerException "github.com/supakornn/game-shop/pkg/player/exception"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepository(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{db: db, logger: logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)
	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Errorf("creating player failed: %v", err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("finding player by id failed: %v", err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
