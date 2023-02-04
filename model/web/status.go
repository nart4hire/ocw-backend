package web

import (
	"errors"
)

type Status uint

const (
	Success Status = iota
	Failed
)

func (s Status) MarshalJSON() ([]byte, error) {
	switch s {
	case Success:
		return []byte("\"success\""), nil
	case Failed:
		return []byte("\"failed\""), nil
	}

	return nil, errors.New("unkown value")
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if string(data) == "\"success\"" {
		*s = Success
		return nil
	} else if string(data) == "\"failed\"" {
		*s = Failed
		return nil
	}

	return errors.New("Unkown type of " + string(data))
}
