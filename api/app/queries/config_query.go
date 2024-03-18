package queries

import (
	"github.com/jmoiron/sqlx"
)

// MessageQueries struct for queries from Message model.
type MessageQueries struct {
	*sqlx.DB
}

var (
	updateAutoSendingConfig = `UPDATE message_sender_service_schema.config SET status=$1 where name='AUTO-SEND-FEATURE';`
)

// UpdateAutoSendingConfig method for updating auto sending config.
func (q *MessageQueries) UpdateAutoSendingConfig(status string) error {
	// Prepare query.
	stmt, err := q.Preparex(updateAutoSendingConfig)
	if err != nil {
		// Return empty object and error.
		return err
	}

	_, err = stmt.Exec(status)
	if err != nil {
		// Return empty object and error.
		return err
	}

	// Return query result.
	return nil
}
