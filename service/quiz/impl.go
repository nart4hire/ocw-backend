package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
	quizRepo "gitlab.informatika.org/ocw/ocw-backend/repository/quiz"
)

type QuizServiceImpl struct {
	quizRepo.QuizRepository
}

func (q QuizServiceImpl) ListAllQuiz(courseId string) ([]quiz.Quiz, error) {
	return q.QuizRepository.GetQuizes(courseId)
}

func (q QuizServiceImpl) GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error) {
	return q.QuizRepository.GetQuizDetail(quizId)
}
