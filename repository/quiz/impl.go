package quiz

import (
	"time"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
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
	result := []quiz.Quiz{}
	err := q.db.Where("course_id = ?", courseId).Find(result).Error
	return result, err
}

func (q *QuizRepositoryImpl) GetQuizDetail(quizId uuid.UUID) (*quiz.Quiz, error) {
	result := &quiz.Quiz{}
	err := q.db.Where("id = ?", quizId).Find(result).Error
	return result, err
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
