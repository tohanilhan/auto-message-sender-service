package models

// Message struct to describe the message model.
type Message struct {
	ID         string `json:"-" db:"id"`
	Content    string `json:"content" db:"content"`
	To         string `json:"to" db:"recipient_phone_number"`
	SendStatus string `json:"-" db:"send_status"`
}

// SendMessageResponse struct to describe the response body for sending a message.
type SendMessageResponse struct {
	MessageID string `json:"messageId"`
	Message   string `json:"message"`
}
