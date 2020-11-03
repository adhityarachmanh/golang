package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "os"
	// "strings"
)

func (app *AppSchema) modMusic() {
	// app.routeRegister("POST", "music/add", app.addMusic)
	app.routeRegister("GET", "music", app.getMusic, false)
}

func (app *AppSchema) getMusic(c *gin.Context) {
	var data []models.MusicSchema
	utils.Block{
		Try: func() {
			app.firestoreByCollection("music", &data)
			for i := 0; i < len(data); i++ {
				jd, _ := json.Marshal(data[i].Song)
				data[i].Song = utils.SongMapping(string(jd))
			}
			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
		},
		Catch: func(e utils.Exception) {
			utils.ResponseAPIError(c, "Something Wrong!")
		},
	}.Go()
}

// func (app *AppSchema) addMusic(c *gin.Context) {
// 	var path string
// 	var req models.RequestSchema
// 	var data []models.MusicSchema
// 	c.BindJSON(&path)
// 	utils.Block{
// 		Try: func() {
// 			jsonFile, _ := os.Open("/Users/adhityarachmanh/projects/CV/api/music.json")
// 			d, _ := ioutil.ReadAll(jsonFile)
// 			json.Unmarshal(d, &req)
// 			ds, _ := json.Marshal(req.Data)
// 			json.Unmarshal(ds, &data)
// 			client, _ := app.Firebase.Firestore(ctx)
// 			for i := 0; i < len(data); i++ {
// 				dt := data[i]
// 				_, _, err := client.Collection("music").Add(ctx, dt)
// 				if err != nil {
// 					utils.ResponseAPIError(c, err.Error())
// 					return
// 				}
// 			}

// 			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
// 		},
// 		Catch: func(e utils.Exception) {
// 			utils.ResponseAPIError(c, fmt.Sprintf("%s", e))
// 		},
// 	}.Go()
// }
