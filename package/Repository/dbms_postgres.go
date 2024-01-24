package repository

import (
	"fmt"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type DbmsPostgres struct {
	db *sqlx.DB
}

func NewDbmsService(db *sqlx.DB) *DbmsPostgres {
	return &DbmsPostgres{db: db}
}

func (r *DbmsPostgres) CreateDbms(dbms models.Dbms) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", dbmsTable)
	row := r.db.QueryRow(query, dbms.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DbmsPostgres) GetDbmsByID(id int) (models.Dbms, error) {
	var dbms models.Dbms
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=($1)", dbmsTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&dbms.ID, &dbms.Name); err != nil {
		return dbms, err
	}
	return dbms, nil
}

func (r *DbmsPostgres) GetAllDbms() ([]models.Dbms, error) {
	var dbms []models.Dbms
	query := fmt.Sprintf("SELECT * FROM %s", dbmsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempDbms models.Dbms
		if err := rows.Scan(&tempDbms.ID, &tempDbms.Name); err != nil {
			return dbms, err
		}
		dbms = append(dbms, tempDbms)
	}
	if err = rows.Err(); err != nil {
		return dbms, err
	}
	return dbms, nil
}
