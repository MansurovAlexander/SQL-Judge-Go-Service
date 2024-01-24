package repository

import (
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
	query := "INSERT INTO $1 (banned_word) VALUES ($2) RETURNING id"
	row := r.db.QueryRow(query, bannedWordsTable, bannedword.BannedWord)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BannedWordPostgres) GetBannedWordByID(id int) (models.BannedWord, error) {
	var bannedword models.BannedWord
	query := "SELECT * FROM $1 WHERE id=$2"
	row := r.db.QueryRow(query, bannedWordsTable, id)
	if err := row.Scan(&bannedword.ID, &bannedword.BannedWord); err != nil {
		return bannedword, err
	}
	return bannedword, nil
}

func (r *BannedWordPostgres) GetAllBannedWords() ([]models.BannedWord, error) {
	var bannedwords []models.BannedWord
	query := "SELECT * FROM $1"
	rows, err := r.db.Query(query, bannedWordsTable)
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

