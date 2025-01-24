package main

import (
	"github.com/supakornn/game-shop/config"
	"github.com/supakornn/game-shop/databases"
	"github.com/supakornn/game-shop/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDB(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
