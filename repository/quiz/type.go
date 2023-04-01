package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
)

type QuizRepository interface {
	GetQuizes(courseId string) ([]quiz.Quiz, error)
	GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error)
	UpdateScore(takeId uuid.UUID, score int) error
	NewTake(quizId uuid.UUID, userEmail string) (uuid.UUID, error)
	IsActiveTake(quizId uuid.UUID, userEmail string) (bool, error)
	GetAllTake(quizId uuid.UUID, userEmail string) ([]quiz.QuizTake, error)
}
