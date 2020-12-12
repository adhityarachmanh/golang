package api

import (
	"arh/pkg/utils"
	"encoding/json"
	// "cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"strings"
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

func (app *AppSchema) routeMiddleware(c *gin.Context) (int, string) {
	if _, ok := c.Request.Header["Authorization"]; !ok {
		return 1, "Token not found."
	}
	// Token, Type := app.getToken(c)

	//------------query check token is registered---------//

	// if Type == "visitors" {
	// 	var visitor models.Visitor
	// 	app.firestoreGetDocument(Type, Token, &visitor)

	// 	if visitor.Uid == "" {
	// 		return 1, "Token not registered."
	// 	}
	// } else if Type == "admins" {
	// 	var admins []models.Admin
	// 	app.firestoreFilter(Type, Filter{Key: "token", Op: "==", Value: Token}, &admins)
	// 	if len(admins) == 0 {
	// 		return 1, "Token not registered."
	// 	}
	// }

	return 0, ""
}
