package scheduler

import (
	"testing"

	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/message"
)

// TestCreateInstance tests the CreateInstance function. It should return a single instance of Scheduler.
func TestCreateInstance(t *testing.T) {
	// Mock MessageSender
	mockMessageSender := &message.MessageSender{} // or create a mock if necessary

	// Test creating the first instance
	scheduler1 := CreateInstance(mockMessageSender)
	if scheduler1 == nil {
		t.Error("First instance creation failed")
	}

	// Test creating subsequent instances
	for i := 0; i < 3; i++ {
		scheduler := CreateInstance(mockMessageSender)
		if scheduler != scheduler1 {
			t.Error("Subsequent instance creation failed, instances are not equal")
		}
	}
}
