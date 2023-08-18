package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"milestone-go/entity"
)

type Game struct {
	DB *sql.DB
}

func (g *Game) GetGameByID(gameID int) (entity.Game, error) {
	query := "SELECT game_id, Title, Price FROM games WHERE game_id = ?"
	row := g.DB.QueryRow(query, gameID)

	var game entity.Game
	err := row.Scan(&game.GameID, &game.Title, &game.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Game{}, errors.New("game not found")
		}
		return entity.Game{}, err
	}

	return game, nil
}

func (g *Game) ShowGames() {
	var games entity.Game

	query := `select * from games;`

	ctx := context.Background()
	rows, err := g.DB.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("No.    Title               Price")
	for rows.Next() {
		err := rows.Scan(&games.GameID, &games.Title, &games.Price)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d  %s          $%d\n", games.GameID, games.Title, games.Price)

	}
}

func (g *Game) BuyGame(userID, gameID, amount int) error {
	query := `INSERT INTO orders (user_id, game_id, amount) VALUES (?, ?, ?);`
	_, err := g.DB.Exec(query, userID, gameID, amount)
	return err
}

func (g *Game) ShowOrders(userID int) {
	query := `
        SELECT o.order_id, g.title, o.amount
        FROM orders o
        JOIN users u ON o.user_id = u.user_id
        JOIN games g ON o.game_id = g.game_id
        WHERE u.user_id = ?;
    `

	ctx := context.Background()
	rows, err := g.DB.QueryContext(ctx, query, userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Order ID\tGame Title\tAmount")
	fmt.Println("------------------------------------------------------------")

	for rows.Next() {
		var orderID, amount int
		var title string
		if err := rows.Scan(&orderID, &title, &amount); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t\t%s\t%d\n", orderID, title, amount)
	}
}
