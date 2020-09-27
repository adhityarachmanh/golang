package api

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func (app *AppSchema) modSocket() {
	app.SocketIO.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	app.SocketIO.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	app.SocketIO.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	app.SocketIO.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
}
