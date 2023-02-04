package utils

import (
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/test"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

func NewTestHandler() (http.Handler, *MockLogger, error) {
	logger := NewMockLogger()
	handler, err := test.CreateServer(logger, &env.Environment{
		AppEnvironment: "DEVELOPMENT",
	})

	if err != nil {
		return nil, nil, err
	}

	logger.CleanLog()
	return handler.GetServer(), logger, nil
}
