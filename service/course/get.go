package course

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/get"
	"gorm.io/gorm"
)

func (c CourseServiceImpl) GetCourse(payload get.GetByStringRequestPayload) (*course.Course, error) {
	packet, err := c.CourseRepository.GetCourse(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.CourseNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetMajor(payload get.GetByUUIDRequestPayload) (*course.Major, error) {
	packet, err := c.CourseRepository.GetMajor(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.MajorNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetFaculty(payload get.GetByUUIDRequestPayload) (*course.Faculty, error) {
	packet, err := c.CourseRepository.GetFaculty(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.FacultyNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllCourse() ([]course.Course, error) {
	packet, err := c.CourseRepository.GetAllCourse()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.CourseNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllMajor() ([]course.Major, error) {
	packet, err := c.CourseRepository.GetAllMajor()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.MajorNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllFaculty() ([]course.Faculty, error) {
	packet, err := c.CourseRepository.GetAllFaculty()

	if err != nil {
		// This should not happen unless data is unpopulated, faculty is the root dependency so there should be at least one
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.FacultyNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllCourseByMajor(payload get.GetByUUIDRequestPayload) ([]course.Course, error) {
	packet, err := c.CourseRepository.GetAllCourseByMajor(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.CourseNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllCourseByFaculty(payload get.GetByUUIDRequestPayload) ([]course.Course, error) {
	packet, err := c.CourseRepository.GetAllCourseByFaculty(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.CourseNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}

func (c CourseServiceImpl) GetAllMajorByFaculty(payload get.GetByUUIDRequestPayload) ([]course.Major, error) {
	packet, err := c.CourseRepository.GetAllMajorByFaculty(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.MajorNotExist)
		}
		// Some Uncaught error
		return nil, err
	}

	return packet, nil
}