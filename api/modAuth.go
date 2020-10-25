// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modAuth() {
	app.routeRegister("POST", "auth/login", app.loginVisitor)
	app.routeRegister("POST", "auth/autologin", app.autoLoginVisitor)
	app.routeRegister("POST", "auth/updatevisitor", app.editVisitor)

}

func (app *AppSchema) editVisitor(c *gin.Context) {
	var visitorRequest models.VisitorRequest
	var visitor models.Visitor

	utils.Tahan{
		Coba: func() {
			app.BindRequestJSON(c, &visitorRequest)
			visitor.Name = visitorRequest.Name
			visitor.Uid = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)
			_, err := client.Collection("visitor").Doc(visitor.Uid).Update(ctx, []firestore.Update{
				{
					Path: "name", Value: visitor.Name,
				},
				{
					Path: "chat", Value: true,
				},
			})
			if err != nil {
				utils.Throw("")
			}
			result, _ := client.Collection("visitor").Doc(visitor.Uid).Get(ctx)
			result.DataTo(&visitor)

			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Tangkap: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Gas()

}
func (app *AppSchema) autoLoginVisitor(c *gin.Context) {
	var visitor models.Visitor
	utils.Tahan{
		Coba: func() {
			visitor.Uid = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)
			result, _ := client.Collection("visitor").Doc(visitor.Uid).Get(ctx)
			result.DataTo(&visitor)
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Tangkap: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Gas()

}

func (app *AppSchema) loginVisitor(c *gin.Context) {

	var visitor models.Visitor
	utils.Tahan{
		Coba: func() {
			visitor.Uid = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)

			result, _ := client.Collection("visitor").Doc(visitor.Uid).Get(ctx)
			if result.Data() != nil {
				utils.Throw("")
			}
			_, err := client.Collection("visitor").Doc(visitor.Uid).Set(ctx, visitor)
			if err != nil {
				utils.Throw("")
			}
			app.loggingMiddleWare(c, "LOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		},
		Tangkap: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Gas()

}
