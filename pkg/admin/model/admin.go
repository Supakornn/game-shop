package model

type AdminCreatingReq struct {
	ID     string `gorm:"primaryKey;type:varchar(64);"`
	Email  string `gorm:"type:varchar(128);unique;not null;"`
	Name   string `gorm:"type:varchar(128);not null;"`
	Avatar string `gorm:"type:varchar(256);not null;default:'';"`
}
