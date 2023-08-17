package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"ngc16/entity"
)

type Game struct {
    DB *sql.DB
}

func (g Game) ShowGames() {
    var games entity.Game

    // query := `select game_id, title, price from games;`
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

// package handler

// import (
// 	"database/sql"
// 	"ngc16/entity"
// )

// func listGames(db *sql.DB) ([]entity.Game, error) {
// 	var games []entity.Game
// 	rows, err := db.Query("SELECT game_id, title, description FROM games")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var game entity.Game
// 		err := rows.Scan(&game.ID, &game.Title, &game.Price)
// 		if err != nil {
// 			return nil, err
// 		}
// 		games = append(games, game)
// 	}

// 	return games, nil
// }