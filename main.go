package main

import (
	"github.com/supakorn/game-shop/config"
	"github.com/supakorn/game-shop/databases"
	"github.com/supakorn/game-shop/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDB(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}
