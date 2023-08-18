package main

import (
	"milestone-go/cli"
	"milestone-go/config"
	"milestone-go/handler"
)

func main() {

	db := config.InitDatabase("root:X0j1SrIMNOBxiMbMJVvz@tcp(containers-us-west-150.railway.app:5645)/railway")

	userHandler := handler.User{DB: db}
	gameHandler := handler.Game{DB: db}
	app := cli.Cli{UserHandler: userHandler, GameHandler: gameHandler, DB: db}
	app.AuthMenu()
}
