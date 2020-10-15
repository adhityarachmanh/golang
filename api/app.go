// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	"encoding/json"
	// "fmt"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()
var Ed utils.BNESchema = utils.BNESchema{}

type AppSchema struct {
	Router   *gin.Engine
	Firebase *firebase.App
	ED       BNESchema
}

func (app *AppSchema) Initialize() {
	app.initializeFirebase()
	app.initializeRoutes()
}

func (app *AppSchema) initializeFirebase() {

	// initial database
	var dbconn interface{}
	utils.GetDecData("firebase", "", &dbconn)
	data, _ := json.Marshal(dbconn)
	opt := option.WithCredentialsJSON(data)
	// opt := option.WithCredentialsFile("key/firebase.json")
	Conn, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	app.Firebase = Conn
}

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.GET("/", app.index)
	app.modSkill()
	app.modCertificate()
	app.ED = BNESchema{}
}

// func (app *AppSchema) initializeSocketIO() {
// 	app.SocketIO, _ = socketio.NewServer(nil)
// 	app.modSocket()
// }

func (app *AppSchema) Run(addr string) {

	// app.Router.GET("/socket.io/*any", gin.WrapH(app.SocketIO))
	// app.Router.POST("/socket.io/*any", gin.WrapH(app.SocketIO))
	allowOrigin, allowMethods, allowedHeaders, Debug := config.GetCorsConfig()
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrigin,
		AllowedMethods:   allowMethods,
		AllowedHeaders:   allowedHeaders,
		AllowCredentials: true,
		Debug:            Debug,
	})
	srv := &http.Server{
		Handler:      c.Handler(app.Router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// go app.SocketIO.Serve()
	log.Fatal(srv.ListenAndServe())
}

func (app *AppSchema) index(c *gin.Context) {

	c.JSON(http.StatusOK, models.ResponseSchema{Status: 0, Message: "Hello Brow"})

}
