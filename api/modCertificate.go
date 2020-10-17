package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modCertificate() {
	app.routeRegister("GET", "certificate", app.getCertificate)
}

func (app *AppSchema) getCertificate(c *gin.Context) {
	utils.Tahan{
		Coba: func() {
			client, _ := app.Firebase.Firestore(ctx)

			var data []models.CertificateSchema
			result := client.Collection("certificates").Documents(ctx)
			for {
				var d models.CertificateSchema
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
		Tangkap: func(e utils.Pengecualian) {
			utils.ResponseAPIError(c, "Server Error!")
		},
	}.Gas()

}
