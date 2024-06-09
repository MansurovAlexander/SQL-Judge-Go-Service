package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BannedWordToAssignPostgres struct {
	db *sqlx.DB
}

func NewBannedWordToAssignService(db *sqlx.DB) *BannedWordToAssignPostgres {
	return &BannedWordToAssignPostgres{db: db}
}

func (r *BannedWordToAssignPostgres) CreateBannedWordToAssign(assignID int, bannedWords, admissionWords map[string][]string) error {
	if len(bannedWords) > len(admissionWords) {
		for k, _ := range bannedWords {
			subtask, _ := strconv.Atoi(k)
			query := fmt.Sprintf("INSERT INTO %s (assign_id, banned_words, subtask_id) VALUES ($1, $2, $3)", bannedWordsToAssignTable)
			_, err := r.db.Exec(query, assignID, bannedWords[k][0], subtask)
			if err != nil {
				return err
			}
		}
		for k, _ := range admissionWords {
			subtask, _ := strconv.Atoi(k)
			var admissionWordsSB strings.Builder
			admissionWordsSB.WriteString("'")
			admissionWordsSB.WriteString(admissionWords[k][0])
			admissionWordsSB.WriteString("'")
			query := fmt.Sprintf("UPDATE %s SET admission_words = $1 WHERE assign_id = $2 AND subtask_id = $3", bannedWordsToAssignTable)
			_, err := r.db.Exec(query, admissionWordsSB.String(), assignID, subtask)
			if err != nil {
				return err
			}
		}
	} else {
		for k, _ := range admissionWords {
			subtask, _ := strconv.Atoi(k)
			query := fmt.Sprintf("INSERT INTO %s (assign_id, admission_words, subtask_id) VALUES ($1, $2, $3)", bannedWordsToAssignTable)
			_, err := r.db.Exec(query, assignID, admissionWords[k][0], subtask)
			if err != nil {
				return err
			}
		}
		for k, _ := range bannedWords {
			subtask, _ := strconv.Atoi(k)
			var bannedWordsSB strings.Builder
			bannedWordsSB.WriteString("'")
			bannedWordsSB.WriteString(bannedWords[k][0])
			bannedWordsSB.WriteString("'")
			query := fmt.Sprintf("UPDATE %s SET banned_words = $1 WHERE assign_id = $2 AND subtask_id = $3", bannedWordsToAssignTable)
			_, err := r.db.Exec(query, bannedWordsSB.String(), assignID, subtask)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *BannedWordToAssignPostgres) GetBannedWordByAssignID(id int) (map[string]string, error) {
	bannedWords := make(map[string]string)

	query := fmt.Sprintf("SELECT banned_words, subtask_id FROM %s WHERE assign_id=$1", bannedWordsToAssignTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempBannedWord string
		var subtask string
		if err := rows.Scan(&tempBannedWord, &subtask); err != nil {
			return nil, err
		}
		bannedWords[subtask] = tempBannedWord
	}
	return bannedWords, nil
}

func (r *BannedWordToAssignPostgres) GetAdmissionWordByAssignID(id int) (map[string]string, error) {
	admissionWords := make(map[string]string)

	query := fmt.Sprintf("SELECT admission_words, subtask_id FROM %s WHERE assign_id=$1", bannedWordsToAssignTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempBannedWord string
		var subtask string
		if err := rows.Scan(&tempBannedWord, &subtask); err != nil {
			return nil, err
		}
		admissionWords[subtask] = tempBannedWord
	}
	return admissionWords, nil
}

func (r *BannedWordToAssignPostgres) DropAllBannedWordsByAssignID(assignId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE assign_id=$1", bannedWordsToAssignTable)
	if _, err := r.db.Exec(query, assignId); err != nil {
		return err
	}
	return nil
}
