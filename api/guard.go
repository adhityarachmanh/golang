package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"encoding/json"
	// "cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type TokenData struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

func (app *AppSchema) getToken(c *gin.Context) (string, string) {
	var tokenHeader string
	var data TokenData
	utils.Block{
		Try: func() {
			tokenHeader = c.GetHeader("Authorization")
			tokenHeader = strings.Split(tokenHeader, " ")[1]
			tokenHeader = utils.Ed.BNE(6, 1).Dec(tokenHeader)
			json.Unmarshal([]byte(tokenHeader), &data)
		},
		Catch: func(e utils.Exception) {
			data = TokenData{Type: "", Token: ""}
		},
	}.Go()
	return data.Token, data.Type
}

func (app *AppSchema) loggingMiddleWare(c *gin.Context, description string) {
	Token, Type := app.getToken(c)
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
	client.Collection(Type).Doc(Token).Collection("logging").Add(ctx, logging)

}

func (app *AppSchema) routeMiddleware(c *gin.Context) (int, string) {
	if _, ok := c.Request.Header["Authorization"]; !ok {
		return 1, "Token not found."
	}
	Token, Type := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	_, err := client.Collection(Type).Doc(Token).Get(ctx)
	if err != nil {
		return 1, "Token not registered."
	}
	return 0, ""
}
