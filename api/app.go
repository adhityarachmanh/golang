// CREATOR : Adhitya Rachman H

package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"context"
	"encoding/json"
	// "fmt"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
	"time"
)

var ctx = context.Background()
var Ed BNESchema = BNESchema{}

type BNESchema struct {
	Genkey func(int) string
	_k     string
	_bL    int
	Enc    func(string) string
	Dec    func(string) string
}
type AppSchema struct {
	Router   *gin.Engine
	Firebase *firebase.App
	ED       BNESchema
}

func (app *AppSchema) Initialize() {
	app.initializeFirebase()
	app.initializeRoutes()
}

func (app *AppSchema) initializeFirebase() {

	// initial database
	var dbconn interface{}
	utils.GetDecData("firebase", "", &dbconn)
	data, _ := json.Marshal(dbconn)
	opt := option.WithCredentialsJSON(data)
	// opt := option.WithCredentialsFile("key/firebase.json")
	Conn, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	app.Firebase = Conn
}

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.GET("/", app.index)
	app.modSkill()
	app.modCertificate()
	app.ED = BNESchema{}
}

// func (app *AppSchema) initializeSocketIO() {
// 	app.SocketIO, _ = socketio.NewServer(nil)
// 	app.modSocket()
// }

func (app *AppSchema) Run(addr string) {

	// app.Router.GET("/socket.io/*any", gin.WrapH(app.SocketIO))
	// app.Router.POST("/socket.io/*any", gin.WrapH(app.SocketIO))
	allowOrigin, allowMethods, allowedHeaders, Debug := config.GetCorsConfig()
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrigin,
		AllowedMethods:   allowMethods,
		AllowedHeaders:   allowedHeaders,
		AllowCredentials: true,
		Debug:            Debug,
	})
	srv := &http.Server{
		Handler:      c.Handler(app.Router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// go app.SocketIO.Serve()
	log.Fatal(srv.ListenAndServe())
}

func (app *AppSchema) index(c *gin.Context) {

	c.JSON(http.StatusOK, models.ResponseSchema{Status: 0, Message: "Hello Brow"})

}

func (self *BNESchema) BNE(bL int, s int) *BNESchema {
	k := utils.EdGenkey(s)
	if k == "" {
		self._k = utils.EdGenkey(-1)
	} else {
		self._k = k
	}
	self._bL = bL
	// untuk ENC

	self.Enc = func(s string) string {
		d := 0
		l := 0
		m := 8
		b := utils.Rank(2, self._bL) - 1
		r := ""
		j := len(s)
		i := 0
		for i < j {
			d = (d << m) + int(s[i])
			i += 1
			l += m
			for l >= self._bL {
				l -= self._bL
				r += string(self._k[(d>>l)&b])
				d &= utils.Rank(2, l) - 1
			}
		}
		if l > 0 {
			r += string(self._k[(d<<(self._bL-l))&b])
		}
		return r
	}

	self.Dec = func(s string) string {
		d := 0
		l := 0
		m := 8
		r := ""
		j := len(s)
		i := 0
		for i < j {
			d = ((d & 255) << self._bL) + strings.Index(self._k, string(s[i]))
			i += 1
			l += self._bL
			if l >= m {
				l -= m
				r += string(rune((d >> l) & 255))
			}
		}
		return r
	}

	//untuk dec
	return self
}
