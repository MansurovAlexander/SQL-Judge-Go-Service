package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BannedWordToAssignPostgres struct {
	db *sqlx.DB
}

func NewBannedWordToAssignService(db *sqlx.DB) *BannedWordToAssignPostgres {
	return &BannedWordToAssignPostgres{db: db}
}

func (r *BannedWordToAssignPostgres) CreateBannedWordToAssign(assignID int, bannedWordID int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (assign_id, banned_word_id) VALUES ($1, $2) RETURNING id", bannedWordsToAssignTable)
	row := r.db.QueryRow(query, assignID, bannedWordID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BannedWordToAssignPostgres) GetBannedWordByAssignID(id int) ([]int, error) {
	var (
		bannedWordsID            []int
		tempAssignId             int
		tempBannedWordToAssignID int
	)
	query := fmt.Sprintf("SELECT banned_word_id FROM %s WHERE assign_id=$1", bannedWordsToAssignTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bannedWordID int
		if err := rows.Scan(&tempBannedWordToAssignID, &tempAssignId, &bannedWordID); err != nil {
			return bannedWordsID, err
		}
		bannedWordsID = append(bannedWordsID, bannedWordID)
	}
	if err = rows.Err(); err != nil {
		return bannedWordsID, err
	}
	return bannedWordsID, nil
}
