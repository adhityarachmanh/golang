// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modAuth() {
	app.routeClientRegister("POST", "auth/login", app.loginVisitor, false)
	app.routeClientRegister("POST", "auth/autologin", app.autoLoginVisitor, true)
	app.routeClientRegister("POST", "auth/updatevisitor", app.editVisitor, true)

}

func (app *AppSchema) editVisitor(c *gin.Context) {
	var visitorRequest models.VisitorRequest
	var visitor models.Visitor

	utils.Block{
		Try: func() {
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
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Go()

}
func (app *AppSchema) autoLoginVisitor(c *gin.Context) {
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			visitor.Uid = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)
			result, _ := client.Collection("visitor").Doc(visitor.Uid).Get(ctx)
			result.DataTo(&visitor)
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Go()

}

func (app *AppSchema) loginVisitor(c *gin.Context) {

	var visitor models.Visitor
	utils.Block{
		Try: func() {
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
			app.loggingMiddleWare(c, "visitor", "LOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Go()

}
