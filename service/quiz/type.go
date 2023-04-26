package quiz

import (
	"context"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
)

type QuizService interface {
	ListAllQuiz(courseId string) ([]quiz.Quiz, error)
	GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error)

	DoTakeQuiz(ctx context.Context, quizId uuid.UUID, email string) (*quiz.QuizDetail, error)
	DoFinishQuiz(ctx context.Context, quizId uuid.UUID, email string, studentAnswer []quiz.Response) (*quiz.QuizTake, error)
	GetSolutionQuiz(ctx context.Context, quizId uuid.UUID, email string) (*quiz.QuizDetail, error)
}
