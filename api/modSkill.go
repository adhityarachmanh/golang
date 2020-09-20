package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *AppSchema) modSkill() {

	app.Router.GET(utils.RouteAPI("skill"), app.getSkill)
	app.Router.POST(utils.RouteAPI("add-skill"), app.addSkill)

}
func (app *AppSchema) getSkill(c *gin.Context) {
	var q_result []models.SkillSchema
	cur, err := app.DB.Collection("skills").Find(ctx, bson.D{})
	if err != nil {
		utils.ResponseAPIError(c, "Error")
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.SkillSchema
		cur.Decode(&result)
		q_result = append(q_result, result)
	}
	utils.ResponseAPI(c, models.ResponseSchema{Data: q_result})
}

func (app *AppSchema) addSkill(c *gin.Context) {
	// var q_result models.SkillSchema
	var request models.BaseSkillSchema
	c.BindJSON(&request)
	_, err := app.DB.Collection("skills").InsertOne(ctx, request)
	if err != nil {
		utils.ResponseAPIError(c, "Error")
		return
	}
	utils.ResponseAPI(c, models.ResponseSchema{Message: "Data berhasil di input."})
}
