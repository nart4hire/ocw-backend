package app

import (
	"fmt"

	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

func (l HttpServer) Version() {
	data, err := l.res.GetStringResource("ascii.art")

	if err == nil {
		fmt.Println(
			l.logUtil.ColoredOutput(
				data,
				log.ForeGreen,
			),
		)
		println()
	}

	data, err = l.res.GetStringResource("version")

	if err == nil {
		fmt.Println(
			l.logUtil.ColoredOutput(
				data,
				log.ForeCyan,
			),
		)
		println()
	}
}
