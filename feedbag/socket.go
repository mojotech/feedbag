package feedbag

import (
	"github.com/fogcreek/logging"
	"github.com/googollee/go-socket.io"
)

func SetupSocketIO() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}

	server.On("connection", func(so socketio.Socket) {
		logging.InfoWithTags([]string{"socket.io"}, "New socket.io connection:", so.Id())
		so.Join("feedbag")
		so.On("disconnection", func() {
			// no op
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		logging.ErrorWithTags([]string{"socket.io"}, "Error on socket.io server", err.Error())
	})

	return server, nil
}

func StartSocketPusher(s *socketio.Server, c <-chan []Activity) error {
	go func() {
		for {
			activities := <-c
			s.BroadcastTo("feedbag", "activity", activities)
		}
	}()

	return nil
}
