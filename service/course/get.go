package course

import (
	"errors"

	domCourse "gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
	"gorm.io/gorm"
)

func (c CourseServiceImpl) GetCourse(payload course.GetByStringRequestPayload) (*domCourse.Course, error) {
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

func (c CourseServiceImpl) GetMajor(payload course.GetByUUIDRequestPayload) (*domCourse.Major, error) {
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

func (c CourseServiceImpl) GetFaculty(payload course.GetByUUIDRequestPayload) (*domCourse.Faculty, error) {
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

func (c CourseServiceImpl) GetAllCourse() ([]domCourse.Course, error) {
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

func (c CourseServiceImpl) GetAllMajor() ([]domCourse.Major, error) {
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

func (c CourseServiceImpl) GetAllFaculty() ([]domCourse.Faculty, error) {
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

func (c CourseServiceImpl) GetAllCourseByMajor(payload course.GetByUUIDRequestPayload) ([]domCourse.Course, error) {
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

func (c CourseServiceImpl) GetAllCourseByFaculty(payload course.GetByUUIDRequestPayload) ([]domCourse.Course, error) {
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

func (c CourseServiceImpl) GetAllMajorByFaculty(payload course.GetByUUIDRequestPayload) ([]domCourse.Major, error) {
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