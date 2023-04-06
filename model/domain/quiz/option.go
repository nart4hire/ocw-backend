package quiz

import "github.com/google/uuid"

type Option struct {
	Id       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	IsAnswer bool      `json:"is_answer"`
}
