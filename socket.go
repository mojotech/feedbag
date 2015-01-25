package main

import (
	"encoding/json"
	"log"

	"github.com/googollee/go-socket.io"
)

func SetupSocketIO() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}

	checkErr(err, "Problem making socket.io server:")
	server.On("connection", func(so socketio.Socket) {
		log.Println("New socket.io connection:", so.Id())
		so.Join("feedbag")
		so.On("disconnection", func() {
			// no op
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("Socket.io server error:", err.Error())
	})

	return server, nil
}

func StartSocketPusher(s *socketio.Server, c <-chan []Activity) error {
	go func() {
		for {
			activities := <-c
			json, err := json.Marshal(activities)
			if err != nil {
				log.Println("** Problem marshaling activities:", err.Error())
			} else {
				s.BroadcastTo("feedbag", "activity", string(json))
			}
		}
	}()

	return nil
}
