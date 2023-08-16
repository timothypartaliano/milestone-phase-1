package handler

import (
	"context"
	"database/sql"
	"errors"
	// "fmt"
	"ngc16/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB *sql.DB
}

func (u User) Register(username, password string, age int) error {
	// fmt.Println(username, password, age)
	ctx := context.Background()
	query := `
		INSERT INTO users (username, password, age)
		VALUES (?, ?, ?)
	`
	
	byteHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.DB.ExecContext(ctx, query, username, string(byteHashed), age)

	if err != nil {
		return err
	}

	return nil
}

func (u User) Login(username, password string) (entity.User, error) {
	user := entity.User{}
	ctx := context.Background()
	query := `
		SELECT user_id, username, password, age
		FROM users WHERE username = ?
	`

	result, err := u.DB.QueryContext(ctx, query, username)
	if err != nil {
		return user, err
	}

	if result.Next() {
		result.Scan(&user.Id, &user.Username, &user.Password, &user.Age)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		return user, err
	} else {
		return user, errors.New("username/password invalid")
	}
}