// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/api/option"
	"log"
	"net/http"
	// "time"
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
	// initial firebase
	var dbconn interface{}
	utils.GetDecData("firebase", "", &dbconn)
	data, _ := json.Marshal(dbconn)
	opt := option.WithCredentialsJSON(data)
	Conn, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	app.Firebase = Conn
}

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.Use(gin.Logger())
	app.Router.LoadHTMLGlob("templates/*.html")
	app.Router.Static("/static", "static")
	app.Router.GET("/", app.index)

	app.admin_auth()

	app.user_skill()
	app.user_certificate()
	app.user_auth()
	app.user_music()
	app.user_project()
	app.user_contact()

}

func (app *AppSchema) BindRequestJSON(c *gin.Context, data interface{}) {
	if config.MODE == "PROD" {
		var binding models.RequestProdSchema
		c.BindJSON(&binding)
		d := utils.Ed.BNE(6, 2).Dec(binding.Data)
		json.Unmarshal([]byte(d), &data)
	} else if config.MODE == "DEV" {
		var binding models.RequestSchema
		c.BindJSON(&binding)

		byt, _ := json.Marshal(binding.Data)
		fmt.Print(binding.Data)
		json.Unmarshal(byt, &data)
	}
}

func (app *AppSchema) routeRegister(method string, url string, handler gin.HandlerFunc, middleware bool) {
	app.Router.Handle(method, utils.RouteAPI(url), func(c *gin.Context) {
		utils.Block{
			Try: func() {
				if middleware {
					status, msg := app.routeMiddleware(c)
					if status == 1 {
						utils.Throw(msg)
					}
				}
				handler(c)
			}, Catch: func(e utils.Exception) {
				utils.ResponseAPIError(c, fmt.Sprint(e))
			},
		}.Go()
	})
}

func (app *AppSchema) Run(addr string) {
	// Middleware
	allowOrigin, allowMethods, allowedHeaders, Debug := config.GetCorsConfig()
	c := cors.New(cors.Options{
		AllowedOrigins: allowOrigin,
		AllowedMethods: allowMethods,
		AllowedHeaders: allowedHeaders,
		// AllowCredentials: true,
		Debug: Debug,
	})
	srv := &http.Server{
		Handler: c.Handler(app.Router),
		Addr:    addr,
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (app *AppSchema) index(c *gin.Context) {
	c.HTML(http.StatusOK, "creator.html", gin.H{"zproduct": "Protofolio", "zcreator": config.CREATOR})
	// c.JSON(http.StatusOK, models.ResponseSchema{Status: 0, Message: "Hello Brow"})

}
