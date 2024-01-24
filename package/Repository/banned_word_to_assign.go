package repository

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type BannedWordToAssignPostgres struct {
	db *sqlx.DB
}

func NewBannedWordToAssignService(db *sqlx.DB) *BannedWordToAssignPostgres {
	return &BannedWordToAssignPostgres{db: db}
}

func (r *BannedWordToAssignPostgres) CreateBannedWordToAssign(bannedwordtoassign models.BannedWordToAssign) (big.Int, error) {
	var id big.Int
	query := "INSERT INTO $1 (assign_id, banned_word_id) VALUES ($2, $3) RETURNING id"
	row := r.db.QueryRow(query, bannedWordsToAssignTable, bannedwordtoassign.AssignID, bannedwordtoassign.BannedWordID)
	if err := row.Scan(&id); err != nil {
		return *big.NewInt(0), err
	}
	return id, nil
}

func (r *BannedWordToAssignPostgres) GetBannedWordByAssignID(id big.Int) ([]int, error) {
	var (
		bannedWordsID            []int
		tempAssignId             big.Int
		tempBannedWordToAssignID big.Int
	)
	query := "SELECT banned_word_id FROM $1 WHERE assign_id=$2"
	rows, err := r.db.Query(query, bannedWordsToAssignTable)
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
