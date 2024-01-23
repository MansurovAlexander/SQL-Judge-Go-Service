package main

import (
	"fmt"
	"log"

	sqljudge "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin"
	handler "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Handler"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error ocured while reading config: " + err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error while initializing DB: %s", err.(*pq.Error).Code.Name())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(sqljudge.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error ocured while running http server: " + err.Error())
	} else {
		fmt.Println("Server is listening!")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
