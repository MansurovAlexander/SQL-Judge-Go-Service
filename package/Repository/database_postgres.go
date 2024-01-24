package repository

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type DatabasePostgres struct {
	db *sqlx.DB
}

func NewDatabaseService(db *sqlx.DB) *DatabasePostgres {
	return &DatabasePostgres{db: db}
}

func (r *DatabasePostgres) CreateDatabase(database models.Database) (int, error) {
	var id int
	query := "INSERT INTO $1 (name, description, creation_script, dbms_id) VALUES ($2, $3, $4, $5) RETURNING id"
	row := r.db.QueryRow(query, databasesTable, database.Name, database.Description, database.CreationScript, database.DbmsID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DatabasePostgres) GetDatabaseByID(id int) (models.Database, error) {
	var database models.Database
	query := "SELECT * FROM $1 WHERE id=$2"
	row := r.db.QueryRow(query, databasesTable, id)
	if err := row.Scan(&database.ID, &database.Name, &database.Description, &database.CreationScript, &database.DbmsID); err != nil {
		return database, err
	}
	return database, nil
}

func (r *DatabasePostgres) GetAllDatabases() ([]models.Database, error) {
	var databases []models.Database
	query := "SELECT * FROM $1"
	rows, err := r.db.Query(query, databasesTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempDatabase models.Database
		if err := rows.Scan(&tempDatabase.ID, &tempDatabase.Name, &tempDatabase.Description, &tempDatabase.CreationScript,
			tempDatabase.DbmsID); err != nil {
			return databases, err
		}
		databases = append(databases, tempDatabase)
	}
	if err = rows.Err(); err != nil {
		return databases, err
	}
	return databases, nil
}

