package api

import (
	// "arh/pkg/models"
	// "arh/pkg/utils"
	// "cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	// "time"
)

func (app *AppSchema) admin_auth() {
	app.routeRegister("POST", "admin/auth/login", app.admin_auth_login, false)

}
func (app *AppSchema) admin_auth_login(c *gin.Context) {

}
