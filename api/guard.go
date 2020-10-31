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
	utils.Block{
		Try: func() {
			token = c.GetHeader("Authorization")
			token = strings.Split(token, " ")[1]
		},
		Catch: func(e utils.Exception) {
			token = ""
		},
	}.Go()
	return token
}

func (app *AppSchema) loggingMiddleWare(c *gin.Context, t string, description string) {
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
	client.Collection(t).Doc(token).Collection("logging").Add(ctx, logging)

}

func (app *AppSchema) routeMiddleware(c *gin.Context, t string) int {
	token := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	_, err := client.Collection(t).Doc(token).Get(ctx)
	if err != nil {
		return 1
	}
	return 0
}
