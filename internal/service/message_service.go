package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin/demo/internal/domain"
	"github.com/gin/demo/internal/infrastructure/repositories"
	"time"
)

type MessageService struct {
	dispatcher        *EventDispatcher
	messageRepository *repositories.MessageRepository
}

func NewMessageService(dispatcher *EventDispatcher, messageRepo *repositories.MessageRepository) *MessageService {
	return &MessageService{
		dispatcher:        dispatcher,
		messageRepository: messageRepo,
	}
}

func (s *MessageService) RetryPendingEvents() {
	for {
		messages, err := s.messageRepository.GetPendingMessages()
		if err != nil {
			fmt.Println("Error fetching pending messages:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, msg := range messages {
			var event domain.CartEvent
			json.Unmarshal([]byte(msg.Payload), &event)
			s.dispatcher.Dispatch(event)
		}
		time.Sleep(10 * time.Second)
	}
}
