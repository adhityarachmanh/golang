// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	// fb "firebase.google.com/go/v4/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	// FCM "google.golang.org/api/fcm/v1"
	"google.golang.org/api/option"
	// "google.golang.org/api/transport"
	"bytes"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()

type AppSchema struct {
	Router   *gin.Engine
	Firebase *firebase.App
	// FCM      *FCM.Service
}

func (app *AppSchema) Initialize() {
	app.initializeFirebase()
	// app.initializeFCM()
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

// func (app *AppSchema) initializeFCM() {
// 	var fcmTokenKey string
// 	utils.GetDecData("fcm", "", &fcmTokenKey)
// 	fcmService, _ := FCM.NewService(ctx, option.WithAPIKey(fcmTokenKey))
// 	app.FCM = fcmService
// }

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.Use(gin.Logger())
	app.Router.LoadHTMLGlob("templates/*.html")
	app.Router.Static("/static", "static")
	app.Router.GET("/", app.index)

	app.admin_auth()
	app.admin_user()

	app.user_skill()
	app.user_certificate()
	app.user_auth()
	app.user_music()
	app.user_project()
	app.user_contact()

}

func (app *AppSchema) sendNotifToAdmin(fmcRequest models.FCMRequest) []int {
	var adminList []models.Admin
	loc, _ := time.LoadLocation("Asia/Jakarta")
	app.firestoreByCollection("admins", &adminList)
	var status []int
	for i := 0; i < len(adminList); i++ {
		fmcRequest.Data["notificationID"] = utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).String())
		admin := adminList[i]
		stat := app.sendNotificationFCM(models.FCM{
			Notification: models.NotificationSchema{
				Title:       fmcRequest.Title,
				Body:        fmcRequest.Body,
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			},
			Data: fmcRequest.Data,
			To:   admin.NotificationToken,
		})
		status = append(status, stat)
	}
	return status
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

func (app *AppSchema) sendNotificationFCM(fcmData models.FCM) int {
	var FCMKey models.FCMKey
	utils.GetDecData("fcm", "", &FCMKey)
	requestBody, _ := json.Marshal(fcmData)
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	request, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(requestBody))
	if err != nil {
		return 1
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "key="+FCMKey.Key)

	resp, err := client.Do(request)
	if err != nil {
		return 1
	}
	var result models.FCMResponse
	json.NewDecoder(resp.Body).Decode(&result)

	// defer resp.Body.Close()
	if result.Success > 0 {
		return 0
	} else if result.Failure > 0 {
		return 1
	} else {
		return 1
	}

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
