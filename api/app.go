// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	// "fmt"
	"github.com/gin-gonic/gin"
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
	DB     *mongo.Database
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}

func (app *AppSchema) Initialize() {
	app.initializeDatabase()
	app.initializeRoutes()

}

func (app *AppSchema) initializeDatabase() {
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
	app.Router = gin.Default()
	app.Router.GET("/", app.index)
	app.modSkill()
	app.modCertificate()
}

func (app *AppSchema) Run(addr string) {
	srv := &http.Server{
		Handler:      corsMiddleware(app.Router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (app *AppSchema) index(c *gin.Context) {

	c.JSON(http.StatusOK, models.ResponseSchema{Status: 0, Message: "Hello Brow"})

}
