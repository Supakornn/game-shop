package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/supakorn/game-shop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

var (
	postgresDBInstance *PostgresDB
	once               sync.Once
)

func NewPostgresDB(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.SSLMode,
			conf.Schema,
		)

		con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to %s", conf.DBName)

		postgresDBInstance = &PostgresDB{con}
	})

	return postgresDBInstance
}

func (db *PostgresDB) ConnectionGetting() *gorm.DB {
	return db.DB
}
