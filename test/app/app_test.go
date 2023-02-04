package app

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	test "gitlab.informatika.org/ocw/ocw-backend/test/utils"
)

func TestUnitAppHandler(t *testing.T) {
	t.Run("HomeEndpoint", func(t *testing.T) {
		t.Run("GetHomeShouldBeOK", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "GET",
				Endpoint: "/",
			})
			assert.Nil(t, err)

			decoder := json.NewDecoder(res.Body)
			responseData := &web.BaseResponse{}

			err = decoder.Decode(responseData)
			assert.Nil(t, err)

			assert.Equal(t, res.StatusCode, http.StatusOK)
			assert.Equal(t, responseData.Status, web.Success)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})

		t.Run("PostHomeShouldBeNotFound", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "POST",
				Endpoint: "/",
			})
			assert.Nil(t, err)

			decoder := json.NewDecoder(res.Body)
			responseData := &web.BaseResponse{}

			err = decoder.Decode(responseData)

			assert.Nil(t, err)
			assert.Equal(t, res.StatusCode, http.StatusNotFound)
			assert.Equal(t, responseData.Status, web.Failed)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})
	})

	t.Run("PingEndpoint", func(t *testing.T) {
		t.Run("GetPingShouldBeOK", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "GET",
				Endpoint: "/ping",
			})
			assert.Nil(t, err)

			decoder := json.NewDecoder(res.Body)
			responseData := &web.BaseResponse{}

			err = decoder.Decode(responseData)
			assert.Nil(t, err)

			assert.Equal(t, res.StatusCode, http.StatusOK)
			assert.Equal(t, responseData.Status, web.Success)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})

		t.Run("GetPingWithSlashShouldBeOk", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "GET",
				Endpoint: "/ping/",
			})
			assert.Nil(t, err)

			assert.Equal(t, res.StatusCode, http.StatusMovedPermanently)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})

		t.Run("GetPingWithManySlashShouldBeOk", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "GET",
				Endpoint: "/ping/////",
			})
			assert.Nil(t, err)

			assert.Nil(t, err)
			assert.Equal(t, res.StatusCode, http.StatusMovedPermanently)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})

		t.Run("PostPingShouldBeOK", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "POST",
				Endpoint: "/ping",
			})
			assert.Nil(t, err)

			decoder := json.NewDecoder(res.Body)
			responseData := &web.BaseResponse{}

			err = decoder.Decode(responseData)

			assert.Nil(t, err)
			assert.Equal(t, res.StatusCode, http.StatusOK)
			assert.Equal(t, responseData.Status, web.Success)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})

		t.Run("PostPingWithPathShouldNotFound", func(t *testing.T) {
			res, log, err := test.ExecuteJSON(test.RequestData{
				Method:   "POST",
				Endpoint: "/ping///ping",
			})
			assert.Nil(t, err)

			decoder := json.NewDecoder(res.Body)
			responseData := &web.BaseResponse{}

			err = decoder.Decode(responseData)

			assert.Nil(t, err)
			assert.Equal(t, res.StatusCode, http.StatusNotFound)
			assert.Equal(t, responseData.Status, web.Failed)
			assert.Equal(t, log.GetCount(test.Info), 1)
		})
	})
}
