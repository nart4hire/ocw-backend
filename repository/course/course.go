package course

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func New(
	db db.Database,
) *CourseRepositoryImpl {
	return &CourseRepositoryImpl{db.Connect()}
}

func (repo CourseRepositoryImpl) IsCourseExist(id string) (bool, error) {
	_, err := repo.GetCourse(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return true, err
	}

	return false, nil
}

func (repo CourseRepositoryImpl) IsMajorExist(id uuid.UUID) (bool, error) {
	_, err := repo.GetMajor(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return true, err
	}

	return false, nil
}

func (repo CourseRepositoryImpl) IsFacultyExist(id uuid.UUID) (bool, error) {
	_, err := repo.GetFaculty(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return true, err
	}

	return false, nil
}

func (repo CourseRepositoryImpl) AddCourse(course course.Course) error {
	return repo.db.Create(&course).Error
}

func (repo CourseRepositoryImpl) AddMajor(major course.Major) error {
	return repo.db.Create(&major).Error
}

func (repo CourseRepositoryImpl) AddFaculty(faculty course.Faculty) error {
	return repo.db.Create(&faculty).Error
}

func (repo CourseRepositoryImpl) GetCourse(id string) (*course.Course, error) {
	result := &course.Course{}
	err := repo.db.First(result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetMajor(id uuid.UUID) (*course.Major, error) {
	result := &course.Major{}
	err := repo.db.First(result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetFaculty(id uuid.UUID) (*course.Faculty, error) {
	result := &course.Faculty{}
	err := repo.db.First(result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllCourse() ([]course.Course, error) {
	var result []course.Course
	err := repo.db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllMajor() ([]course.Major, error) {
	var result []course.Major
	err := repo.db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllFaculty() ([]course.Faculty, error) {
	var result []course.Faculty
	err := repo.db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllCourseByMajor(id uuid.UUID) ([]course.Course, error) {
	var result []course.Course
	err := repo.db.InnerJoins("Major", repo.db.Where(&course.Major{ID: id})).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllCourseByFaculty(id uuid.UUID) ([]course.Course, error) {
	var result []course.Course
	err := repo.db.InnerJoins("Faculty", repo.db.Where(&course.Faculty{ID: id})).InnerJoins("Major").Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllMajorByFaculty(id uuid.UUID) ([]course.Major, error) {
	var result []course.Major
	err := repo.db.InnerJoins("Faculty", repo.db.Where(&course.Faculty{ID: id})).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) UpdateCourse(course course.Course) error {
	return repo.db.Save(course).Error
}

func (repo CourseRepositoryImpl) UpdateMajor(major course.Major) error {
	return repo.db.Save(major).Error
}

func (repo CourseRepositoryImpl) UpdateFaculty(faculty course.Faculty) error {
	return repo.db.Save(faculty).Error
}

func (repo CourseRepositoryImpl) DeleteCourse(id string) error {
	return repo.db.Where("id = ?", id).Delete(&course.Course{}).Error
}
