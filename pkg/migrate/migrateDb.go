package migrate

import (
	"fmt"
	"log"

	"github.com/Safwanseban/jwt-auth/pkg/config"
	"github.com/Safwanseban/jwt-auth/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDB() {
	var conf = config.ConfigLoad()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", conf.Db_Host, conf.Db_Port, conf.Db_User,
		conf.DB_Password, conf.Db_Name)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	fmt.Println(conf.Db_Host)
	if err != nil {
		log.Println("error connecting to gorm db")
	}
	db.AutoMigrate(&models.User{})

}
