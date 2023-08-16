package main

import (
	"ngc16/cli"
	"ngc16/config"
	"ngc16/handler"
)

func main() {
	db := config.InitDatabase("root:@tcp(localhost:3306)/mp1")
	// fmt.Println(db, "<---")

	userHandler := handler.User{DB: db}
	app := cli.Cli{UserHandler: userHandler}
	app.AuthMenu()
}