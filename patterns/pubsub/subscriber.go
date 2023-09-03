package pubsub

import (
	"context"
	"log"
	"sync"
)

type message struct {
	data []byte
}

type subscriber struct {
	sync.Mutex

	name    string
	handler chan *message
	quit    chan struct{}
}

func newSubscriber(name string) *subscriber {
	return &subscriber{
		name:    name,
		handler: make(chan *message, 100),
		quit:    make(chan struct{}),
	}
}

func (s *subscriber) run(ctx context.Context) {
	for {
		select {
		case msg := <-s.handler:
			log.Println(s.name, string(msg.data))
		case <-s.quit:
			return
		case <-ctx.Done():
			return
		}
	}
}
