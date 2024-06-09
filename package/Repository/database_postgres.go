package repository

import (
	"fmt"
	"strconv"

	components "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Components"
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
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

	err := components.PrepareDB(r.db, database.FileName)
	if err != nil {
		return 0, err
	}

	var lastDbName string
	query := "SELECT datname FROM pg_database ORDER BY oid DESC LIMIT 1"
	row := r.db.QueryRow(query)
	if err = row.Scan(&lastDbName); err != nil {
		return 0, err
	}

	dbmsId, _ := strconv.Atoi(database.DbmsID)
	query = fmt.Sprintf("INSERT INTO %s (name, file_name, description, dbms_id) VALUES ($1, $2, $3, $4) RETURNING id", databasesTable)
	row = r.db.QueryRow(query, lastDbName, database.FileName, database.Description, dbmsId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DatabasePostgres) DeleteDataBaseByID(id int) (string, error) {
	var dbName string
	query := fmt.Sprintf("SELECT file_name FROM %s WHERE id=$1", databasesTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&dbName); err != nil {
		return "", err
	}
	query = fmt.Sprintf("DROP DATABASE %s", dbName)
	_, err := r.db.Exec(query)
	if err != nil {
		return dbName, err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", databasesTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		return dbName, err
	}
	return dbName, nil
}

func (r *DatabasePostgres) GetDatabaseByID(id int) (viewmodels.DatabaseViewModel, error) {
	var database viewmodels.DatabaseViewModel
	query := "SELECT db.id, db.name, db.file_name, db.description, dbms.name FROM databases db JOIN dbms ON db.dbms_id = dbms.id WHERE db.id = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&database.ID, &database.Name, &database.FileName, &database.Description, &database.DbmsName); err != nil {
		return database, err
	}
	return database, nil
}

func (r *DatabasePostgres) GetAllDatabases() ([]viewmodels.DatabaseViewModel, error) {
	var databases []viewmodels.DatabaseViewModel
	query := "SELECT db.id, db.name, db.file_name, db.description, dbms.name FROM databases db JOIN dbms ON db.dbms_id = dbms.id"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempDatabase viewmodels.DatabaseViewModel
		if err := rows.Scan(&tempDatabase.ID, &tempDatabase.Name, &tempDatabase.FileName, &tempDatabase.Description, &tempDatabase.DbmsName); err != nil {
			return databases, err
		}
		databases = append(databases, tempDatabase)
	}
	if err = rows.Err(); err != nil {
		return databases, err
	}
	return databases, nil
}
