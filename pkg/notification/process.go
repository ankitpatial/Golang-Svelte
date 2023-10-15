package notification

import (
	"context"
	"roofix/pkg/util/log"
	"sync"
)

func Process(ctx context.Context, msg *Message) error {
	if msg == nil {
		log.Warn("abort notification.Process(), message is nil")
	}

	log.InfoBullet("notification.Process()")
	var wg sync.WaitGroup

	// save to DB
	if msg.SaveToDB {
		wg.Add(1)
		go func(c context.Context, wg *sync.WaitGroup, m Message) {
			defer wg.Done()
			saveToDB(c, &m)
		}(ctx, &wg, *msg)

	}

	// send WS message
	wg.Add(1)
	go func(c context.Context, wg *sync.WaitGroup, m Message) {
		defer wg.Done()
		wsMessage(c, &m)
	}(ctx, &wg, *msg)

	// send push notifications

	wg.Wait()
	return nil
}
