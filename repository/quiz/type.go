package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
)

type QuizRepository interface {
	GetQuizes(courseId string) ([]quiz.Quiz, error)
	GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error)
	UpdateScore(quizId uuid.UUID, score int) error
	NewTake(quizId uuid.UUID, userEmail string) (uuid.UUID, error)
	IsUserContributor(id string, email string) (bool, error)
	NewQuiz(quiz quiz.Quiz) error
	UpdateQuiz(quiz quiz.Quiz) error
	Delete(quizId uuid.UUID) error
	IsActiveTake(quizId uuid.UUID, userEmail string) (bool, error)
	GetAllTake(quizId uuid.UUID, userEmail string) ([]quiz.QuizTake, error)
	GetLastTake(quizId uuid.UUID, userEmail string) (*quiz.QuizTake, error)
}
