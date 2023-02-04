package log

import (
	"fmt"
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

type RequestLogMiddleware struct {
	LogUtil log.LogUtils
	Logger  logger.Logger
}

var methodColor = map[string]log.Color{
	"GET":    log.ForeGreen,
	"POST":   log.ForeBlue,
	"PUT":    log.ForeCyan,
	"PATCH":  log.ForeMagenta,
	"DELETE": log.ForeRed,
}

var statusColor = map[int]log.Color{
	0:   log.ForeCyan,
	200: log.ForeGreen,
	300: log.ForeBlue,
	400: log.ForeYellow,
	500: log.ForeRed,
}

func (rl RequestLogMiddleware) colorizeMethod(method string) string {
	val, ok := methodColor[method]

	if ok {
		return rl.LogUtil.ColoredOutput(method, val)
	} else {
		return rl.LogUtil.ColoredOutput(method, log.ForeWhite)
	}
}

func (rl RequestLogMiddleware) colorizeCode(code int) string {
	if code < 200 {
		return rl.LogUtil.ColoredOutput(fmt.Sprint(code), statusColor[0])
	} else if code < 300 {
		return rl.LogUtil.ColoredOutput(fmt.Sprint(code), statusColor[200])
	} else if code < 400 {
		return rl.LogUtil.ColoredOutput(fmt.Sprint(code), statusColor[300])
	} else if code < 500 {
		return rl.LogUtil.ColoredOutput(fmt.Sprint(code), statusColor[400])
	} else {
		return rl.LogUtil.ColoredOutput(fmt.Sprint(code), statusColor[500])
	}
}

func (rl RequestLogMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ww := chiMiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			delta := time.Since(startTime)
			status := ww.Status()
			path := r.URL.Path
			method := r.Method

			rl.Logger.Info(
				fmt.Sprintf("Request %s %s %s (%dms)",
					rl.colorizeCode(status),
					rl.colorizeMethod(method),
					path,
					delta.Milliseconds(),
				),
			)
		}()

		next.ServeHTTP(ww, r)
	})
}
