// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"time"
)

func (app *AppSchema) user_auth() {
	app.routeRegister("POST", "auth/login", app.user_auth_login_visitor, false)
	app.routeRegister("POST", "auth/autologin", app.user_auth_autologin_visitor, false)
	app.routeRegister("POST", "auth/updatevisitor", app.user_auth_edit_visitor, true)
	app.routeRegister("POST", "auth/notif", app.user_auth_notif, true)
	app.routeRegister("POST", "auth/chat/active", app.user_auth_chat_active, true)

}

func (app *AppSchema) user_auth_chat_active(c *gin.Context) {
	var adminList []models.Admin
	var chatRequest models.ActiveChat
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			client, _ := app.Firebase.Firestore(ctx)
			uid, _ := app.getToken(c)
			app.BindRequestJSON(c, &chatRequest)
			app.firestoreByCollection("admins", &adminList)
			app.firestoreUpdate("visitors", uid, []firestore.Update{
				{
					Path: "name", Value: chatRequest.Name,
				},
				{
					Path: "chat", Value: true,
				},
			})
			documentID := utils.UUID()
			client.Collection("visitors").Doc(uid).Collection("chating").Doc(documentID).Set(ctx, models.ChatingSchema{
				Id:               documentID,
				Text:             "Welcome to my website.",
				CreatedAt:        chatRequest.Time,
				CustomProperties: map[string]interface{}{},
				User: models.ChatingUserSchema{
					Color:          4294967295,
					ContainerColor: 4279858655,
					CustomProperties: models.CustomPropertiesSchema{
						Read: false,
					},
					FirstName: "Super",
					LastName:  "Admin",
					Name:      "superadmin",
				},
			})

			app.firestoreGetDocument("visitors", uid, &visitor)
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}

func (app *AppSchema) user_auth_notif(c *gin.Context) {
	var fmcRequest models.FCMRequest
	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &fmcRequest)
			status := app.sendNotifToAdmin(fmcRequest)
			utils.ResponseAPI(c, models.ResponseSchema{Data: status})

		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}

func (app *AppSchema) user_auth_edit_visitor(c *gin.Context) {
	var visitorRequest models.VisitorRequest
	var visitor models.Visitor
	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &visitorRequest)
			uid, _ := app.getToken(c)
			// loc, _ := time.LoadLocation("Asia/Jakarta")
			// chatID := utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).Format(time.RFC3339))
			app.firestoreGetDocument("visitors", uid, &visitor)
			app.firestoreUpdate("visitors", uid, []firestore.Update{
				{
					Path: "token", Value: visitorRequest.Token,
				},
				{
					Path: "ip_address", Value: visitorRequest.IPAddress,
				},
			})

			app.firestoreGetDocument("visitors", uid, &visitor)

			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}
func (app *AppSchema) user_auth_autologin_visitor(c *gin.Context) {
	var visitor models.Visitor
	// var findById models.Visitor
	var visitorExist []models.Visitor

	var visitorBanned []models.BannedVisitor
	var visitorRequest models.VisitorRequest
	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &visitorRequest)
			client, _ := app.Firebase.Firestore(ctx)
			uid, _ := app.getToken(c)
			loc, _ := time.LoadLocation("Asia/Jakarta")
			app.firestoreGetDocument("visitors", uid, &visitor)

			if visitor.Uid == "" {
				app.firestoreFilter("visitors", Filter{Key: "ip_address", Op: "==", Value: visitorRequest.IPAddress}, &visitorExist)
				if len(visitorExist) != 0 {
					uid = visitorExist[0].Uid
				} else {
					visitor.Uid = uid
					visitor.IPAddress = visitorRequest.IPAddress
					client.Collection("visitors").Doc(visitor.Uid).Set(ctx, visitor)
				}
			} else if visitor.Uid != "" && visitor.IPAddress != visitorRequest.IPAddress {
				app.firestoreUpdate("visitors", uid, []firestore.Update{
					{
						Path: "ip_address", Value: visitorRequest.IPAddress,
					},
				})
				app.firestoreFilter("banned", Filter{Key: "ip_address", Op: "==", Value: visitor.IPAddress}, &visitorBanned)
				if len(visitorBanned) != 0 {
					documentID := time.Now().In(loc).String()
					client.Collection("banned").Doc(documentID).Set(ctx, models.BannedVisitor{
						DocumentID: documentID,
						Uid:        uid,
						IPAddress:  visitorRequest.IPAddress,
					})
				}

				app.firestoreUpdate("visitors", uid, []firestore.Update{
					{
						Path: "time_visit", Value: time.Now().In(loc).Format(time.RFC3339),
					},
				})
			}
			app.firestoreGetDocument("visitors", uid, &visitor)
			// app.loggingMiddleWare(c, "AUTOLOGIN_SUCCESS")
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})
		}, Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()

}

func (app *AppSchema) user_auth_login_visitor(c *gin.Context) {
	var visitor models.Visitor
	var visitorExist []models.Visitor

	var visitorRequest models.VisitorRequest
	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &visitorRequest)
			client, _ := app.Firebase.Firestore(ctx)
			loc, _ := time.LoadLocation("Asia/Jakarta")
			uid, _ := app.getToken(c)
			app.firestoreFilter("visitors", Filter{Key: "ip_address", Op: "==", Value: visitorRequest.IPAddress}, &visitorExist)
			if len(visitorExist) != 0 {
				app.firestoreGetDocument("visitors", visitorExist[0].Uid, &visitor)
			} else {
				visitor.Uid = uid
				visitor.IPAddress = visitorRequest.IPAddress
				visitor.TimeVisit = time.Now().In(loc).Format(time.RFC3339)
				client.Collection("visitors").Doc(uid).Set(ctx, visitor)
				app.firestoreGetDocument("visitors", uid, &visitor)
				app.sendNotifToAdmin(models.FCMRequest{
					Title: "New Visitor",
					Body:  "IP Address : " + visitor.IPAddress,
					Data: map[string]interface{}{
						"uid":  visitor.Uid,
						"type": "NOTIFICATION_VISITOR",
					},
				})
			}
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})

		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
