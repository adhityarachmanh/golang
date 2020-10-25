// CREATOR : Adhitya Rachman H

package api

import (
	// "arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modSkill() {
	app.routeRegister("POST", "skill", app.getSkill)
	// if config.MODE == "DEV" {
	// 	app.routeRegister("POST", "add-skill", app.addSkill)
	// }

}
func (app *AppSchema) getSkill(c *gin.Context) {
	var data []models.SkillSchema
	var d models.SkillSchema
	utils.Tahan{
		Coba: func() {
			client, _ := app.Firebase.Firestore(ctx)
			result := client.Collection("skill").Documents(ctx)
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
		Tangkap: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Gas()
}

// func (app *AppSchema) addSkill(c *gin.Context) {

// }
