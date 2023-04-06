package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
)

type QuizService interface {
	ListAllQuiz(courseId string) ([]quiz.Quiz, error)
	GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error)
}
