package server

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type clientStreams struct {
	mtx     sync.RWMutex
	streams map[room.UserId]chan *api.Event
}

func newStreams() clientStreams {
	return clientStreams{streams: make(map[room.UserId]chan *api.Event, 1000)}
}

func (s *clientStreams) sendSingle(user room.UserId, event *api.Event) {
	stream, ok := s.streams[user]
	if ok {
		stream <- event
	}
}

func (s *clientStreams) send(users []room.UserId, event *api.Event) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	for _, user := range users {
		s.sendSingle(user, event)
	}
}

func (s *clientStreams) startNew(srv api.Doko_StartSessionServer, user room.UserId) {
	stream := s.createStream(user)
	go func() {
		defer s.removeStream(user)
		for {
			select {
			case <-srv.Context().Done():
				log.Printf("no longer waiting for messages to %s", user)
				return
			case event := <-stream:
				if s, ok := status.FromError(srv.Send(event)); ok {
					switch s.Code() {
					case codes.OK:
						// pass
					case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
						log.Printf("client %s terminated connection", user)
						return

					default:
						log.Printf("failed to send to client %s: %v", user, s.Err())
					}
				}
			}
		}
	}()
}

func (s *clientStreams) createStream(user room.UserId) (stream chan *api.Event) {
	stream = make(chan *api.Event, 10)
	s.streams[user] = stream
	return
}

func (s *clientStreams) removeStream(user room.UserId) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if stream, ok := s.streams[user]; ok {
		delete(s.streams, user)
		close(stream)
	}
	log.Printf("closed table stream for %s", user)
}

func (s *clientStreams) isOnline(user room.UserId) bool {
	_, ok := s.streams[user]
	return ok
}
