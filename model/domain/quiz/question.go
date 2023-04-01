package quiz

import "github.com/google/uuid"

type Question struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"string"`
	Options     []Option  `json:"options"`
}
