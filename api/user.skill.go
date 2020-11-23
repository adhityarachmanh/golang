// CREATOR : Adhitya Rachman H

package api

import (
	// "arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) user_skill() {
	app.routeRegister("GET", "skill", app.getSkill, false)
	// if config.MODE == "DEV" {
	// 	app.routeRegister("POST", "add-skill", app.addSkill)
	// }

}
func (app *AppSchema) getSkill(c *gin.Context) {
	var data []models.SkillSchema
	utils.Block{
		Try: func() {
			app.firestoreByCollection("skill", &data)
			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}
