package quiz

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
	userDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	model "gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/provider/storage"
	quizRepo "gitlab.informatika.org/ocw/ocw-backend/repository/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	tokenUtil "gitlab.informatika.org/ocw/ocw-backend/utils/token"
	"gorm.io/gorm"
)

type QuizServiceImpl struct {
	quizRepo.QuizRepository
	storage.Storage
	tokenUtil.TokenUtil
	logger.Logger
	*env.Environment
}

// TODO: should be for admins, make ones for users which doesnt expose minio link
func (q QuizServiceImpl) ListAllQuiz(courseId string) ([]quiz.Quiz, error) {
	return q.QuizRepository.GetQuizes(courseId)
}

func (q QuizServiceImpl) GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error) {
	return q.QuizRepository.GetQuizDetail(quizId)
}

func (q QuizServiceImpl) getQuizDetail(ctx context.Context, quizId uuid.UUID) (*quiz.QuizDetail, error) {
	detail, err := q.QuizRepository.GetQuizDetail(quizId)

	if err != nil {
		return nil, err
	}

	payload, err := q.Storage.Get(ctx, detail.QuizPath)

	if err != nil {
		return nil, err
	}

	result := &quiz.QuizDetail{}

	decoder := json.NewDecoder(bytes.NewReader(payload))
	err = decoder.Decode(result)

	return result, err
}

func (q QuizServiceImpl) DoTakeQuiz(ctx context.Context, quizId uuid.UUID, email string) (*quiz.QuizDetail, error) {
	result, err := q.getQuizDetail(ctx, quizId)

	if err != nil {
		return nil, err
	}

	taken, err := q.IsActiveTake(quizId, email)

	if err != nil {
		return nil, err
	}

	if !taken {
		_, err = q.NewTake(quizId, email)

		if err != nil {
			return nil, err
		}
	}

	for i := range result.Problems {
		for j := range result.Problems[i].Answer {
			result.Problems[i].Answer[j].IsSolution = nil
		}
	}

	return result, nil
}

func (q QuizServiceImpl) GetSolutionQuiz(ctx context.Context, quizId uuid.UUID, user token.UserClaim) (*quiz.QuizDetail, error) {
	result, err := q.getQuizDetail(ctx, quizId)

	if err != nil {
		return nil, err
	}

	last, err := q.GetLastTake(quizId, user.Email)

	if err != nil {
		return nil, err
	}

	if last == nil && user.Role == userDomain.Student {
		return nil, web.NewResponseError("user is not allow to access this data", "ERR_NOT_ALLOWED")
	}

	taken, err := q.IsActiveTake(quizId, user.Email)

	if err != nil {
		return nil, err
	}

	if taken && user.Role == userDomain.Student {
		return nil, web.NewResponseError("user is not allow to access this data", "ERR_NOT_ALLOWED")
	}

	return result, nil
}

func (q QuizServiceImpl) checkAnswer(detail *quiz.QuizDetail, studentAnswer []quiz.Response) float64 {
	answerDict := map[uuid.UUID][]uuid.UUID{}
	totalProblem := len(detail.Problems)

	for _, problem := range detail.Problems {
		correctAnswerId := []uuid.UUID{}
		for _, answer := range problem.Answer {
			if *answer.IsSolution {
				correctAnswerId = append(correctAnswerId, answer.Id)
			}
		}

		answerDict[problem.Id] = correctAnswerId
	}

	correctAnswer := 0
	for _, responseItem := range studentAnswer {
		numCorrect := 0

		for _, correctId := range answerDict[responseItem.ProblemId] {
			if responseItem.AnswerId == correctId {
				numCorrect++
			}
		}

		if numCorrect == len(answerDict[responseItem.ProblemId]) {
			correctAnswer++
		}
	}

	return float64(correctAnswer) / float64(totalProblem) * 100
}

func (q QuizServiceImpl) DoFinishQuiz(ctx context.Context, quizId uuid.UUID, email string, studentAnswer []quiz.Response) (*quiz.QuizTake, error) {
	taken, err := q.IsActiveTake(quizId, email)

	if err != nil {
		return nil, err
	}

	if !taken {
		return nil, web.NewResponseError("user not yet do take the quiz", "NOT_TAKEN_QUIZ_YET")
	}

	result, err := q.getQuizDetail(ctx, quizId)

	if err != nil {
		return nil, err
	}

	score := q.checkAnswer(result, studentAnswer)

	data, err := q.QuizRepository.GetLastTake(quizId, email)
	data.IsFinished = true
	data.Score = int(score)

	if err != nil {
		return nil, err
	}

	err = q.QuizRepository.UpdateScore(email, quizId, int(score))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q QuizServiceImpl) isQuizContributor(courseId string, email string) error {
	_, err := q.QuizRepository.IsUserContributor(courseId, email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseError("course and user combination not found", "NOT_OWNER")
		}

		return err
	}

	return nil
}

