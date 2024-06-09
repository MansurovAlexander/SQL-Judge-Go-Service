package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	server "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin"
	handler "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Handler"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error ocured while reading config: " + err.Error())
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
		logrus.Fatalf("error while initializing DB: %s", err.Error())
	}
	if err := createPostgresEnv(); err != nil {
		logrus.Fatalf("error while creating postgres env: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(server.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes(viper.GetString("white-ip-4"), viper.GetString("white-ip-6"))); err != nil {
			logrus.Fatal("error ocured while running http server: " + err.Error())
		}
	}()
	logrus.Print("SQL-JUDGE app started...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("SQL-JUDGE app shutting down...")
	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error ocured while closing server: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error ocured while closing database connection: %s", err.Error())
	}
}

func createPostgresEnv() error {
	return os.Setenv("PGPASSWORD", viper.GetString("db.password"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
