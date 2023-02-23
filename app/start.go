package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
)

func (l *HttpServer) Start() {
	listenAddr := fmt.Sprintf("%s:%d", l.env.ListenAddress, l.env.ListenPort)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: l.server,
	}

	serverCtx, cancelServer := context.WithCancel(context.Background())
	l.reporter.Start(serverCtx)
	l.mail.Start(serverCtx)

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		defer cancelServer()
		defer l.log.Info("🛑 Server is successfully shut down")

		<-sig

		forceQuit, cancelForceQuit := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelForceQuit()

		group := sync.WaitGroup{}
		group.Add(1)

		go func() {
			defer group.Done()

			l.log.Info("⏱️ Gracefully shutdown....")
			<-forceQuit.Done()

			if forceQuit.Err() == context.DeadlineExceeded {
				l.log.Error("⏱️ Waiting timeout, force shutdown...")
			}
		}()

		err := server.Shutdown(forceQuit)
		if err != nil {
			l.log.Error(err.Error())
		}

		cancelForceQuit()
		group.Wait()
	}()

	l.log.Info(fmt.Sprintf("🌎 Server Listen at %s",
		l.logUtil.ColoredOutput(
			"http://"+listenAddr,
			log.ForeGreen,
		),
	))

	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		l.log.Error("🔥 Failed to start server")
		l.log.Error(err.Error())
		os.Exit(1)
	}

	<-serverCtx.Done()
}
