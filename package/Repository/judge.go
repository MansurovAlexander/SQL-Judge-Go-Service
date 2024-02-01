package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/bdwilliams/go-jsonify/jsonify" 
)

type JudgePostgres struct {
	db *sqlx.DB
}

func NewJudgeService(db *sqlx.DB) *JudgePostgres {
	return &JudgePostgres{db: db}
}

func (r *JudgePostgres) CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName string, assignID, studentID, timeLimit, memoryLimit int) (int, error) {
	var grade string
	if err := initConfig(); err != nil {
		return 0, err
	}
	cfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	dbProvider, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return 0, err
	}
	row := dbProvider.QueryRow("SELECT 1 FROM pg_database WHERE datname=$1", strings.ToLower(dbName))
	var count int
	if row.Scan(&count); count < 1 {
		_, err = dbProvider.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil {
			return 0, err
		}
		err = dbProvider.Close()
		if err != nil {
			return 0, err
		}
	}
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
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
	}
	initSavepoint, err := testDB.Begin()
	if err != nil {
		initSavepoint.Rollback()
		return 0, err
	}
	correctAnswer, err := testDB.Query(correctScript)
	if err != nil {
		initSavepoint.Rollback()
		return 0, err
	}
	err = initSavepoint.Rollback()
	if err != nil {
		return 0, err
	}
	initSavepoint, err = testDB.Begin()
	if err != nil {
		initSavepoint.Rollback()
		return 0, err
	}
	startTime := time.Now()
	fmt.Print(inputedScript)
	studentAnswer, err := testDB.Query(inputedScript)
	totalTime := time.Since(startTime)
	if err != nil {
		initSavepoint.Rollback()
		logrus.Fatalf("error while running student script: %s", err.(*pq.Error).Code.Name())
		return 0, err
	}
	err = initSavepoint.Rollback()
	if err != nil {
		return 0, err
	}
	studentAnswersJson:=jsonify.Jsonify(studentAnswer)
	fmt.Print(studentAnswersJson)
	if correctAnswer != studentAnswer {
		grade = "Неверный ответ!"
	} else if float32(totalTime) > float32(timeLimit) {
		grade = "Не выполнено ограничение по времени"
	} else {
		grade = "Все правильно!"
	} /*else if memory>memoryLimit{
		grade="Не выполнено ограничение по времени"
	}*/
	var id int
	query := "INSERT INTO submission(assign_id, student_id, time, memory, script, grade) VALUES ($1,$2,$3,$4,$5,$6) returning id"
	row = r.db.QueryRow(query, assignID, studentID, totalTime, 0, inputedScript, grade)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
