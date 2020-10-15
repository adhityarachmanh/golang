package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modSkill() {

	app.Router.GET(utils.RouteAPI("skill"), app.getSkill)
	if config.MODE == "DEV" {
		app.Router.POST(utils.RouteAPI("add-skill"), app.addSkill)
	}

}
func (app *AppSchema) getSkill(c *gin.Context) {
	client, _ := app.Firebase.Firestore(ctx)

	var data []models.SkillSchema
	result := client.Collection("skill").Documents(ctx)
	for {
		var d models.SkillSchema
		doc, err := result.Next()

		if err != nil {
			break
		}
		doc.DataTo(&d)
		data = append(data, d)
	}
	defer client.Close()
	utils.ResponseAPI(c, models.ResponseSchema{Data: data})
}

func (app *AppSchema) addSkill(c *gin.Context) {

}
