// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (app *AppSchema) user_auth() {
	app.routeRegister("POST", "auth/login", app.user_auth_login_visitor, false)
	app.routeRegister("POST", "auth/autologin", app.user_auth_autologin_visitor, true)
	app.routeRegister("POST", "auth/updatevisitor", app.user_auth_edit_visitor, true)

}

func (app *AppSchema) user_auth_edit_visitor(c *gin.Context) {
	var visitorRequest models.VisitorRequest
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &visitorRequest)
			uid, _ := app.getToken(c)
			loc, _ := time.LoadLocation("Asia/Jakarta")
			// chatID := utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).Format(time.RFC3339))
			client, _ := app.Firebase.Firestore(ctx)
			result, _ := client.Collection("visitors").Doc(uid).Get(ctx)
			result.DataTo(&visitor)
			_, err := client.Collection("visitors").Doc(uid).Update(ctx, []firestore.Update{
				{
					Path: "name", Value: visitorRequest.Name,
				},
				{
					Path: "chat", Value: visitorRequest.Chat,
				},
				{
					Path: "token", Value: visitorRequest.Token,
				},
			})
			if err != nil {
				utils.Throw("")
			}

			messageID := fmt.Sprint(time.Now().In(loc).Unix())
			if !visitor.Chat {
				_, err = client.Collection("visitors").Doc(uid).Collection("chating").Doc(messageID).Set(ctx, models.ChatingSchema{
					CostumeProperties: models.CustomeProperties{
						Uid:       uid,
						Fullname:  visitor.Name,
						Read:      false,
						MessageID: messageID,
					},
					User: models.ChatingUser{
						Uid:       "",
						Username:  "superadmin",
						Firstname: "Super",
						Lastname:  "Admin",
						Read:      false,
					},
					CreatedAt: time.Now().In(loc).Format(time.RFC3339),
					Text:      "Welcome to my Website",
					Id:        messageID,
				})
			}

			result, _ = client.Collection("visitors").Doc(uid).Get(ctx)
			result.DataTo(&visitor)

			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}
func (app *AppSchema) user_auth_autologin_visitor(c *gin.Context) {
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			visitor.Uid, _ = app.getToken(c)
			app.firestoreGetDocument("visitors", visitor.Uid, &visitor)
			// app.loggingMiddleWare(c, "AUTOLOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}

func (app *AppSchema) user_auth_login_visitor(c *gin.Context) {

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
