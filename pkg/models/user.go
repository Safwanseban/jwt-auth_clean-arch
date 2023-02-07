package models

import (
	"database/sql"
	"fmt"

	"github.com/Safwanseban/jwt-auth/pkg/config"
	"github.com/Safwanseban/jwt-auth/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	R          = config.ConfigLoad()
	insertUser = `INSERT INTO users (username,password)
	VALUES ($1, $2)
	RETURNING id`
)

type User struct {
	gorm.Model
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) InsertOne() error {
	db, err := db.ConnectToDb(R, &sql.DB{})
	fmt.Println(R.Db_Host)
	if err != nil {
		fmt.Println("hai")
		return err
	}
	if err := db.QueryRow(insertUser, user.Username, user.Password).Scan(user); err != nil {
		return err
	}
	return nil
}

func CreateUser(ctx *gin.Context, user User) error {
	return user.InsertOne()
}
