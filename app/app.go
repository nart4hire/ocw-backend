package app

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/middleware"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gitlab.informatika.org/ocw/ocw-backend/routes"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/reporter"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
	"gitlab.informatika.org/ocw/ocw-backend/utils/res"
)

type HttpServer struct {
	server          *chi.Mux
	log             logger.Logger
	logUtil         log.LogUtils
	res             res.Resource
	env             *env.Environment
	reporter        reporter.Reporter
	mail            mail.MailQueue
	middlewaresName []string
}

func New(
	middlewares middleware.MiddlewareCollection,
	routes routes.RouteCollection,
	env *env.Environment,
	log logger.Logger,
	logUtil log.LogUtils,
	res res.Resource,
	reporter reporter.Reporter,
	mailqueue mail.MailQueue,
) *HttpServer {
	r := chi.NewRouter()

	middlewareHandlers, middlewareName := middlewares.Register()
	for _, handler := range middlewareHandlers {
		r.Use(handler.Handle)
	}

	for _, group := range routes.Register() {
		r.Group(group.Register)
	}

	return &HttpServer{
		server:          r,
		log:             log,
		res:             res,
		logUtil:         logUtil,
		env:             env,
		reporter:        reporter,
		middlewaresName: middlewareName,
		mail:            mailqueue,
	}
}
