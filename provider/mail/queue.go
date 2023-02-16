package mail

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type MailQueueImpl struct {
	provider  MailProvider
	queue     []Mail
	mutex     sync.Mutex
	logUtil   logger.Logger
	env       *env.Environment
	isStarted bool
}

func NewQueue(
	provider MailProvider,
	log logger.Logger,
	env *env.Environment,
) *MailQueueImpl {
	return &MailQueueImpl{
		provider:  provider,
		queue:     []Mail{},
		mutex:     sync.Mutex{},
		logUtil:   log,
		isStarted: false,
		env:       env,
	}
}

func (mq *MailQueueImpl) Send(mail Mail) {
	if !mq.isStarted {
		return
	}

	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	mq.queue = append(mq.queue, mail)
}

func (mq *MailQueueImpl) Flush() {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	if len(mq.queue) == 0 {
		return
	}

	sendBuffer := make([]Mail, len(mq.queue))
	copy(sendBuffer, mq.queue)

	go func(buffer *[]Mail) {
		for _, data := range *buffer {
			err := mq.provider.Send(data.To, data.Subject, data.Message)

			if err != nil {
				mq.logUtil.Error(
					fmt.Sprintf("Problem with sending mail: %s", err.Error()),
				)
			}
		}
	}(&sendBuffer)

	mq.queue = []Mail{}
}

func (mq *MailQueueImpl) Start(ctx context.Context) {
	go func() {
		mq.isStarted = true
		defer func() {
			mq.isStarted = false
			mq.logUtil.Info("🛑 Mail queue has been stopped.")
		}()

		interval := time.Duration(mq.env.MailingInterval)
		timer := time.NewTicker(interval * time.Millisecond)
		defer timer.Stop()

		mq.logUtil.Info(
			fmt.Sprintf("Mailing started to listen... (interval: %dms)", interval),
		)

		isLoop := true

		for isLoop {
			select {
			case <-ctx.Done():
				isLoop = false
			case <-timer.C:
				mq.Flush()
			}
		}
	}()
}
