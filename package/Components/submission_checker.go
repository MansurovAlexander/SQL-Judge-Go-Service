package components

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	JDGECODES "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/StatusCodes"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type CheckScriptsResult struct {
	Result int
	Script string
}

type SubmissionResult struct {
	Status int
	Script string
	Memory float32
	Time   float32
}

func CheckSubmission(inputedScript, correctScript, banned_words, admission_words map[string]string, dbName string, timeLimit, memoryLimit float32) (map[int]SubmissionResult, error) {
	var exitStatus int
	results := make(map[int]SubmissionResult)
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), dbName, viper.GetString("db.password"), viper.GetString("db.sslmode")))
	if err != nil {
		return nil, err
	}
	for key, _ := range correctScript {
		var tempResult SubmissionResult
		var time float32
		if inputedScript[key] == "" {
			exitStatus = JDGECODES.JudgeStatusWrongAnswer
		} else if ContainsBannedWords(inputedScript[key], banned_words[key]) {
			exitStatus = JDGECODES.JudgeStatusBadWord
		} else if !ContainsAdmissionWords(inputedScript[key], admission_words[key]) {
			exitStatus = JDGECODES.JudgeStatusAdmissionWord
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "select") {
			exitStatus, time = CheckSelectScript(testDB, inputedScript[key], correctScript[key], timeLimit)
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "insert") {
			exitStatus = CheckInsertScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "update") {
			exitStatus = CheckUpdateScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "delete") {
			exitStatus = CheckDeleteScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "index") {
			exitStatus, time = CheckIndexScript(testDB, inputedScript[key], correctScript[key], timeLimit)
		}
		key_to_int, _ := strconv.Atoi(key)

		tempResult.Time = time
		tempResult.Memory = 0
		tempResult.Status = exitStatus
		tempResult.Script = inputedScript[key]

		results[key_to_int] = tempResult
	}
	return results, nil
}

func CheckAssignScripts(inputedScripts map[string]string, dbName string) (map[int]CheckScriptsResult, error) {
	var exitStatus int
	results := make(map[int]CheckScriptsResult)
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), dbName, viper.GetString("db.password"), viper.GetString("db.sslmode")))
	if err != nil {
		return nil, err
	}

	for key, _ := range inputedScripts {
		var tempResult CheckScriptsResult
		if inputedScripts[key] == "" {
			exitStatus = JDGECODES.JudgeStatusUnknownError
		} else if strings.Contains(strings.ToLower(inputedScripts[key]), "select") {
			exitStatus = RunSelectScript(inputedScripts[key], testDB)
		} else if strings.Contains(strings.ToLower(inputedScripts[key]), "insert") {
			exitStatus = RunIUDScript(inputedScripts[key], testDB)
		} else if strings.Contains(strings.ToLower(inputedScripts[key]), "update") {
			exitStatus = RunIUDScript(inputedScripts[key], testDB)
		} else if strings.Contains(strings.ToLower(inputedScripts[key]), "delete") {
			exitStatus = RunIUDScript(inputedScripts[key], testDB)
		} else if strings.Contains(strings.ToLower(inputedScripts[key]), "index") {
			exitStatus = RunCreateScripts(inputedScripts[key], testDB)
		}
		key_to_int, _ := strconv.Atoi(key)
		tempResult.Result = exitStatus
		tempResult.Script = inputedScripts[key]
		results[key_to_int] = tempResult
	}
	return results, nil
}

func RunSelectScript(inputedScript string, dbConn *sqlx.DB) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	defer tx.Rollback()

	_, err = tx.QueryContext(ctx, inputedScript)
	if err != nil {
		if err.Error() == "pq: выполнение оператора отменено по запросу пользователя" {
			return JDGECODES.JudgeStatusTimeLimitExceed
		} else if err.Error() == sql.ErrNoRows.Error() {
			return JDGECODES.JudgeStatusEmptyOutput
		}
		return JDGECODES.JudgeStatusUnknownError
	} else {
		return JDGECODES.JudgeStatusChecked
	}
}

// запуск insert, update, delete скриптов (IUD)
func RunIUDScript(inputedScript string, dbConn *sqlx.DB) int {
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	defer tx.Rollback()
	_, err = tx.Exec(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	} else {
		return JDGECODES.JudgeStatusChecked
	}
}

func RunCreateScripts(inputedScript string, dbConn *sqlx.DB) int {
	indexName := ParseIndexName(inputedScript)
	if indexName == "" {
		return JDGECODES.JudgeStatusUnknownError
	}
	_, err := dbConn.Exec(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if _, err := dbConn.Exec("DROP INDEX $1", indexName); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	return JDGECODES.JudgeStatusChecked
}

func CheckSelectScript(dbConn *sqlx.DB, inputedScript, correctScript string, timeLimit float32) (int, float32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Два скрипта проверки
	// Левый отвечает за проверку недостающего количества строк
	// Правый отвечает за проверку лишних строк
	var exceptScriptLeft, exceptScriptRight strings.Builder
	exceptScriptLeft.WriteString("(")
	exceptScriptLeft.WriteString(correctScript)
	exceptScriptLeft.WriteString(") EXCEPT (")
	exceptScriptLeft.WriteString(inputedScript)
	exceptScriptLeft.WriteString(")")

	exceptScriptRight.WriteString("(")
	exceptScriptRight.WriteString(inputedScript)
	exceptScriptRight.WriteString(") EXCEPT (")
	exceptScriptRight.WriteString(correctScript)
	exceptScriptRight.WriteString(")")

	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError, 0
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, exceptScriptLeft.String()).Scan()
	if err != nil {
		if err.Error() == "pq: выполнение оператора отменено по запросу пользователя" {
			return JDGECODES.JudgeStatusTimeLimitExceed, 5000
		} else if err != sql.ErrNoRows {
			return JDGECODES.JudgeStatusWrongAnswer, 0
		}
	}

	err = tx.QueryRowContext(ctx, exceptScriptRight.String()).Scan()
	if err != nil {
		if err.Error() == "pq: выполнение оператора отменено по запросу пользователя" {
			return JDGECODES.JudgeStatusTimeLimitExceed, 5000
		}
		if err != sql.ErrNoRows {
			return JDGECODES.JudgeStatusWrongAnswer, 0
		}
	}

	if err := tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError, 0
	}

	return CheckScriptByTime(dbConn, inputedScript, timeLimit)
}

