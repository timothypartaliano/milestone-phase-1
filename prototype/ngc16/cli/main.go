package cli

import (
	"fmt"
	"log"
	"ngc16/handler"
	"os"
)

type Cli struct{
	UserHandler handler.User
}

func (c Cli) AuthMenu() {
	// fmt.Println("\n\n -----")
	fmt.Println("welcome to Games Store!")
	fmt.Println("[COMMAND]: 	Description")
	fmt.Println("[sign-up]: 	Register/create new user")
	fmt.Println("[sign-in]: 	login into existing user")
	// fmt.Println("press anything to exit")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "sign-up":
		c.Register()
	case "sign-in":
		c.Login()
	default:
		os.Exit(1)
	}
}

func (c Cli) MainMenu() {
	fmt.Println("\nwelcome to Games Store")
	fmt.Println("[COMMAND]: 		Description")
	fmt.Println("[list-games]: 		Retrieve all games")
	fmt.Println("[list-orders]: 		Retrieve all orders")
	fmt.Println("[create-products]: 	Create new products")
	fmt.Println("or press anything to back to main menu")
	// fmt.Println("press anything to exit")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (c Cli) Register() {
	fmt.Println("\nPlease Register!")

	var username, password string
	var age int

	fmt.Println("Please input username:")
	_, errUsername := fmt.Scanln(&username)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	fmt.Println("Please input age:")
	_, errAge := fmt.Scanln(&age)

	if errUsername != nil ||
		errPassword != nil ||
		errAge != nil {
		log.Fatal("failed to scan input")
	}

	err := c.UserHandler.Register(username, password, age)
	if err != nil {
		fmt.Println("failed to create new user...")
	} else {
		fmt.Println("\nSuccess register user!")
	}
	c.AuthMenu()
}

func (c Cli) Login() {
	fmt.Println("\nPlease Login!")
	
	var username, password string

	fmt.Println("Please input username:")
	_, errUsername := fmt.Scanln(&username)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	if errUsername != nil ||
		errPassword != nil {
		log.Fatal("failed to scan input")
	}
	// user, err := c.UserHandler.Login(username, password)
	_, err := c.UserHandler.Login(username, password)
	if err != nil {
		// fmt.Println("error: ", err.Error())
		fmt.Println("\nFailed to login!")
		c.Login()
	} else {
		fmt.Println("\nSuccess login user!")
		// fmt.Println(user, "<---")
		c.MainMenu()
	}

	// c.AuthMenu()
}