package repositories

import (
	"database/sql"
	"github.com/gin/demo/internal/domain"
	"time"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) SaveMessage(eventType, payload, status string) (int64, error) {
	result, err := r.DB.Exec(`
		INSERT INTO messages (event_type, payload, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`,
		eventType, payload, status, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *MessageRepository) UpdateMessageStatus(id int64, status string) error {
	_, err := r.DB.Exec(`
		UPDATE messages SET status = ?, updated_at = ? WHERE id = ?`,
		status, time.Now(), id)
	return err
}

func (r *MessageRepository) GetPendingMessages() ([]domain.Message, error) {
	rows, err := r.DB.Query(`SELECT id, event_type, payload FROM messages WHERE status = ?`, domain.StatusPending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var msg domain.Message
		if err := rows.Scan(&msg.ID, &msg.EventType, &msg.Payload); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
