// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modCertificate() {
	app.routeRegister("POST", "certificate", app.getCertificate, true)
}

func (app *AppSchema) getCertificate(c *gin.Context) {
	var data []models.CertificateSchema
	var d models.CertificateSchema
	utils.Block{
		Try: func() {
			client, _ := app.Firebase.Firestore(ctx)
			result := client.Collection("certificates").Documents(ctx)
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
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Go()

}
