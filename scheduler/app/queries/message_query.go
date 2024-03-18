package queries

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/models"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/repository"
)

// MessageQueries struct for queries from Book model.
type MessageQueries struct {
	*sqlx.DB
}

var (
	getUnsentMessagesQuery       string = "SELECT id,content,send_status,recipient_phone_number FROM message_sender_service_schema.messages where send_status = $1 ORDER BY creation_time DESC LIMIT 2;"
	updateMessageSendStatusQuery string = `UPDATE message_sender_service_schema.messages SET send_status = $1, send_time=$2::timestamp, messageid=$3 WHERE id = $4`
)

// GetMessages method for getting all books.
func (q *MessageQueries) GetUnsendedMessages() ([]models.Message, error) {
	// Define books variable.
	messages := []models.Message{}

	// Prepare query.
	stmt, err := q.Preparex(getUnsentMessagesQuery)
	if err != nil {
		// Return empty object and error.
		return messages, err
	}

	err = stmt.Select(&messages, repository.NotSentStatus)
	if err != nil {
		// Return empty object and error.
		return messages, err
	}

	// Return query result.
	return messages, nil
}

func (q *MessageQueries) UpdateMessageStatus(resp *models.SendMessageResponse, id string, timestamp time.Time) error {

	// Prepare query.
	stmt, err := q.Preparex(updateMessageSendStatusQuery)
	if err != nil {
		// Return empty object and error.
		return err
	}

	_, err = stmt.Exec(repository.SentStatus, timestamp, resp.MessageID, id)
	if err != nil {
		// Return empty object and error.
		return err
	}

	// Return query result.
	return nil
}
