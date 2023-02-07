package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Safwanseban/jwt-auth/pkg/config"

	_ "github.com/lib/pq"
)

func ConnectToDb(conf *config.Config, db *sql.DB) (*sql.DB, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", conf.Db_Host, conf.Db_Port, conf.Db_User,
		conf.DB_Password, conf.Db_Name)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error")
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("errord")
		return nil, err
	}
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err!=nil{
	// 	log.Println("errord")
	// 	return nil,err
	// }

	// m, err := migrate.NewWithDatabaseInstance(

	// 	"file://db/migrations",
	// 	"postgres", driver)
	if err != nil {
		return nil, err
	}
	
	return db, nil

}
