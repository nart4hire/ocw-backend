package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/quiz"
)

// AddQuiz Request Payload
//
//	@Description	Information that should be available when you add a quiz, full json cached in redis while actively editted
type AddQuizRequestPayload struct {
	// Web Token that was appended to the link
	AddQuizToken string

	// Quiz ID
	ID uuid.UUID `json:"id" validate:"required"`

	// Quiz Name
	Name string `json:"name" validate:"required"`

	// Course ID
	CourseID string `json:"course_id" validate:"required"`

	// Tentative: Media
	Media string `json:"media"`

	// Quiz Description (Can be left empty)
	Description string `json:"description"`

	// Contributor Email
	CreatorEmail string `json:"creator_email" validate:"required,email" example:"someone@example.com"`

	// Tentative
	// NomorUrut int `json:"no_urut"`
}

// AddProblem Request Payload
//
//	@Description	Information that should be available when you add a problem to a quiz
type AddProblemRequestPayload struct {
	// Web Token that was appended to the link
	AddProblemToken string

	// Problem ID
	ID uuid.UUID `json:"id" validate:"required"`

	// Tentative: Media
	Media string `json:"media"`

	// Problem Statement
	Statement string `json:"statement" validate:"required"`

	// Problem Type
	Type quiz.ProblemType `json:"type" validate:"required"`

	// Associated Quiz ID
	QuizID uuid.UUID `json:"quiz_id" validate:"required"`
}

// AddAnswer Request Payload
//
//	@Description	Information that should be available when you add an answer to a problem
type AddAnswerRequestPayload struct {
	// Web Token that was appended to the link
	AddAnswerToken string

	// Answer ID
	ID uuid.UUID `json:"id" validate:"required"`

	// Tentative: Media
	Media string `json:"media"`

	// Answer Statement
	Statement string `json:"statement" validate:"required"`

	// Associated Problem ID
	ProblemID uuid.UUID `json:"quiz_id" validate:"required"`

	// AnswerTruth Value
	IsAnswer bool `json:"is_answer" validate:"required"`
}

// UpdateQuiz Request Payload
//
//	@Description	Information that should be available when you update a quiz, full json cached in redis while actively editted
type UpdateQuizRequestPayload struct {
	// Web Token that was appended to the link
	AddQuizToken string

	// Quiz ID, Set by param
	ID uuid.UUID `json:"id"`

	// Quiz Name
	Name string `json:"name"`

	// Course ID
	CourseID string `json:"course_id"`

	// Tentative: Media
	Media string `json:"media"`

	// Quiz Description (Can be left empty)
	Description string `json:"description"`

	// Contributor Email
	CreatorEmail string `json:"creator_email" validate:"email" example:"someone@example.com"`

	// Tentative
	// NomorUrut int `json:"no_urut"`
}

// UpdateProblem Request Payload
//
//	@Description	Information that should be available when you update a problem to a quiz
type UpdateProblemRequestPayload struct {
	// Web Token that was appended to the link
	AddProblemToken string

	// Problem ID, Set by param
	ID uuid.UUID `json:"id"`

	// Tentative: Media
	Media string `json:"media"`

	// Problem Statement
	Statement string `json:"statement"`

	// Problem Type
	Type quiz.ProblemType `json:"type"`

	// Associated Quiz ID
	QuizID uuid.UUID `json:"quiz_id"`
}

// UpdateAnswer Request Payload
//
//	@Description	Information that should be available when you update an answer to a quiz
type UpdateAnswerRequestPayload struct {
	// Web Token that was appended to the link
	AddAnswerToken string

	// Answer ID, Set by param
	ID uuid.UUID `json:"id"`

	// Tentative: Media
	Media string `json:"media"`

	// Answer Statement
	Statement string `json:"statement"`

	// Associated Problem ID
	ProblemID uuid.UUID `json:"quiz_id"`

	// AnswerTruth Value
	IsAnswer bool `json:"is_answer"`
}

// DeleteQuiz Request Payload
//
//	@Description	Information that should be available when you delete using uuid
type DeleteRequestPayload struct {
	// Web Token that was appended to the link
	DeleteToken string

	// Quiz ID, provided by query
	ID uuid.UUID
}

// GetUUID Request Payload
//
//	@Description	Information that should be available when you get using uuid
type GetRequestPayload struct {
	// Major/Faculty ID, provided by query
	ID uuid.UUID
}
