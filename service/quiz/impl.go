package quiz

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/provider/storage"
	quizRepo "gitlab.informatika.org/ocw/ocw-backend/repository/quiz"
)

type QuizServiceImpl struct {
	quizRepo.QuizRepository
	storage.Storage
}

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

func (q QuizServiceImpl) GetSolutionQuiz(ctx context.Context, quizId uuid.UUID, email string) (*quiz.QuizDetail, error) {
	result, err := q.getQuizDetail(ctx, quizId)

	if err != nil {
		return nil, err
	}

	_, err = q.GetLastTake(quizId, email)

	if err != nil {
		return nil, err
	}

	taken, err := q.IsActiveTake(quizId, email)

	if err != nil {
		return nil, err
	}

	if taken {
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

	err = q.QuizRepository.UpdateScore(data.Id, int(score))

	if err != nil {
		return nil, err
	}

	return data, nil
}
