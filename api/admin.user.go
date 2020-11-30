package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	"time"
	// "fmt"
)

func (app *AppSchema) admin_user() {
	app.routeRegister("POST", "admin/user/banned", app.admin_user_banned, true)
}

func (app *AppSchema) admin_user_banned(c *gin.Context) {
	var visitor models.VisitorRequest
	var visitorBanned []models.BannedVisitor
	utils.Block{
		Try: func() {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			client, _ := app.Firebase.Firestore(ctx)
			app.BindRequestJSON(c, &visitor)
			if !visitor.Banned {
				documentID := utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).String())
				client.Collection("banned").Doc(documentID).Set(ctx, models.BannedVisitor{
					DocumentID: documentID,
					Uid:        visitor.Uid,
					IPAddress:  visitor.IPAddress,
				})
			} else {
				app.firestoreFilter("banned", Filter{Key: "uid", Op: "==", Value: visitor.Uid}, &visitorBanned)
				for i := 0; i < len(visitorBanned); i++ {
					data := visitorBanned[i]
					client.Collection("banned").Doc(data.DocumentID).Delete(ctx)
				}
			}
			client.Collection("visitors").Doc(visitor.Uid).Update(ctx, []firestore.Update{
				{
					Path: "banned", Value: !visitor.Banned,
				},
			})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
