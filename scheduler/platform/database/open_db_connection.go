package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.MessageQueries // load queries from Message model
	*queries.ConfigQueries  // load queries from Config model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *sqlx.DB
		err error
	)

	// Define a new Database connection.
	db, err = PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		MessageQueries: &queries.MessageQueries{DB: db}, // from Book model
		ConfigQueries:  &queries.ConfigQueries{DB: db},  // from Config model
	}, nil
}
