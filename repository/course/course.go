package course

import (
	"errors"

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

func (repo CourseRepositoryImpl) AddCourse(course course.Course) error {
	return repo.db.Create(&course).Error
}

func (repo CourseRepositoryImpl) GetCourse(id string) (*course.Course, error) {
	result := &course.Course{}
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

func (repo CourseRepositoryImpl) GetAllCourseByMajor(abbr string) ([]course.Course, error) {
	var result []course.Course
	err := repo.db.InnerJoins("Major", repo.db.Where(&course.Major{Abbreviation: abbr})).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllCourseMyFaculty(abbr string) ([]course.Course, error) {
	var result []course.Course
	err := repo.db.InnerJoins("Faculty", repo.db.Where(&course.Faculty{Abbreviation: abbr})).InnerJoins("Major").Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) UpdateCourse(course course.Course) error {
	return repo.db.Save(course).Error
}

func (repo CourseRepositoryImpl) DeleteCourse(id string) error {
	return repo.db.Where("id = ?", id).Delete(&course.Course{}).Error
}
