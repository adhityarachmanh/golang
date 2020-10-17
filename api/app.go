// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	"encoding/json"
	// "fmt"
	// "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
	"time"
)

var ctx = context.Background()

type AppSchema struct {
	Router   *gin.Engine
	Firebase *firebase.App
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
	app.modAuth()

}

func (app *AppSchema) getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	utils.Tahan{
		Coba: func() {
			token = strings.Split(token, " ")[1]
		},
		Tangkap: func(e utils.Pengecualian) {
			token = ""
		},
	}.Gas()
	// token = strings.TrimSpace(token)
	return token
}

func (app *AppSchema) loggingMiddleWare(c *gin.Context, description string) {
	token := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	var logging models.Logging
	var status int = 0
	loc, _ := time.LoadLocation("Asia/Jakarta")
	logging.MacAddress, _ = utils.GetMacAddr()
	logging.IPAddress, _ = utils.ExternalIP()
	logging.Message = description
	logging.URL = c.Request.RequestURI
	logging.CreatedAt = time.Now().In(loc)
	if config.MODE == "PROD" {
		logging.URL = strings.ReplaceAll(c.Request.RequestURI, "/", "")
		logging.URL = strings.ReplaceAll(c.Request.RequestURI, "."+config.CREATOR, "")
		logging.URL = utils.Ed.BNE(6, 1).Dec(logging.URL)
	}
	if status == 0 {
		client.Collection("visitor").Doc(token).Collection("logging").Add(ctx, logging)
	}

}

func (app *AppSchema) routeMiddleware(c *gin.Context) (int, string) {

	if _, ok := c.Request.Header["Authorization"]; !ok {
		return 1, "Token tidak ditemukan"
	}
	token := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	_, err := client.Collection("visitor").Doc(token).Get(ctx)
	if err != nil {
		return 1, "Token tidak terdaftar"
	}

	return 0, token
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func (app *AppSchema) routeRegister(method string, url string, handler gin.HandlerFunc) {
	// modified route
	app.Router.Handle(method, utils.RouteAPI(url), func(c *gin.Context) {
		// handle API KEY
		withoutMiddleware := []string{"auth/login"}
		_, found := Find(withoutMiddleware, url)
		if !found {
			r, msg := app.routeMiddleware(c)
			if r == 1 {
				utils.ResponseAPIError(c, msg)
				return
			}
			handler(c)
		} else {
			if _, ok := c.Request.Header["Authorization"]; !ok {
				utils.ResponseAPIError(c, "Token tidak ditemukan")
				return
			}
			handler(c)
		}

	})
}

func (app *AppSchema) Run(addr string) {
	// Middleware
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
	log.Fatal(srv.ListenAndServe())
}

func (app *AppSchema) index(c *gin.Context) {

	c.JSON(http.StatusOK, models.ResponseSchema{Status: 0, Message: "Hello Brow"})

}
