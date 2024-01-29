package repository

import (
	"fmt"
	"math/big"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type JudgePostgres struct {
	db *sqlx.DB
}

func NewJudgeService(db *sqlx.DB) *JudgePostgres {
	return &JudgePostgres{db: db}
}

func (r *JudgePostgres) CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName string, assignID, studentID big.Int, timeLimit, memoryLimit int) (big.Int, error) {
	if err := initConfig(); err != nil {
		return *big.NewInt(-1), err
	}
	cfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	transaction, err := r.db.Begin()
	if err != nil {
		return *big.NewInt(-1), err
	}
	_, err = r.db.Exec(dbCreateScript)
	if err != nil {
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, dbName, cfg.Password, cfg.SSLMode))
	if err != nil {
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	if err = testDB.Ping(); err != nil {
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	initSavepoint, err := testDB.Begin()
	if err != nil {
		initSavepoint.Rollback()
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	correctAnswer, err := testDB.Query(correctScript)
	if err != nil {
		initSavepoint.Rollback()
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	err = initSavepoint.Rollback()
	if err != nil {
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	initSavepoint, err = testDB.Begin()
	if err != nil {
		initSavepoint.Rollback()
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	studentAnswer, err := testDB.Query(inputedScript)
	if err != nil {
		initSavepoint.Rollback()
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	err = initSavepoint.Rollback()
	if err != nil {
		transaction.Rollback()
		return *big.NewInt(-1), err
	}
	if correctAnswer != studentAnswer {
		var id big.Int
		query := "INSERT INTO submission(assign_id, student_id, time, memory, script, grade) VALUES ($1,$2,$3,$4,$5,$6) returning id"
		row := r.db.QueryRow(query, assignID, studentID, 0, 0, inputedScript, "Неправильный вывод!")
		if err := row.Scan(&id); err != nil {
			return *big.NewInt(-1), err
		}
		return id, nil
	}
	if err := transaction.Rollback(); err != nil {
		return *big.NewInt(-1), err
	}
	return *big.NewInt(-1), nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
