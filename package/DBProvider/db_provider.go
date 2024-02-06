package dbprovider

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func IsExist(dbName string) (bool, *sqlx.DB, error) {
	if err := initConfig(); err != nil {
		return false, nil, err
	}
	if err := godotenv.Load(); err != nil {
		return false, nil, err
	}
	cfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	dbProvider, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return false, nil, err
	}
	row := dbProvider.QueryRow("SELECT 1 FROM pg_database WHERE datname=$1", strings.ToLower(dbName))
	var count int
	if row.Scan(&count); count < 1 {
		err = dbProvider.Close()
		if err != nil {
			return false, nil, err
		}
		return false, nil, nil
	}
	err = dbProvider.Close()
	if err != nil {
		return true, nil, err
	}
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, strings.ToLower(dbName), cfg.Password, cfg.SSLMode))
	if err != nil {
		return false, nil, err
	}
	return true, testDB, nil
	/*testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, strings.ToLower(dbName), cfg.Password, cfg.SSLMode))
	if err != nil {
		return 0, err
	}
	if err = testDB.Ping(); err != nil {
		return 0, err
	}
	_, err = testDB.Exec(dbCreateScript)
	if err != nil {
		return 0, err
	}*/
}
func CreateTestDB(dbName, createScript string) (*sqlx.DB, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	dbProvider, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	transaction, err := dbProvider.Begin()
	if err != nil {
		return nil, err
	}
	_, err = dbProvider.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, strings.ToLower(dbName), cfg.Password, cfg.SSLMode))
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	_, err = testDB.Exec(createScript)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return testDB, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
