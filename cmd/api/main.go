package main

import (
	"database/sql"
	"log"

	http "github.com/Safwanseban/jwt-auth/pkg/api"
	"github.com/Safwanseban/jwt-auth/pkg/config"
	"github.com/Safwanseban/jwt-auth/pkg/db"
	"github.com/Safwanseban/jwt-auth/pkg/migrate"
	"github.com/gin-gonic/gin"
)

var conf = config.ConfigLoad()

func init() {

	migrate.MigrateDB()
}

type app struct {
	engine *gin.Engine
	db     *sql.DB
}

func main() {

	r := http.NewServerHttp(app{}.engine)
	_, err := db.ConnectToDb(conf, app{}.db)
	if err != nil {
		log.Println("err conncting to db", err)
	}

	r.Run(conf.Port)
}
