package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type RequestData struct {
	Method   string
	Endpoint string
	Data     interface{}
	Headers  map[string]string
}

func ExecuteJSON(reqData RequestData) (*http.Response, *MockLogger, error) {
	r, log, err := NewTestHandler()

	if err != nil {
		return nil, nil, err
	}

	var req *http.Request

	if reqData.Data == nil {
		req = httptest.NewRequest(reqData.Method, reqData.Endpoint, nil)
	} else {
		byteData, err := json.Marshal(reqData.Data)

		if err != nil {
			return nil, nil, err
		}

		reader := bytes.NewReader(byteData)
		req = httptest.NewRequest(reqData.Method, reqData.Endpoint, reader)
	}

	if reqData.Headers != nil {
		for key, value := range reqData.Headers {
			req.Header.Add(key, value)
		}
	}

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	return res, log, nil
}
