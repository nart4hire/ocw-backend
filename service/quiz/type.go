package quiz

import (
	"context"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	model "gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
)

type QuizService interface {
	ListAllQuiz(courseId string) ([]quiz.Quiz, error)
	GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error)

	DoTakeQuiz(ctx context.Context, quizId uuid.UUID, email string) (*quiz.QuizDetail, error)
	DoFinishQuiz(ctx context.Context, quizId uuid.UUID, email string, studentAnswer []quiz.Response) (*quiz.QuizTake, error)
	GetSolutionQuiz(ctx context.Context, quizId uuid.UUID, user token.UserClaim) (*quiz.QuizDetail, error)
	isQuizContributor(courseId string, email string) error
	NewQuiz(payload model.AddQuizRequestPayload) (*model.LinkResponse, error)
	GetQuiz(payload model.UpdateQuizRequestPayload) (*model.LinkResponse, error)
	DeleteQuiz(payload model.DeleteRequestPayload) error
}
