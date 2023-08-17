package handler

import (
	"context"
	"database/sql"
	"errors"
	"milestone-go/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB *sql.DB
}

func (u User) Register(username, email, password string) error {
	ctx := context.Background()
	query := `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
	`
	
	byteHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.DB.ExecContext(ctx, query, username, email, string(byteHashed))

	if err != nil {
		return err
	}

	return nil
}

func (u User) Login(email, password string) (entity.User, error) {
	user := entity.User{}
	ctx := context.Background()
	query := `
		SELECT user_id, username, email, password
		FROM users WHERE email = ?
	`

	result, err := u.DB.QueryContext(ctx, query, email)
	if err != nil {
		return user, err
	}

	if result.Next() {
		result.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		return user, err
	} else {
		return user, errors.New("email/password invalid")
	}
}