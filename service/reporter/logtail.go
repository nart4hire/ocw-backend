package reporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

type LogtailReporter struct {
	env        *env.Environment
	logUtil    log.LogUtils
	logQueue   []ReporterPayload
	mutex      sync.Mutex
	httpClient *http.Client
	isStarted  bool
}

func New(
	env *env.Environment,
	logUtil log.LogUtils,
) *LogtailReporter {
	return &LogtailReporter{
		env,
		logUtil,
		[]ReporterPayload{},
		sync.Mutex{},
		&http.Client{
			Transport: &http.Transport{
				IdleConnTimeout: time.Duration(env.HttpReqTimeout) * time.Second,
			},
		},
		false,
	}
}

func (l *LogtailReporter) Send(payload ReporterPayload) {
	if !l.isStarted {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.logQueue = append(l.logQueue, payload)
}

func (l *LogtailReporter) Flush() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	payloadBytes, err := json.Marshal(l.logQueue)

	if err != nil {
		l.logUtil.PrintFormattedOutput(
			fmt.Sprintf("Some error happened when parse json: %s", err),
			"REPORT",
			"ERROR",
			log.ForeRed,
		)
	} else {
		l.logQueue = []ReporterPayload{}

		go func() {
			reader := bytes.NewReader(payloadBytes)
			req, err := http.NewRequest("POST", "https://in.logtail.com", reader)

			if err != nil {
				l.logUtil.PrintFormattedOutput(
					fmt.Sprintf("Some error happened when creating request: %s", err),
					"REPORT",
					"ERROR",
					log.ForeRed,
				)
				return
			}

			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", l.env.LogtailToken))
			req.Header.Add("Content-Type", "application/json")

			res, err := l.httpClient.Do(req)

			if err != nil {
				l.logUtil.PrintFormattedOutput(
					fmt.Sprintf("Some error happened when sending request: %s", err),
					"REPORT",
					"ERROR",
					log.ForeRed,
				)
				return
			}

			if res.StatusCode != http.StatusOK {
				l.logUtil.PrintFormattedOutput(
					fmt.Sprintf("Request respose is not 200 OK: got %d", res.StatusCode),
					"REPORT",
					"ERROR",
					log.ForeRed,
				)
			}
		}()
	}

}

func (l *LogtailReporter) Start(ctx context.Context) {
	if l.env.AppEnvironment != "PRODUCTION" {
		l.logUtil.PrintFormattedOutput(
			"Reporter is not started due to non-production environment",
			"REPORT",
			"WARNING",
			log.ForeYellow)
		return
	}

	go func() {
		l.isStarted = true
		defer func() { l.isStarted = false }()
		defer l.Flush()

		timer := time.NewTicker(time.Second)
		defer timer.Stop()

		l.logUtil.PrintFormattedOutput(
			"Reporter started to listen...",
			"REPORT",
			"INFO",
			log.ForeBlue,
		)

		for {
			select {
			case <-ctx.Done():
				break
			case <-timer.C:
				l.Flush()
			}
		}
	}()
}

func (l *LogtailReporter) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.logQueue = []ReporterPayload{}
}
