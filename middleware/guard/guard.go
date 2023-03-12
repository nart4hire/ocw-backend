package guard

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type GuardMiddleware struct {
	Token  token.TokenUtil
	Role   []user.UserRole
	Logger logger.Logger
	wrapper.WrapperUtil
}

type ContextKey string

const UserContext ContextKey = "user_claim"

func (g GuardMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if len(g.Role) > 0 {
			authorization := r.Header.Get("Authorization")

			if authorization != "" {
				g.Logger.Info("Unauthorized access detected")

				w.WriteHeader(http.StatusUnauthorized)
				payload := g.WrapperUtil.ErrorResponseWrap("authorization is required", nil)

				parser := json.NewEncoder(w)
				parser.Encode(payload)
				return
			}

			tokenString := strings.Split(authorization, " ")[1]
			claim, err := g.Token.Validate(tokenString, authToken.Access)

			if err != nil {
				g.Logger.Info("Invalid token request")
				parser := json.NewEncoder(w)

				w.WriteHeader(http.StatusUnauthorized)
				payload := g.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
				parser.Encode(payload)
				return
			}

			isAuthorized := false

			for _, user := range g.Role {
				if user == claim.Role {
					isAuthorized = true
				}
			}

			if !isAuthorized {
				g.Logger.Info("Unauthorized user access")
				parser := json.NewEncoder(w)

				w.WriteHeader(http.StatusForbidden)
				payload := g.WrapperUtil.ErrorResponseWrap("current user role is prohibited to access this resources", nil)
				parser.Encode(payload)
				return
			}

			ctx := context.WithValue(r.Context(), UserContext, claim)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		next.ServeHTTP(w, r)
	})
}
