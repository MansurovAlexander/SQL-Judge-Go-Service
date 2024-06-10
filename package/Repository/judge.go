package repository

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	components "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Components"
	JDGECODES "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/StatusCodes"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type JudgePostgres struct {
	db *sqlx.DB
}

func NewJudgeService(db *sqlx.DB) *JudgePostgres {
	return &JudgePostgres{db: db}
}

func (r *JudgePostgres) CheckSubmission(inputedScript, correctScript, banned_words, admission_words map[string]string, dbName string, submissionID, assignID, studentID, timeLimit, memoryLimit int) (int, error) {
	var exitStatus int
	testDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), dbName, viper.GetString("db.password"), viper.GetString("db.sslmode")))
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError, err
	}
	for key, _ := range correctScript {
		if inputedScript[key] == "" {
			exitStatus = JDGECODES.JudgeStatusWrongAnswer
		} else if ContainsBannedWords(inputedScript[key], banned_words[key]) {
			exitStatus = JDGECODES.JudgeStatusBadWord
		} else if !ContainsAdmissionWords(inputedScript[key], admission_words[key]) {
			exitStatus = JDGECODES.JudgeStatusAdmissionWord
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "select") {
			exitStatus = CheckSelectScript(testDB, inputedScript[key], correctScript[key], timeLimit)
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "insert") {
			exitStatus = CheckInsertScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "update") {
			exitStatus = CheckUpdateScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "delete") {
			exitStatus = CheckDeleteScript(testDB, inputedScript[key], correctScript[key])
		} else if strings.Contains(strings.ToLower(inputedScript[key]), "index") {
			exitStatus = CheckIndexScript(testDB, inputedScript[key], correctScript[key], timeLimit)
		}
		key_to_int, _ := strconv.Atoi(key)
		err = CreateSubmissionResults(r.db, submissionID, assignID, timeLimit, memoryLimit, studentID, exitStatus, key_to_int, inputedScript[key])
		if err != nil {
			return JDGECODES.JudgeStatusUnknownError, err
		}
	}
	return JDGECODES.JudgeStatusChecked, nil
}

func CheckSelectScript(dbConn *sqlx.DB, inputedScript, correctScript string, timeLimit int) int {
	var exitStatus int

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
		return JDGECODES.JudgeStatusUnknownError
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, exceptScriptLeft.String()).Scan()
	if err != nil {
		if err.Error() == "pq: выполнение оператора отменено по запросу пользователя" {
			return JDGECODES.JudgeStatusTimeLimitExceed
		} else if err != sql.ErrNoRows {
			return JDGECODES.JudgeStatusWrongAnswer
		}
	}

	err = tx.QueryRowContext(ctx, exceptScriptRight.String()).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return JDGECODES.JudgeStatusWrongAnswer
		}
	}

	if err := tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}

	exitStatus = CheckScriptByTime(dbConn, inputedScript, timeLimit)

	return exitStatus
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

func CheckIndexScript(dbConn *sqlx.DB, inputedScript, correctScript string, timeLimit int) int {
	indexName := components.ParseIndexName(inputedScript)
	if indexName == "" {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	_, err := dbConn.Exec(inputedScript)
	if err != nil {
		return JDGECODES.JudgeStatusWrongAnswer
	}
	statusCode := CheckScriptByTime(dbConn, correctScript, timeLimit)
	if _, err := dbConn.Exec("DROP INDEX $1", indexName); err != nil {
		return statusCode
	}
	return statusCode
}

func CheckScriptByTime(dbConn *sqlx.DB, inputedScript string, timeLimit int) int {
	if timeLimit == 0 {
		return JDGECODES.JudgeStatusCorrectAnswer
	}
	//Подготовка скрипта для проверки времени выполнения
	var performanceTest strings.Builder
	performanceTest.WriteString("EXPLAIN (ANALYSE, COSTS OFF, TIMING OFF) ")
	performanceTest.WriteString(inputedScript)
	tx, err := dbConn.Begin()
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	timeTest, err := tx.Query(performanceTest.String())
	if err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}
	defer timeTest.Close()
	for timeTest.Next() {
		var testResults string
		if err = timeTest.Scan(&testResults); err != nil {
			return JDGECODES.JudgeStatusUnknownError
		}
		if strings.Contains(testResults, "Execution time") {
			tempString := testResults
			tempString = strings.Replace(tempString, "Execution time: ", "", -1)
			tempString = strings.Replace(tempString, " ms", "", -1)
			timeFloat, err := strconv.ParseFloat(tempString, 64)
			if err != nil {
				return JDGECODES.JudgeStatusUnknownError
			}
			if timeFloat > float64(timeLimit) {
				return JDGECODES.JudgeStatusTimeLimitExceed
			}
		}
	}

	if err = tx.Rollback(); err != nil {
		return JDGECODES.JudgeStatusUnknownError
	}

	return JDGECODES.JudgeStatusCorrectAnswer
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
	RemoveEmpty(&admissionWordList)
	for i := 0; i < len(admissionWordList); i++ {
		reg, _ := regexp.Compile(`\b` + strings.ToLower(admissionWordList[i]) + `\b`)
		if !reg.MatchString(inputedScript) {
			return false
		}
	}
	return true
}

func RemoveEmpty(slice *[]string) {
	i := 0
	p := *slice
	for _, entry := range p {
		if strings.Trim(entry, " ") != "" {
			p[i] = entry
			i++
		}
	}
	*slice = p[0:i]
}
