package repository

import (
	"fmt"
	"time"

	dbprovider "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/DBProvider"
	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type JudgePostgres struct {
	db *sqlx.DB
}

func NewJudgeService(db *sqlx.DB) *JudgePostgres {
	return &JudgePostgres{db: db}
}

func (r *JudgePostgres) CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName string, assignID, studentID, timeLimit, memoryLimit int) (int, error) {
	var grade string
	/*if methods.ContainsBannedWords(inputedScript, bannedWords) {
		grade = "Скрипт содержит запрещенные слова."
		var id int
		query := "INSERT INTO submission(assign_id, student_id, time, memory, script, grade) VALUES ($1,$2,$3,$4,$5,$6) returning id"
		row := r.db.QueryRow(query, assignID, studentID, 0, 0, inputedScript, grade)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}*/
	isExist, db, err := dbprovider.IsExist(dbName)
	if err != nil {
		return 0, err
	}
	if !isExist {
		db, err = dbprovider.CreateTestDB(dbName, dbCreateScript)
		if err != nil {
			return 0, err
		}
	}
	transaction, err := db.Begin()
	if err != nil {
		return 0, err
	}
	correctAnswer, err := db.Query(correctScript)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	err = transaction.Rollback()
	if err != nil {
		return 0, err
	}
	transaction, err = db.Begin()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	startTime := time.Now()
	fmt.Print(inputedScript)
	studentAnswer, err := db.Query(inputedScript)
	totalTime := time.Since(startTime)
	if err != nil {
		transaction.Rollback()
		logrus.Fatalf("error while running student script: %s", err.(*pq.Error).Code.Name())
		return 0, err
	}
	err = transaction.Rollback()
	if err != nil {
		return 0, err
	}
	studentAnswersJson := jsonify.Jsonify(studentAnswer)
	correctAnswerJson := jsonify.Jsonify(correctAnswer)
	fmt.Print(studentAnswersJson)
	fmt.Print(correctAnswerJson)
	if len(correctAnswerJson) != len(studentAnswersJson) {
		grade = "Неверный ответ!"
	} else if float32(timeLimit) != 0 && float32(totalTime) > float32(timeLimit) {
		grade = "Не выполнено ограничение по времени"
	} else {
		for i := 0; i < len(correctAnswerJson); i++ {
			if correctAnswerJson[i] != studentAnswersJson[i] {
				grade = "Не верный ответ!"
				break
			}
		}
		if grade != "Неверный ответ!" {
			grade = "Все правильно!"
		}
	} /*else if memory>memoryLimit{
		grade="Не выполнено ограничение по времени"
	}*/
	var id int
	query := "INSERT INTO submission(assign_id, student_id, time, memory, script, grade) VALUES ($1,$2,$3,$4,$5,$6) returning id"
	row := r.db.QueryRow(query, assignID, studentID, totalTime, 0, inputedScript, grade)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
