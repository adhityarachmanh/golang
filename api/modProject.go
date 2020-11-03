package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	// "cloud.google.com/go/firestore"

	"github.com/gin-gonic/gin"
)

func (app *AppSchema) modProject() {
	app.routeRegister("GET", "projects", app.getProject, false)
}

func (app *AppSchema) getProject(c *gin.Context) {
	var projects []models.ProjectSchema
	var skills []models.SkillSchema
	var vendors []models.VendorSchema

	utils.Block{
		Try: func() {
			// init data
			app.firestoreByCollection("project", &projects)
			app.firestoreByCollection("skill", &skills)
			app.firestoreByCollection("vendor", &vendors)

			for i := 0; i < len(projects); i++ {
				project := projects[i]
				for j := 0; j < len(project.Tools); j++ {
					tool := project.Tools[j]
					for k := 0; k < len(skills); k++ {
						skill := skills[k]
						if skill.Name == tool {
							projects[i].Tools[j] = skill
						}
					}
				}
				for j := 0; j < len(project.Vendors); j++ {
					vendor := project.Vendors[j]
					for k := 0; k < len(vendors); k++ {
						_vendor := vendors[k]
						if _vendor.Slug == vendor {
							projects[i].Vendors[j] = _vendor
						}
					}
				}
			}

			utils.ResponseAPI(c, models.ResponseSchema{Data: projects})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Telah terjadi kesalahan!")
		},
	}.Go()
}
