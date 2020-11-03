// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modAuth() {
	app.routeRegister("POST", "auth/login", app.loginVisitor, false)
	app.routeRegister("POST", "auth/autologin", app.autoLoginVisitor, true)
	app.routeRegister("POST", "auth/updatevisitor", app.editVisitor, true)

}

func (app *AppSchema) editVisitor(c *gin.Context) {
	var visitorRequest models.VisitorRequest
	var visitor models.Visitor

	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &visitorRequest)
			visitor.Name = visitorRequest.Name
			visitor.Uid, _ = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)
			_, err := client.Collection("visitors").Doc(visitor.Uid).Update(ctx, []firestore.Update{
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
			result, _ := client.Collection("visitors").Doc(visitor.Uid).Get(ctx)
			result.DataTo(&visitor)

			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}
func (app *AppSchema) autoLoginVisitor(c *gin.Context) {
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			visitor.Uid, _ = app.getToken(c)
			app.firestoreGetDocument("visitors", visitor.Uid, visitor)
			app.loggingMiddleWare(c, "AUTOLOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}

func (app *AppSchema) loginVisitor(c *gin.Context) {

	var visitor models.Visitor
	utils.Block{
		Try: func() {
			visitor.Uid, _ = app.getToken(c)
			client, _ := app.Firebase.Firestore(ctx)

			result, _ := client.Collection("visitors").Doc(visitor.Uid).Get(ctx)
			if result.Data() != nil {
				utils.Throw("")
			}
			_, err := client.Collection("visitors").Doc(visitor.Uid).Set(ctx, visitor)
			if err != nil {
				utils.Throw("")
			}
			app.loggingMiddleWare(c, "LOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}
