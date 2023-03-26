package quiz

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ProblemType int

const (
	Choice ProblemType = iota
)

var roleMapping = map[ProblemType]string{
	Choice: "choice",
}

func (ur *ProblemType) Scan(value interface{}) error {
	val := value.(string)

	for key, label := range roleMapping {
		if label == val {
			*ur = key
			return nil
		}
	}

	return fmt.Errorf("invalid user role")
}

func (u ProblemType) Value() (driver.Value, error) {
	value, ok := roleMapping[u]

	if !ok {
		return nil, fmt.Errorf("invalid user role")
	}

	return value, nil
}

func (u *ProblemType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	for key, label := range roleMapping {
		if label == s {
			*u = key
			return nil
		}
	}

	return fmt.Errorf("unkown role, given %s", s)
}

func (u ProblemType) MarshalJSON() ([]byte, error) {
	s, ok := roleMapping[u]

	if !ok {
		return nil, errors.New("unkown user role")
	}

	return json.Marshal(s)
}
