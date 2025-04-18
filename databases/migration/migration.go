package main

import (
	"github.com/supakornn/game-shop/config"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDB(conf.Database)

	tx := db.Connect().Begin()

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purcheseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purcheseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
