package api

// import (
// 	// "arh/pkg/utils"
// 	"encoding/json"
// 	// "fmt"
// 	// "github.com/google/uuid"
// 	socketio "github.com/googollee/go-socket.io"
// )

// type SocketIO struct {
// 	ID   string      `json:"id"`
// 	Data interface{} `json:"data"`
// }

// func (app *AppSchema) modSocket() {
// 	var socket_map map[string]socketio.Conn = make(map[string]socketio.Conn)

// 	app.SocketIO.OnEvent("/", "initialize", func(s socketio.Conn, msg string) {
// 		socket_map[msg] = s
// 	})
// 	app.SocketIO.OnEvent("/", "send", func(s socketio.Conn, msg string) {
// 		var data SocketIO
// 		json.Unmarshal([]byte(msg), &data)
// 		for _, value := range socket_map {
// 			value.Emit("receiver", data.Data)
// 		}

// 	})
// 	app.SocketIO.OnEvent("/", "dc", func(s socketio.Conn, msg string) {
// 		delete(socket_map, msg)
// 	})

// }
