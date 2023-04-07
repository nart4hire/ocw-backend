package quiz

import "gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"

type FinishQuizPayload struct {
	Data []quiz.Response `json:"data"`
}
