package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *AppSchema) modCertificate() {

	app.Router.GET(utils.RouteAPI("certificate"), app.getCertificate)

}

func (app *AppSchema) getCertificate(c *gin.Context) {

	var q_result []models.CertificateSchema
	cur, err := app.DB.Collection("certificates").Find(ctx, bson.D{})
	if err != nil {
		utils.ResponseAPIError(c, "Error")
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.CertificateSchema
		cur.Decode(&result)
		q_result = append(q_result, result)
	}
	utils.ResponseAPI(c, models.ResponseSchema{Data: q_result})

}
