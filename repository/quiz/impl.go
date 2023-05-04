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

func (q *QuizRepositoryImpl) UpdateScore(quizId uuid.UUID, score int) error {
	return q.db.
		Model(&quiz.QuizTake{}).
		Where("quiz_id = ?", quizId).
		Updates(quiz.QuizTake{
			Score:      score,
			IsFinished: true,
		}).Error
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

func (q *QuizRepositoryImpl) NewQuiz(quiz quiz.Quiz) error {
	return q.db.Create(&quiz).Error
}

func (q *QuizRepositoryImpl) UpdateQuiz(quiz quiz.Quiz) error {
	return q.db.Save(quiz).Error
}

func (q *QuizRepositoryImpl) GetQuizLink(quizId uuid.UUID) (string, error) {
	result := &quiz.Quiz{}
	err := q.db.Where("id = ?", quizId).First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", web.NewResponseError("Record not found", "ERR_NOT_FOUND")
	}

	return result.QuizPath, nil
}

func (q *QuizRepositoryImpl) Delete(quizId uuid.UUID) error {
	return q.db.Delete(&quiz.Quiz{}, quizId).Error
}

func (q *QuizRepositoryImpl) IsActiveTake(quizId uuid.UUID, userEmail string) (bool, error) {
	var result int64 = 0
	err := q.db.
		Model(&quiz.QuizTake{}).
		Where("quiz_id = ? AND email = ? AND is_finished = false", quizId, userEmail).
		Count(&result).
		Error

	if err != nil {
		return false, nil
	}

	return result > 0, nil
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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return result, err
}
