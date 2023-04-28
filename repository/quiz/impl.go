package quiz

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type QuizRepositoryImpl struct {
	db *gorm.DB
}

func New(
	db db.Database,
) *QuizRepositoryImpl {
	return &QuizRepositoryImpl{db.Connect()}
}

func (q *QuizRepositoryImpl) GetQuizes(courseId string) ([]quiz.Quiz, error) {
	result := &[]quiz.Quiz{}
	err := q.db.Where("course_id = ?", courseId).Find(result).Error

	return *result, err
}

func (q *QuizRepositoryImpl) GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error) {
	result := &quiz.Quiz{}
	err := q.db.Where("id = ?", quizId).First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, web.NewResponseError("Record not found", "ERR_NOT_FOUND")
	}

	return result, nil
}

func (q *QuizRepositoryImpl) UpdateScore(takeId uuid.UUID, score int) error {
	return q.db.
		Model(&quiz.QuizTake{}).
		Update("score", score).
		Update("is_finished", true).
		Where("id = ?", takeId).Error
}

func (q *QuizRepositoryImpl) NewTake(quizId uuid.UUID, userEmail string) (uuid.UUID, error) {
	id := uuid.New()
	err := q.db.Create(
		&quiz.QuizTake{
			Id:         id,
			Email:      userEmail,
			StartTime:  time.Now(),
			QuizId:     quizId,
			IsFinished: false,
			Score:      0,
		},
	).Error

	return id, err
}

func (q *QuizRepositoryImpl) IsUserContributor(id string, email string) (bool, error) {
	err := q.db.Where("id = ? AND email = ?", id, email).Find(&course.Course{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func(q *QuizRepositoryImpl) NewQuiz(quiz quiz.Quiz) error {
	return q.db.Create(&quiz).Error
}

func(q *QuizRepositoryImpl) GetQuizPath(quizId uuid.UUID) (string, error) {
	result := quiz.Quiz{}
	err := q.db.Where("id = ?", quizId).Find(&result).Error

	if err != nil {
		return "", err
	}

	return result.QuizPath, nil
}

func(q *QuizRepositoryImpl) Delete(quizId uuid.UUID) error {
	return q.db.Delete(&quiz.Quiz{}, quizId).Error
}

func (q *QuizRepositoryImpl) IsActiveTake(quizId uuid.UUID, userEmail string) (bool, error) {
	result := struct{ cnt int }{}
	err := q.db.
		Select("COUNT(*) as cnt").
		Where("quiz_id = ? AND email = ? AND is_finished = false", quizId, userEmail).
		Find(result).
		Error

	if err != nil {
		return false, nil
	}

	return result.cnt > 0, nil
}

func (q *QuizRepositoryImpl) GetAllTake(quizId uuid.UUID, userEmail string) ([]quiz.QuizTake, error) {
	result := []quiz.QuizTake{}
	err := q.db.
		Where("quiz_id = ? AND email = ?", quizId, userEmail).
		Find(result).Error

	return result, err
}

func (q *QuizRepositoryImpl) GetLastTake(quizId uuid.UUID, userEmail string) (*quiz.QuizTake, error) {
	result := &quiz.QuizTake{}
	err := q.db.
		Where("quiz_id = ? AND email = ?", quizId, userEmail).
		Last(result).Error

	return result, err
}