func (q QuizServiceImpl) NewQuiz(payload model.AddQuizRequestPayload) (*model.LinkResponse, error) {
	// Validate Role
	claim, err := q.TokenUtil.Validate(payload.AddQuizToken, token.Access)

	// Invalid Token
	if err != nil {
		return &model.LinkResponse{}, web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == userDomain.Student {
		return &model.LinkResponse{}, web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// Validate Ownership
	if err := q.isQuizContributor(payload.CourseID, claim.Email); err != nil {
		return &model.LinkResponse{}, err
	}

	path := fmt.Sprintf("%s/%s.json", q.BucketQuizBasePath, strings.ReplaceAll(uuid.New().String(), "-", ""))
	uploadLink, err := q.Storage.CreatePutSignedLink(context.Background(), path)

	if err != nil {
		q.Logger.Error("Some error happened when generate link")
		q.Logger.Error(err.Error())
		return &model.LinkResponse{}, err
	}

	err = q.QuizRepository.NewQuiz(quiz.Quiz{
		Id:           uuid.New(),
		Name:         payload.Name,
		CourseId:     payload.CourseID,
		CreatorEmail: claim.Email,
		QuizPath:     path,
	})

	if err != nil {
		q.Logger.Error("Some error happened when insert to repository")
		q.Logger.Error(err.Error())
		return &model.LinkResponse{}, err
	}

	return &model.LinkResponse{UploadLink: uploadLink}, nil
}

func (q QuizServiceImpl) UpdateQuiz(payload model.UpdateQuizRequestPayload) (*model.LinkResponse, error) {
	// Validate Role
	claim, err := q.TokenUtil.Validate(payload.UpdateQuizToken, token.Access)

	// Invalid Token
	if err != nil {
		return &model.LinkResponse{}, web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == userDomain.Student {
		return &model.LinkResponse{}, web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// Get Quiz Detail
	prev, err := q.QuizRepository.GetQuizDetail(payload.ID)

	if err != nil {
		return &model.LinkResponse{}, err
	}

	// Validate Ownership
	if err := q.isQuizContributor(prev.CourseId, claim.Email); err != nil {
		return &model.LinkResponse{}, err
	}

	path := fmt.Sprintf("%s/%s.json", q.BucketQuizBasePath, strings.ReplaceAll(uuid.New().String(), "-", ""))
	uploadLink, err := q.Storage.CreatePutSignedLink(context.Background(), path)

	if err != nil {
		q.Logger.Error("Some error happened when generate link")
		q.Logger.Error(err.Error())
		return &model.LinkResponse{}, err
	}

	err = q.QuizRepository.UpdateQuiz(quiz.Quiz{
		Id:           prev.Id,
		Name:         prev.Name,
		CourseId:     prev.CourseId,
		CreatorEmail: prev.CreatorEmail,
		QuizPath:     path,
	})

	if err != nil {
		q.Logger.Error("Some error happened when inserting new link")
		q.Logger.Error(err.Error())
		return &model.LinkResponse{}, err
	}

	return &model.LinkResponse{UploadLink: uploadLink}, nil
}

func (q QuizServiceImpl) GetQuizLink(payload model.GetRequestPayload) (*model.PathResponse, error) {
	// Validate Role
	claim, err := q.TokenUtil.Validate(payload.GetToken, token.Access)

	// Invalid Token
	if err != nil {
		return &model.PathResponse{}, web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == userDomain.Student {
		return &model.PathResponse{}, web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// Get Quiz Detail
	quiz, err := q.QuizRepository.GetQuizDetail(payload.ID)

	if err != nil {
		return &model.PathResponse{}, err
	}

	// Validate Ownership
	if err := q.isQuizContributor(quiz.CourseId, claim.Email); err != nil {
		return &model.PathResponse{}, err
	}

	return &model.PathResponse{Path: quiz.QuizPath}, nil
}

func (q QuizServiceImpl) DeleteQuiz(payload model.DeleteRequestPayload) error {
	// Validate Role
	claim, err := q.TokenUtil.Validate(payload.DeleteToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == userDomain.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// Get Quiz Detail
	quiz, err := q.QuizRepository.GetQuizDetail(payload.ID)

	if err != nil {
		return err
	}

	// Validate Ownership
	if err := q.isQuizContributor(quiz.CourseId, claim.Email); err != nil {
		return err
	}

	return q.QuizRepository.Delete(payload.ID)
}
