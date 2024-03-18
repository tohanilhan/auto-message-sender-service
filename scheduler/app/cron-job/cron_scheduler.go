package scheduler

import (
	"fmt"
	"log"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/message"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/repository"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/utils"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/platform/database"
)

var MessageScheduler *Scheduler

var lock = &sync.Mutex{}

// Scheduler is responsible for triggering automatic message sending.
type Scheduler struct {
	messageSender *message.MessageSender
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler(messageSender *message.MessageSender) *Scheduler {
	return &Scheduler{
		messageSender: messageSender,
	}
}

// CreateInstance creates a single instance of Scheduler. Used Singleton pattern to avoid creating multiple instances of the scheduler
func CreateInstance(messageSender *message.MessageSender) *Scheduler {
	if MessageScheduler == nil {
		lock.Lock()
		defer lock.Unlock()
		if MessageScheduler == nil {
			MessageScheduler = NewScheduler(messageSender)
		}
	}
	return MessageScheduler
}

// StartScheduler starts the scheduler to trigger automatic message sending every configured interval.
func (s *Scheduler) StartScheduler() {
	fmt.Println("Starting scheduler")
	// create cron job
	c := cron.New(cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))

	c.AddFunc(utils.EnvironmentVars.CronJobSchedule, func() {

		db, err := database.OpenDBConnection()
		if err != nil {
			log.Println(err)
		}

		// check if sending is enabled
		config, err := s.checkAutoSendingConfig(db)
		if err != nil {
			log.Println(err)
		}

		if config == repository.AutoSendOn {
			//get all messages
			messages, err := db.GetUnsendedMessages()
			if err != nil {
				log.Println(err)
			}

			if len(messages) == 0 {
				return
			}

			// Trigger message sending process
			err = s.messageSender.Send(messages)
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("Auto sending is off")
		}

	})

	c.Start()

	// block the main thread
	select {}
}

// CheckAutoSendingConfig enables or disables the scheduler.
func (s *Scheduler) checkAutoSendingConfig(db *database.Queries) (string, error) {
	//get config
	config, err := db.GetConfig()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return config, nil
}
