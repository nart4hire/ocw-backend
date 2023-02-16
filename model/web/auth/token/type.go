package token

import (
	"encoding/json"
	"fmt"
)

type TokenType int

const (
	Access TokenType = iota
	Refresh
)

var tokenMapping = map[TokenType]string{
	Access:  "access",
	Refresh: "refresh",
}

func (t *TokenType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	for key, label := range tokenMapping {
		if label == s {
			*t = key
			return nil
		}
	}

	return fmt.Errorf("unknown token type, given %s", s)
}

func (t TokenType) MarshalJSON() ([]byte, error) {
	s, ok := tokenMapping[t]

	if !ok {
		return nil, fmt.Errorf("unkown token type, given %d", t)
	}

	return json.Marshal(s)
}
