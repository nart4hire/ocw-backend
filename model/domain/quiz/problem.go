package quiz

import "github.com/google/uuid"

type QuizMedia struct {
	Id   uuid.UUID `json:"id"`
	Url  string    `json:"url"`
	Type string    `json:"type"`
}

type AnswerOption struct {
	Id         uuid.UUID   `json:"id"`
	MediaId    []uuid.UUID `json:"media_id"`
	Answer     string      `json:"answer"`
	IsSolution *bool       `json:"is_solution"`
}

type QuizProblem struct {
	Id       uuid.UUID      `json:"id"`
	MediaId  []uuid.UUID    `json:"media_id"`
	Question string         `json:"question"`
	Answer   []AnswerOption `json:"answers"`
}

type QuizDetail struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	CourseId    string        `json:"course_id"`
	Description string        `json:"description"`
	Help        string        `json:"help"`
	Media       []QuizMedia   `json:"media"`
	Problems    []QuizProblem `json:"problems"`
}

type Response struct {
	ProblemId uuid.UUID `json:"problem_id"`
	AnswerId  uuid.UUID `json:"answer_id"`
}