func CheckInsertScript(dbConn *sqlx.DB, inputedScript, correctScript string) int {
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	_, err = tx.Exec(correctScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	row, err := tx.Query(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if err = row.Scan(); err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if err := tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	return JDGECODES.JudgeStatusCorrectAnswer
}

func CheckUpdateScript(dbConn *sqlx.DB, inputedScript, correctScript string) int {
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	_, err = tx.Exec(correctScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	row, err := tx.Query(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if err = row.Scan(); err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if err := tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	return JDGECODES.JudgeStatusCorrectAnswer
}

func CheckDeleteScript(dbConn *sqlx.DB, inputedScript, correctScript string) int {
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	_, err = tx.Exec(correctScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	row, err := tx.Query(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	if err = row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return JDGECODES.JudgeStatusCorrectAnswer
		} else {
			return JDGECODES.JudgeStatusUnknownError
		}
	}
	if err := tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	return JDGECODES.JudgeStatusWrongAnswer
}

func CheckIndexScript(dbConn *sqlx.DB, inputedScript, correctScript string, timeLimit float32) (int, float32) {
	indexName := ParseIndexName(inputedScript)
	if indexName == "" {
		return JDGECODES.JudgeStatusWrongAnswer, 0
	}
	_, err := dbConn.Exec(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer, 0
	}
	statusCode, time := CheckScriptByTime(dbConn, correctScript, timeLimit)
	if _, err := dbConn.Exec("DROP INDEX $1", indexName); err != nil {
		return statusCode, time
	}
	return statusCode, time
}

func CheckScriptByTime(dbConn *sqlx.DB, inputedScript string, timeLimit float32) (int, float32) {
	var timeFloat float64

	if timeLimit == 0 {
		return JDGECODES.JudgeStatusCorrectAnswer, 0
	}
	//Подготовка скрипта для проверки времени выполнения
	var performanceTest strings.Builder
	performanceTest.WriteString("EXPLAIN (ANALYSE, COSTS OFF, TIMING OFF) ")
	performanceTest.WriteString(inputedScript)
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError, 0
	}
	timeTest, err := tx.Query(performanceTest.String())
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError, 0
	}
	defer timeTest.Close()
	for timeTest.Next() {
		var testResults string
		if err = timeTest.Scan(&testResults); err != nil {
			return JDGECODES.JudgeStatusUnknownError, 0
		}
		if strings.Contains(testResults, "Execution time") {
			tempString := testResults
			tempString = strings.Replace(tempString, "Execution time: ", "", -1)
			tempString = strings.Replace(tempString, " ms", "", -1)
			timeFloat, err = strconv.ParseFloat(tempString, 32)
			if err != nil {
				return JDGECODES.JudgeStatusUnknownError, 0
			}
			if float32(timeFloat) > timeLimit {
				return JDGECODES.JudgeStatusTimeLimitExceed, float32(timeFloat)
			}
		}
	}

	if err = tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError, float32(timeFloat)
	}

	return JDGECODES.JudgeStatusCorrectAnswer, float32(timeFloat)
}

func CheckExistSubmission(dbConn *sqlx.DB, studentID, assignID, subtaskID int) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM submission WHERE student_id = $1 AND assign_id = $2 and subtask_id = $3)"
	row := dbConn.QueryRow(query, studentID, assignID, subtaskID)
	if err := row.Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func CreateSubmissionResults(dbConn *sqlx.DB, submissionID, assignID, time, memory, studentID, statusId, subtaskID int, inputedScript string) error {
	exist, _ := CheckExistSubmission(dbConn, studentID, assignID, subtaskID)
	if exist {
		query := "UPDATE submission SET script = $1, time = $2, memory = $3, status_id = $4 WHERE student_id = $5 AND assign_id = $6 AND subtask_id = $7"
		_, err := dbConn.Exec(query, inputedScript, time, memory, statusId, studentID, assignID, subtaskID)
		if err != nil {
			return err
		}
		return nil
	}
	query := "INSERT INTO submission(submission_id, assign_id, student_id, time, memory, script, status_id, subtask_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err := dbConn.Exec(query, submissionID, assignID, studentID, time, memory, inputedScript, statusId, subtaskID)
	if err != nil {
		return err
	}
	return nil
}

func ContainsBannedWords(inputedScript, bannedWords string) bool {
	bannedWordList := strings.Split(bannedWords, ",")
	RemoveEmpty(&bannedWordList)
	for i := 0; i < len(bannedWordList); i++ {
		reg, _ := regexp.Compile(`\b` + strings.ToLower(bannedWordList[i]) + `\b`)
		if reg.MatchString(inputedScript) {
			return true
		}
	}
	return false
}

func ContainsAdmissionWords(inputedScript, admissionWords string) bool {
	admissionWordList := strings.Split(admissionWords, ",")

	//Из парсера
	RemoveEmpty(&admissionWordList)
	for i := 0; i < len(admissionWordList); i++ {
		reg, _ := regexp.Compile(`\b` + strings.ToLower(admissionWordList[i]) + `\b`)
		if !reg.MatchString(inputedScript) {
			return false
		}
	}
	return true
}
