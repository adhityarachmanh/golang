package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	// "cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func (app *AppSchema) getToken(c *gin.Context) string {
	var token string
	utils.Tahan{
		Coba: func() {
			token = c.GetHeader("Authorization")
			token = strings.Split(token, " ")[1]
		},
		Tangkap: func(e utils.Exception) {
			token = ""
		},
	}.Gas()
	return token
}

func (app *AppSchema) loggingMiddleWare(c *gin.Context, description string) {
	token := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	var logging models.Logging
	loc, _ := time.LoadLocation("Asia/Jakarta")
	logging.MacAddress, _ = utils.GetMacAddr()
	logging.IPAddress, _ = utils.ExternalIP()
	logging.Message = description
	logging.URL = c.Request.RequestURI
	logging.UserAgent = c.Request.UserAgent()
	logging.CreatedAt = time.Now().In(loc)
	if config.MODE == "PROD" {
		logging.URL = strings.ReplaceAll(c.Request.RequestURI, "/", "")
		logging.URL = strings.ReplaceAll(logging.URL, "."+strings.ToLower(config.CREATOR), "")
		logging.URL = utils.Ed.BNE(6, 1).Dec(logging.URL)
	}
	client.Collection("visitor").Doc(token).Collection("logging").Add(ctx, logging)

}

func (app *AppSchema) routeMiddleware(c *gin.Context) int {
	token := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	_, err := client.Collection("visitor").Doc(token).Get(ctx)
	if err != nil {
		return 1
	}
	return 0
}
