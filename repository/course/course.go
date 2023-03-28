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
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (repo CourseRepositoryImpl) IsMajorExist(id uuid.UUID) (bool, error) {
	_, err := repo.GetMajor(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (repo CourseRepositoryImpl) IsFacultyExist(id uuid.UUID) (bool, error) {
	_, err := repo.GetFaculty(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (repo CourseRepositoryImpl) AddCourse(course course.Course) error {
	return repo.db.Create(&course).Error
}

func (repo CourseRepositoryImpl) AddMajor(major course.Major) error {
	if id, err := uuid.NewUUID(); err != nil {
		major.ID = id
	}

	return repo.db.Create(&major).Error
}

func (repo CourseRepositoryImpl) AddFaculty(faculty course.Faculty) error {
	if id, err := uuid.NewUUID(); err != nil {
		faculty.ID = id
	}

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
	err := repo.db.Where("major_id = ?", id).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllCourseByFaculty(id uuid.UUID) ([]course.Course, error) {
	var result []course.Course
	err := repo.db.
		Joins("JOIN major ON major.id = course.major_id").
		Joins("JOIN faculty ON faculty.id = major.fac_id").
		Where("faculty.id = ?", id).
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetAllMajorByFaculty(id uuid.UUID) ([]course.Major, error) {
	var result []course.Major
	err := repo.db.Where("fac_id = ?", id).Find(&result).Error

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

func (repo CourseRepositoryImpl) GetMajorByAbbr(abbr string) (*course.Major, error) {
	result := &course.Major{}
	err := repo.db.First(result, "abbreviation = ?", abbr).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo CourseRepositoryImpl) GetFacultyByAbbr(abbr string) (*course.Faculty, error) {
	result := &course.Faculty{}
	err := repo.db.First(result, "abbreviation = ?", abbr).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
