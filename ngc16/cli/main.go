package cli

import (
	"database/sql"
	"fmt"
	"log"
	"ngc16/handler"
	"os"
)

type Cli struct{
	UserHandler handler.User
	GameHandler handler.Game
	DB *sql.DB
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
	fmt.Println("[buy-game]:        Buy a game")
	fmt.Println("[list-orders]: 		Retrieve all orders")
	fmt.Println("[create-products]: 	Create new products")
	fmt.Printf("or press anything to back to main menu\n")
	// fmt.Println("press anything to exit")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "list-games":
		c.ListGame()
	case "buy-game":
        c.BuyGame()
	default:
		os.Exit(1)
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

func (c Cli) ListGame() {
	c.GameHandler.ShowGames()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(games)
}

func (c Cli) BuyGame() {
	fmt.Println("\nPlease select a game to buy:")

    c.GameHandler.ShowGames()

    var gameID int
    fmt.Print("Enter the Game ID you want to buy: ")
    _, err := fmt.Scanln(&gameID)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Print("Enter the amount you want to buy: ")
    var amount int
    _, err = fmt.Scanln(&amount)
    if err != nil {
        log.Fatal(err.Error())
    }

    userID := 1
    err = c.GameHandler.BuyGame(userID, gameID, amount)
    if err != nil {
        fmt.Println("Failed to buy the game:", err)
    } else {
        fmt.Println("Game bought successfully!")
    }
    c.MainMenu()
}