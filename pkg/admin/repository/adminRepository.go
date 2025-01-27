package repository

import "github.com/supakornn/game-shop/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}
