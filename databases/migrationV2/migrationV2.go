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

	itemsAdding(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func itemsAdding(tx *gorm.DB) {
	items := []entities.Item{
		{
			Name:        "Sword",
			Description: "A sword is a bladed melee weapon.",
			Price:       100,
			Picture:     "https://example.com/sword.jpg",
		},
		{
			Name:        "Shield",
			Description: "A shield is a piece of personal armour held in the hand.",
			Price:       50,
			Picture:     "https://example.com/shield.jpg",
		},
		{
			Name:        "Bow",
			Description: "A bow is a flexible.",
			Price:       200,
			Picture:     "https://example.com/bow.jpg",
		},
		{
			Name:        "Arrow",
			Description: "An arrow is a fin-stabilized projectile that is launched via a bow",
			Price:       10,
			Picture:     "https://example.com/arrow.jpg",
		},
	}

	tx.CreateInBatches(items, len(items))
}
