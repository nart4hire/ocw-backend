package recoverer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type RecovererMiddleware struct {
	Logger logger.Logger
	wrapper.WrapperUtil
}

func (rm RecovererMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				parser := json.NewEncoder(w)

				w.WriteHeader(http.StatusInternalServerError)
				payload := rm.WrapperUtil.ErrorResponseWrap("internal server error", nil)

				err := parser.Encode(payload)

				if err != nil {
					rm.Logger.Error("Failed to parse error:" + err.Error())
					rm.Logger.Error("")

					return
				}

				stacks := strings.Split(string(debug.Stack()), "\n")
				rm.Logger.Error("Some panic occured when processing request:")
				rm.Logger.Error(fmt.Sprint(rec))
				rm.Logger.Error("")

				rm.Logger.Error("Stack Trace:")
				for _, val := range stacks {
					rm.Logger.Error(val)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
