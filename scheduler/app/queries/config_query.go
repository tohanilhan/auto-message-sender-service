package queries

import "github.com/jmoiron/sqlx"

// MessageQueries struct for queries from Book model.
type ConfigQueries struct {
	*sqlx.DB
}

var (
	getConfigQuery = `SELECT status FROM message_sender_service_schema.config;`
)

func (q *ConfigQueries) GetConfig() (string, error) {
	// Define books variable.
	var config string

	// Prepare query.
	stmt, err := q.Preparex(getConfigQuery)
	if err != nil {
		// Return empty object and error.
		return config, err
	}

	err = stmt.Get(&config)
	if err != nil {
		// Return empty object and error.
		return config, err
	}

	// Return query result.
	return config, nil
}
