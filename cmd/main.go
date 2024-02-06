package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	sqljudge "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin"
	handler "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Handler"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error ocured while reading config: " + err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error ocured while load env variables: " + err.Error())
	}
	/*connstring := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.name"), viper.GetString("db.sslmode"))
	cmd := exec.Command("goose", "-dir", "db/migrations", "postgres", connstring, "up")
	path, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("error while getting path of project: %s", err.Error())
	}
	cmd.Dir = path
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		logrus.Fatalf("error while commit migration: %s.", err.Error())
	}*/
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error while initializing DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(sqljudge.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
