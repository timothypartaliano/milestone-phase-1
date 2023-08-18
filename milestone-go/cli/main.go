package cli

import (
	"database/sql"
	"fmt"
	"log"
	"milestone-go/entity"
	"milestone-go/handler"

	"os"
)

type Cli struct {
	UserHandler  handler.User
	GameHandler  handler.Game
	DB           *sql.DB
	LoggedInUser entity.User
}

func (c *Cli) AuthMenu() {
	fmt.Println("Welcome to ultimate Games Store!")
	fmt.Println("[COMMAND]: 	Description")
	fmt.Println("[sign-up]: 	Register/create new user")
	fmt.Println("[sign-in]: 	login into existing user")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "sign-up":
		c.Register()
	case "sign-in":
		loggedInUser := c.Login()
		if loggedInUser.Id != 0 {
			c.LoggedInUser = loggedInUser
		}
		c.MainMenu()
	default:
		os.Exit(1)
	}
}

func (c *Cli) MainMenu() {
	fmt.Println("\nWelcome to ultimate Games Store!")
	fmt.Println("[COMMAND]: 		Description")
	fmt.Println("[list-games]: 		Retrieve all games")
	fmt.Println("[buy-games]:        	Buy a game")
	fmt.Println("[list-orders]: 		Retrieve all orders")
	fmt.Println("or press anything to back to main menu")
	fmt.Println("\n-----")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "list-games":
		c.ListGame()
	case "buy-games":
		c.BuyGame(c.LoggedInUser)
	case "list-orders":
		c.ListOrders()
	default:
		os.Exit(1)
	}
}

func (c *Cli) Register() {
	fmt.Println("\nPlease Register!")

	var username, email, password string

	fmt.Println("Please input username:")
	_, errUsername := fmt.Scanln(&username)

	fmt.Println("Please input email:")
	_, errEmail := fmt.Scanln(&email)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	if errUsername != nil ||
		errPassword != nil ||
		errEmail != nil {
		log.Fatal("failed to scan input")
	}

	err := c.UserHandler.Register(username, email, password)
	if err != nil {
		fmt.Println("failed to create new user...")
	} else {
		fmt.Println("\nSuccess register user!")
	}
	c.AuthMenu()
}

func (c *Cli) Login() entity.User {
	fmt.Println("\nPlease Login!")

	var email, password string

	fmt.Println("Please input email:")
	_, errEmail := fmt.Scanln(&email)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	if errEmail != nil || errPassword != nil {
		log.Fatal("failed to scan input")
	}

	loggedInUserID, err := c.UserHandler.Login(email, password)
	if err != nil {
		fmt.Println("\nFailed to login!")
		c.AuthMenu()
		return entity.User{}
	} else {
		fmt.Println("\nSuccess login user!")
		return loggedInUserID
	}
}

func (c *Cli) ListGame() {
	c.GameHandler.ShowGames()

	c.MainMenu()
}

func (c *Cli) BuyGame(loggedInUser entity.User) {
	fmt.Println("\nPlease select a game to buy:")

	c.GameHandler.ShowGames()

	var game_id int
	fmt.Print("Enter the Game ID you want to buy: ")
	_, err := fmt.Scanln(&game_id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("Enter the amount you want to buy: ")
	var amount int
	_, err = fmt.Scanln(&amount)
	if err != nil {
		log.Fatal(err.Error())
	}

	selectedGame, err := c.GameHandler.GetGameByID(game_id)
	if err != nil {
		fmt.Println("Failed to retrieve game information:", err)
		c.MainMenu()
		return
	}

	totalPrice := selectedGame.Price * amount

	fmt.Printf("Total Price: $%d\n", totalPrice)

	userID := loggedInUser.Id
	err = c.GameHandler.BuyGame(userID, game_id, amount)
	if err != nil {
		fmt.Println("Failed to buy the game:", err)
	} else {
		fmt.Println("Game bought successfully!")
	}
	c.MainMenu()
}

func (c *Cli) ListOrders() {
	fmt.Println("\nList of Orders:")
	c.GameHandler.ShowOrders(c.LoggedInUser.Id)
	c.MainMenu()
}
