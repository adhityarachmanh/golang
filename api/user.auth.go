// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	// "fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (app *AppSchema) user_auth() {
	app.routeRegister("POST", "auth/login", app.user_auth_login_visitor, false)
	app.routeRegister("POST", "auth/autologin", app.user_auth_autologin_visitor, false)
	app.routeRegister("POST", "auth/updatevisitor", app.user_auth_edit_visitor, true)
	app.routeRegister("POST", "auth/notif", app.user_auth_notif, true)

}

func (app *AppSchema) user_auth_notif(c *gin.Context) {
	var adminList []models.Admin
	var visitor models.Visitor
	var fmcRequest models.FCMRequest
	var status []string
	utils.Block{
		Try: func() {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			notificationID := utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).String())
			uid, _ := app.getToken(c)
			app.BindRequestJSON(c, &fmcRequest)
			app.firestoreByCollection("admins", &adminList)
			app.firestoreGetDocument("visitors", uid, &visitor)

			for i := 0; i < len(adminList); i++ {
				admin := adminList[i]
				app.sendNotificationFCM(models.FCM{
					Notification: models.NotificationSchema{
						Title:       visitor.Name,
						Body:        fmcRequest.Message,
						ClickAction: "FLUTTER_NOTIFICATION_CLICK",
					},
					Data: map[string]string{
						"notificationID": notificationID,
						"uid":            visitor.Uid,
						"name":           visitor.Name,
					},
					To: admin.NotificationToken,
				})
				status = append(status, admin.NotificationToken)
			}
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
					Path: "name", Value: visitorRequest.Name,
				},
				{
					Path: "chat", Value: visitorRequest.Chat,
				},
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
				visitor.Uid = uid
				visitor.IPAddress = visitorRequest.IPAddress
				client.Collection("visitors").Doc(visitor.Uid).Set(ctx, visitor)
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
				app.firestoreGetDocument("visitors", uid, &visitor)
			}

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
			visitor.Uid, _ = app.getToken(c)
			app.firestoreFilter("visitors", Filter{Key: "ip_address", Op: "==", Value: visitorRequest.IPAddress}, &visitorExist)
			if len(visitorExist) != 0 {
				visitor = visitorExist[0]
			} else {
				visitor.IPAddress = visitorRequest.IPAddress
				client.Collection("visitors").Doc(visitor.Uid).Set(ctx, visitor)
			}
			utils.ResponseAPI(c, models.ResponseSchema{Data: visitor})

		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
