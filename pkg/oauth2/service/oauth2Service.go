package service

import (
	_adminModel "github.com/supakornn/game-shop/pkg/admin/model"
	_playerModel "github.com/supakornn/game-shop/pkg/player/model"
)

type Oauth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsPlayerExist(playerID string) bool
	IsAdminExist(adminID string) bool
}
