package service

import (
	"github.com/gin/demo/internal/application"
	"github.com/gin/demo/internal/infrastructure/repositories"
)

type MessageService struct {
	dispatcher        *application.EventDispatcher
	messageRepository *repositories.MessageRepository
}

func NewMessageService(dispatcher *application.EventDispatcher, messageRepo *repositories.MessageRepository) *MessageService {
	return &MessageService{
		dispatcher:        dispatcher,
		messageRepository: messageRepo,
	}
}

//func (s *MessageService) RetryPendingEvents() {
//	for {
//		messages, err := s.messageRepository.GetPendingMessages()
//		if err != nil {
//			fmt.Println("Error fetching pending messages:", err)
//			time.Sleep(5 * time.Second)
//			continue
//		}
//
//		for _, msg := range messages {
//			var event domain.AccountEvent
//			json.Unmarshal([]byte(msg.Payload), &event)
//			s.dispatcher.Dispatch("",event)
//		}
//		time.Sleep(10 * time.Second)
//	}
//}
