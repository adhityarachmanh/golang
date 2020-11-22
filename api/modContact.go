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

func (app *AppSchema) modContact() {
	// app.routeRegister("POST", "music/add", app.addMusic)
	app.routeRegister("GET", "contact", app.getContact, false)
}

func (app *AppSchema) getContact(c *gin.Context) {
	var data []models.ContactSchema
	var d models.ContactSchema
	utils.Block{
		Try: func() {
			client, _ := app.Firebase.Firestore(ctx)
			result := client.Collection("contact").Documents(ctx)
			for {
				doc, err := result.Next()
				if err != nil {
					break
				}
				doc.DataTo(&d)
				data = append(data, d)
			}
			defer client.Close()
			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
