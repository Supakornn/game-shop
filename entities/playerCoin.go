package entities

import "time"

type PlayerCoin struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID  string    `gorm:"type:varchar(64);not null;"`
	Amount    int64     `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
}
