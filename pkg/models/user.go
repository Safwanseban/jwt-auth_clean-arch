package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Safwanseban/jwt-auth/pkg/config"
	"github.com/Safwanseban/jwt-auth/pkg/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	R          = config.ConfigLoad()
	id         uint
	insertUser = `INSERT INTO users (username,password)
	VALUES ($1, $2)
	RETURNING id`
	selectUser = `SELECT username,password FROM users WHERE username=$1`
)

type User struct {
	gorm.Model
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) InsertOne() error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db, err := db.ConnectToDb(R, &sql.DB{})

	if err != nil {
		fmt.Println("hai")
		return err
	}
	if err := db.QueryRowContext(ctx, insertUser, user.Username, user.Password).Scan(&id); err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func (user User) FindUser() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db, err := db.ConnectToDb(R, &sql.DB{})
	if err != nil {
		return nil, err
	}

	if err := db.QueryRowContext(ctx, selectUser, user.Username).Scan(&user.Username, &user.Password); err != nil {
		return nil, err
	}
	// fmt.Println(user.Password)
	return &user, nil
}

func (user *User) HashUserPassword(password string) ([]byte, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return hashPassword, nil

}
func (user *User) CheckPassword(hashedPass string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(user.Password)); err != nil {
		fmt.Println(user.Password, hashedPass)
		return err
	}

	return nil
}

