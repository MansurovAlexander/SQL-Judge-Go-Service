package repository

import (
	"fmt"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type BannedWordPostgres struct {
	db *sqlx.DB
}

func NewBannedWordService(db *sqlx.DB) *BannedWordPostgres {
	return &BannedWordPostgres{db: db}
}

func (r *BannedWordPostgres) CreateBannedWord(bannedword models.BannedWord) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (banned_word) VALUES ($1) RETURNING id", bannedWordsTable)
	row := r.db.QueryRow(query, bannedword.BannedWord)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BannedWordPostgres) GetBannedWordByID(id int) (models.BannedWord, error) {
	var bannedword models.BannedWord
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", bannedWordsTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&bannedword.ID, &bannedword.BannedWord); err != nil {
		return bannedword, err
	}
	return bannedword, nil
}

func (r *BannedWordPostgres) GetAllBannedWords() ([]models.BannedWord, error) {
	var bannedwords []models.BannedWord
	query := fmt.Sprintf("SELECT * FROM %s", bannedWordsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempBannedWord models.BannedWord
		if err := rows.Scan(&tempBannedWord.ID, &tempBannedWord.BannedWord); err != nil {
			return bannedwords, err
		}
		bannedwords = append(bannedwords, tempBannedWord)
	}
	if err = rows.Err(); err != nil {
		return bannedwords, err
	}
	return bannedwords, nil
}
