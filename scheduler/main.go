package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	scheduler "github.com/tohanilhan/auto-message-sender-service/scheduler/app/cron-job"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/message"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/utils"
)

func init() {
	// Load .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse environment variables.
	err = utils.ParseEnvironmentVariables()
	if err != nil {
		log.Fatal("Error parsing environment variables")
	}

	fmt.Printf("%+v\n", utils.EnvironmentVars)

}

func main() {
	// Initialize message sender
	messageSender := message.NewMessageSender(utils.EnvironmentVars.WebhookApiUrl)

	// Initialize scheduler using Singleton pattern
	cronScheduler := scheduler.CreateInstance(messageSender)

	cronScheduler.StartScheduler()
}
