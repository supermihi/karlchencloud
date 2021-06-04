package server

import (
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type eventStream chan *pb.Event

type ClientStream struct {
	events eventStream
	kicked chan int
}

type ClientStreams struct {
	mtx     sync.RWMutex
	clients map[room.UserId]ClientStream
}

func NewClientStreams() ClientStreams {
	return ClientStreams{clients: make(map[room.UserId]ClientStream)}
}

// SendSingle synchronously sends an event to a single user stream.
func (cs *ClientStreams) SendSingle(user room.UserId, event *pb.Event) {
	cs.mtx.RLock()
	client, ok := cs.clients[user]
	if ok {
		client.events <- event
	}
	cs.mtx.RUnlock()
}

// Send synchronously sends an event to a list of users.
func (cs *ClientStreams) Send(users []room.UserId, event *pb.Event) {
	for _, user := range users {
		cs.SendSingle(user, event)
	}
}

func (cs *ClientStreams) StartNew(srv pb.Doko_StartSessionServer, user room.UserId) chan int {
	client := cs.createStream(user)
	go func() {
		for {
			select {
			case <-srv.Context().Done():
				log.Printf("no longer waiting for messages to %s", user)
				cs.onStreamEndedByClient(user)
				return
			case event, ok := <-client.events:
				if !ok {
					log.Printf("client kicked, returning stream loop")
					return
				}
				if s, sendOk := status.FromError(srv.Send(event)); sendOk {
					switch s.Code() {
					case codes.OK:
						// pass
					case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
						log.Printf("client %s terminated connection", user)
						cs.onStreamEndedByClient(user)
						return
					default:
						log.Printf("failed to Send to client %s: %v", user, s.Err())
					}
				} else {
					log.Printf("unknonw error: %s", s.Err())
				}
			}
		}
	}()
	return client.kicked
}

const userStreamBufferSize = 10

func (cs *ClientStreams) createStream(user room.UserId) ClientStream {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()
	if existingClient, exists := cs.clients[user]; exists {
		delete(cs.clients, user)
		close(existingClient.events)
		close(existingClient.kicked)
		log.Printf("kicked %s because she started a new session", user)
	}
	client := ClientStream{events: make(eventStream, userStreamBufferSize),
		kicked: make(chan int)}
	cs.clients[user] = client
	return client
}

func (cs *ClientStreams) onStreamEndedByClient(user room.UserId) {
	cs.mtx.Lock()
	if client, exists := cs.clients[user]; exists {
		delete(cs.clients, user)
		close(client.events)
		close(client.kicked)
		log.Printf("closed table stream for %s", user)
	}
	cs.mtx.Unlock()

}

func (cs *ClientStreams) IsOnline(user room.UserId) bool {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	_, ok := cs.clients[user]
	return ok
}
