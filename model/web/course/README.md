### Routes

- Getting:
course/major - Get all majors
course/major/{id} - Get major with id
course/major/courses/{id} - Get all courses in major with id
course/faculty - Get all faculties
course/faculty/{id} - Get faculty with id
course/faculty/courses/{id} - Get all courses in faculty with id
course/faculty/majors/{id} - Get all majors in faculty with id
course - get all courses
course/{id} - get course with id

- Adding:
course/major - add a major
course/faculty - add a faculty
course - add a course

- Updating:
course/major/{id} - update a major
course/faculty/{id} - update a faculty
course/{id} - update a course

- Deleting:
course/{id} - delete course with id

### Cannot Delete Majors/Faculties ###
Notes: ID is interchangeable with abbreviation (has OR relationship)