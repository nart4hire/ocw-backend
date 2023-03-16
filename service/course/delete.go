package course

import "gitlab.informatika.org/ocw/ocw-backend/model/web/course/delete"

func (c CourseServiceImpl) DeleteCourse(payload delete.DeleteByStringRequestPayload) error {
	err := c.CourseRepository.DeleteCourse(payload.ID)

	if err != nil {
		// Uncaught error
		return err
	}

	return nil
}