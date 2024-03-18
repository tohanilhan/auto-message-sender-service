package message

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/models"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/repository"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/utils"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/platform/cache"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/platform/database"
	"github.com/useinsider/go-pkg/insrequester"
)

type MessageSender struct {
	WebhookURL string
	SentTime   time.Time
}

func NewMessageSender(url string) *MessageSender {
	return &MessageSender{
		WebhookURL: url,
	}
}

func (s *MessageSender) Send(messages []models.Message) error {
	// Send message concurrently

	ch := make(chan string, len(messages))

	for _, message := range messages {
		go func(message models.Message) {

			// handle request
			bodyByte, err := s.handleRequest(message)
			if err != nil {
				// Send error message to channel
				ch <- err.Error()
			} else {
				// Send message to channel
				ch <- string(bodyByte)
			}
		}(message)
	}

	// Handle response
	err := s.handleResponse(ch, messages)
	if err != nil {
		// log error
		log.Println(err)
	}

	return nil
}

func (s *MessageSender) handleRequest(message models.Message) ([]byte, error) {

	// Create a new requester
	requester := insrequester.NewRequester().Load()

	// Set headers
	headers := map[string]interface{}{
		"Content-Type":   "application/json",
		"x-ins-auth-key": utils.EnvironmentVars.WebhookApiKey,
	}

	// Marshal the message body to JSON
	body, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	// Create a new request entity
	requestEntity := insrequester.RequestEntity{
		Endpoint: s.WebhookURL,
		Headers:  insrequester.Headers{headers},
		Body:     body,
	}

	// Send the request
	resp, err := requester.Post(requestEntity)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is not 202
	if resp.StatusCode != 202 {
		return nil, fmt.Errorf("error code:%s", resp.Status)
	}

	// read response body
	defer resp.Body.Close()

	boyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return boyByte, nil
}

func (s *MessageSender) handleResponse(ch chan string, message []models.Message) error {

	// Wait for all messages to be sent
	for _, message := range message {
		response := <-ch
		// Handle response
		if !strings.Contains(response, "error") {
			responseStruct, err := s.unmarshalResponse(response)
			if err != nil {
				// log error
				log.Println(err)
				continue
			}

			timestamp := time.Now().UTC()
			s.SentTime = timestamp

			// Update db and cache response concurrently
			wg := sync.WaitGroup{}
			wg.Add(2)

			go func() {
				defer wg.Done()
				// Update db
				err = s.updateMessageOnDatabase(responseStruct, message.ID)
				if err != nil {
					// log error
					log.Println(err)
				}
			}()

			go func() {
				defer wg.Done()
				// cache response
				err = s.storeResponseToCache(responseStruct)
				if err != nil {
					// log error
					log.Println(err)
				}
			}()

			wg.Wait()
			// log response
			return nil
		}
	}

	return fmt.Errorf(<-ch)
}

func (s *MessageSender) unmarshalResponse(message string) (*models.SendMessageResponse, error) {
	// unmashal response
	responseStruct := &models.SendMessageResponse{}
	err := json.Unmarshal([]byte(message), &responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}

func (s *MessageSender) updateMessageOnDatabase(resp *models.SendMessageResponse, id string) error {
	// update db

	db, err := database.OpenDBConnection()
	if err != nil {
		// log error
		log.Println(err)
		return err
	}

	// update message status
	err = db.UpdateMessageStatus(resp, id, s.SentTime)
	if err != nil {
		// log error
		log.Println(err)
		return err
	}

	return nil

}

// storeResponseToCache caches the response of the message
func (s *MessageSender) storeResponseToCache(response *models.SendMessageResponse) error {
	// cache response

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		return err
	}

	// value := map[string]interface{}{
	// 	"messageId": response.MessageID,
	// 	"sent_time": s.SentTime.Format("2006-01-02 15:04:05"),
	// }

	// // stringfy value
	// byteValue, err := json.Marshal(value)
	// if err != nil {
	// 	return err
	// }

	// _, err = connRedis.RPush(context.Background(), "sent_messages", byteValue).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	err = connRedis.Set(context.Background(), "sent_message:"+response.MessageID, s.SentTime.Format("2006-01-02 15:04:05"), repository.MessageTTL).Err()
	if err != nil {
		return err
	}

	return nil
}
