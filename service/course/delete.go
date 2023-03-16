package course

import "gitlab.informatika.org/ocw/ocw-backend/model/web/course/get"

func (c CourseServiceImpl) DeleteCourse(payload get.GetByStringRequestPayload) error {
	err := c.CourseRepository.DeleteCourse(payload.ID)

	if err != nil {
		// Uncaught error
		return err
	}

	return nil
}