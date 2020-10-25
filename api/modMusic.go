package api

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

func songMapping(song string) []interface{} {
	keyMap := map[string]int{
		"1": 24,
		"!": 25,
		"2": 26,
		"@": 27,
		"3": 28,
		"4": 29,
		"$": 30,
		"5": 31,
		"%": 32,
		"6": 33,
		"^": 34,
		"7": 35,
		"8": 36,
		"*": 37,
		"9": 38,
		"(": 39,
		"0": 40,
		"q": 41,
		"Q": 42,
		"w": 43,
		"W": 44,
		"e": 45,
		"E": 46,
		"r": 47,
		"t": 48,
		"T": 49,
		"y": 50,
		"Y": 51,
		"u": 52,
		"i": 53,
		"I": 54,
		"o": 55,
		"O": 56,
		"p": 57,
		"P": 58,
		"a": 59,
		"s": 60,
		"S": 61,
		"d": 62,
		"D": 63,
		"f": 64,
		"g": 65,
		"G": 66,
		"h": 67,
		"H": 68,
		"j": 69,
		"J": 70,
		"k": 71,
		"l": 72,
		"L": 73,
		"z": 74,
		"Z": 75,
		"x": 76,
		"c": 77,
		"C": 78,
		"v": 79,
		"V": 80,
		"b": 81,
		"B": 82,
		"n": 83,
		"m": 84,
		"_": 0,
		"|": -1,
	}
	var newArrResult []interface{}

	var isArray bool = false
	var arrayTemp []int
	arrSplit := strings.Split(song, "")

	for i := 0; i < len(arrSplit); i++ {
		d := arrSplit[i]

		if d == "[" {
			isArray = true
			// return
		} else if d == "]" {
			isArray = false

			newArrResult = append(newArrResult, arrayTemp)
			arrayTemp = nil
			// return
		} else {
			if isArray {
				c := keyMap[d]
				arrayTemp = append(arrayTemp, c)
			} else {
				c := keyMap[d]
				newArrResult = append(newArrResult, c)
			}
		}

	}
	return newArrResult
}
func (app *AppSchema) modMusic() {
	// app.routeRegister("POST", "music/add", app.addMusic)
	app.routeRegister("GET", "music", app.getMusic)
}

func (app *AppSchema) getMusic(c *gin.Context) {
	var data []models.MusicSchema
	var d models.MusicSchema
	utils.Tahan{
		Coba: func() {
			// app.loggingMiddleWare(c, "ACCESS_API")
			client, _ := app.Firebase.Firestore(ctx)
			result := client.Collection("music").OrderBy("title", firestore.Asc).Documents(ctx)
			for {
				doc, err := result.Next()
				if err != nil {
					break
				}
				doc.DataTo(&d)
				d.Song = songMapping(d.SongString)
				d.SongString = ""
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

func (app *AppSchema) addMusic(c *gin.Context) {
	var path string
	var req models.RequestSchema
	var data []models.MusicSchema
	c.BindJSON(&path)
	utils.Tahan{
		Coba: func() {
			jsonFile, _ := os.Open("/Users/adhityarachmanh/projects/CV/api/music.json")
			d, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(d, &req)
			ds, _ := json.Marshal(req.Data)
			json.Unmarshal(ds, &data)
			client, _ := app.Firebase.Firestore(ctx)
			for i := 0; i < len(data); i++ {
				dt := data[i]
				_, _, err := client.Collection("music").Add(ctx, dt)
				if err != nil {
					utils.ResponseAPIError(c, err.Error())
					return
				}
			}

			utils.ResponseAPI(c, models.ResponseSchema{Data: data})
		},
		Tangkap: func(e utils.Exception) {
			utils.ResponseAPIError(c, fmt.Sprintf("%s", e))
		},
	}.Gas()
}
