package app

import (
	"fmt"

	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

var colorMap = map[string][]log.Color{
	"GET":    {log.BackGreen, log.ForeWhite},
	"POST":   {log.BackBlue, log.ForeWhite},
	"PUT":    {log.BackCyan, log.ForeWhite},
	"PATCH":  {log.BackMagenta, log.ForeWhite},
	"DELETE": {log.BackRed, log.ForeWhite},
}

func colorizeMethod(name string) string {
	res := ""
	val, ok := colorMap[name]

	if ok {
		res = string(val[0]) + string(val[1])
	} else {
		res = string(log.ForeBlack) + string(log.BackWhite)
	}

	res = res + " "

	res = res + name

	for i := 0; i < 8-(len(name)+1); i++ {
		res = res + " "
	}

	return res + string(log.Reset)
}

func (l HttpServer) ListRoute() {
	routeData := map[string][]string{}

	for _, route := range l.server.Routes() {
		for method := range route.Handlers {
			name := route.Pattern

			if routeData[method] == nil {
				routeData[method] = []string{name}
			} else {
				routeData[method] = append(routeData[method], name)
			}
		}
	}

	l.log.Info("Routes Information:")
	l.log.Info("")

	loggedMethod := []string{
		"GET", "POST", "PUT", "PATCH", "DELETE",
	}

	for _, method := range loggedMethod {
		for _, pattern := range routeData[method] {
			l.log.Info(
				fmt.Sprintf("%s %s",
					colorizeMethod(method),
					pattern),
			)
		}
	}

	l.log.Info("")
}

func (l HttpServer) ListMiddleware() {
	if len(l.middlewaresName) > 0 {
		l.log.Info("Registered Middlewares:")
		for _, middleware := range l.middlewaresName {
			l.log.Info("- " + l.logUtil.ColoredOutput(middleware, log.ForeGreen))
		}
	} else {
		l.log.Info("No middleware registered")
	}

	l.log.Info("")
}
