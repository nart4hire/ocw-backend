package httputil

import (
	"encoding/json"
	"net/http"
)

func (HttpUtilImpl) ParseJson(r *http.Request, output interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(output)
}
