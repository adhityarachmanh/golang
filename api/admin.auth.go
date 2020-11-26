package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	// "time"
	// "fmt"
)

func (app *AppSchema) admin_auth() {
	app.routeRegister("POST", "admin/auth/login", app.admin_auth_login, false)
	app.routeRegister("POST", "admin/auth/autologin", app.admin_auto_login, true)
}

func (app *AppSchema) admin_auto_login(c *gin.Context) {
	var adminRequest models.AdminRequest
	var admin models.Admin
	var adminExists []models.Admin

	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &adminRequest)
			Token, Type := app.getToken(c)
			app.firestoreFilter(Type, Filter{Key: "token", Op: "==", Value: Token}, &adminExists)
			admin = adminExists[0]
			admin.Password = ""
			utils.ResponseAPI(c, models.ResponseSchema{Data: admin})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}

func (app *AppSchema) admin_auth_login(c *gin.Context) {
	var adminRequest models.AdminRequest
	var admin models.Admin
	var adminExists []models.Admin

	utils.Block{
		Try: func() {
			app.BindRequestJSON(c, &adminRequest)
			NewToken, Type := app.getToken(c)
			password := utils.HashAndSalt(adminRequest.Password)
			app.firestoreFilter("admins", Filter{Key: "username", Op: "==", Value: adminRequest.Username}, &adminExists)
			if len(adminExists) == 0 {
				utils.ResponseAPIError(c, "Username tidak terdaftar.")
				return
			}
			admin = adminExists[0]
			if admin.Password != password {
				utils.ResponseAPIError(c, "Password tidak benar.")
				return
			}
			app.firestoreUpdate(Type, admin.Uid, []firestore.Update{
				{
					Path: "token", Value: NewToken,
				},
			})
			admin.Password = ""
			utils.ResponseAPI(c, models.ResponseSchema{Data: admin})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
