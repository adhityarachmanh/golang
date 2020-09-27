// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	// "fmt"
	"github.com/gin-gonic/gin"
	// socketio "github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

// var Ed utils.BNESchema = utils.BNESchema{}
var ctx = context.TODO()

type AppSchema struct {
	Router *gin.Engine
	// SocketIO *socketio.Server
	DB *mongo.Database
}

func (app *AppSchema) Initialize() {
	app.initializeDatabase()
	// app.initializeSocketIO()
	app.initializeRoutes()
}

func (app *AppSchema) initializeDatabase() {
	// initial database
	var dbconn models.DatabaseSchema
	utils.GetDecData("db", "", &dbconn)
	clientOptions := options.Client().ApplyURI(dbconn.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)

	}
	app.DB = client.Database(dbconn.Db)
}

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.GET("/", app.index)
	app.modSkill()
	app.modCertificate()
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
