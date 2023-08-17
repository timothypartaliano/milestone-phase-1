package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"milestone-go/entity"
)

type Game struct {
    DB *sql.DB
}

func (g Game) ShowGames() {
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

func (g Game) BuyGame(userID, gameID, amount int) error {
    query := `INSERT INTO orders (user_id, game_id, amount) VALUES (?, ?, ?);`
    _, err := g.DB.Exec(query, userID, gameID, amount)
    return err
}   

func (g Game) ShowOrders() {
    query := `
        SELECT o.order_id, u.username, g.title, o.amount
        FROM orders o
        JOIN users u ON o.user_id = u.user_id
        JOIN games g ON o.game_id = g.game_id;
    `

    ctx := context.Background()
    rows, err := g.DB.QueryContext(ctx, query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Order ID\tUsername\tGame Title\tAmount")
    fmt.Println("------------------------------------------------------------")

    for rows.Next() {
        var orderID, amount int
        var username, title string
        if err := rows.Scan(&orderID, &username, &title, &amount); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d\t\t%s\t\t%s\t%d\n", orderID, username, title, amount)
    }
}