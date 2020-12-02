package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "os"
	// "strings"
)

func (app *AppSchema) user_contact() {
	// app.routeRegister("POST", "music/add", app.addMusic)
	app.routeRegister("GET", "contact", app.getContact, false)
}

func (app *AppSchema) getContact(c *gin.Context) {
	var data models.InformationSchema
	utils.Block{
		Try: func() {
			app.firestoreGetDocument("config", "information", &data)
			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
