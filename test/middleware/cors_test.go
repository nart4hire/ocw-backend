package middleware

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	test "gitlab.informatika.org/ocw/ocw-backend/test/utils"
)

func TestPreflight(t *testing.T) {
	t.Run("PreflightShouldBeOk", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method:   "OPTIONS",
			Endpoint: "/ping",
			Headers: map[string]string{
				"Access-Control-Request-Method":  "GET",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin":                         "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "https://inkubatorit.com")
		assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "GET")
	})

	t.Run("PreflightShouldBeOkOnNotFound", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method:   "OPTIONS",
			Endpoint: "/not-found",
			Headers: map[string]string{
				"Access-Control-Request-Method":  "GET",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin":                         "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "https://inkubatorit.com")
		assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "GET")
	})

	t.Run("PreflightAllowedPatchMethod", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method:   "OPTIONS",
			Endpoint: "/ping",
			Headers: map[string]string{
				"Access-Control-Request-Method":  "PATCH",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin":                         "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "https://inkubatorit.com")
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Methods"), "PATCH")
	})
}
