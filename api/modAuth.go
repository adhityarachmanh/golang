package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modAuth() {
	app.routeRegister("POST", "auth/login", app.loginVisitor)
	app.routeRegister("POST", "auth/autologin", app.autoLoginVisitor)

}

func (app *AppSchema) autoLoginVisitor(c *gin.Context) {
	// var chatting []models.Chatting
	var uid string
	var visitor models.Visitor
	uid = app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)

	result, err := client.Collection("visitor").Doc(uid).Get(ctx)
	if err != nil {
		utils.ResponseAPIError(c, "")
		return
	}
	result.DataTo(&visitor)

	// for i := 0; i < len(chatting); i++ {
	// 	c := chatting[i]
	// 	client.Collection("visitor").Doc(uid).Collection("chatting").Add(ctx, c)
	// }
	utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
}

func (app *AppSchema) loginVisitor(c *gin.Context) {
	var uid string
	var visitor models.Visitor
	utils.Tahan{
		Coba: func() {
			uid = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)

			result, _ := client.Collection("visitor").Doc(uid).Get(ctx)
			if result.Data() != nil {
				utils.ResponseAPIError(c, "")
				return
			}
			_, err := client.Collection("visitor").Doc(uid).Set(ctx, visitor)
			if err != nil {
				utils.ResponseAPIError(c, "")
				return
			}
			app.loggingMiddleWare(c, "LOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		},
		Tangkap: func(e utils.Pengecualian) {
			utils.ResponseAPIError(c, "Server Error!")
		},
	}.Gas()

}
