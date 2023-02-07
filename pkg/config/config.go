package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructures:"PORT"`
	Db_Host     string `mapstructures:"DB_HOST"`
	Db_User     string `mapstructures:"DB_USER"`
	DB_Password string `mapstructures:"DB_PASSWORD"`
	Db_Port     uint `mapstructures:"DB_PORT"`
	Db_Name     string `mapstructures:"DB_NAME"`
}

func ConfigLoad() (env *Config) {

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error readinmg env file", err)
	}
	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(err)
	}
	return 

}
