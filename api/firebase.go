package api

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"bytes"
	"cloud.google.com/go/firestore"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
	"time"
)

type Filter struct {
	Key   string
	Op    string
	Value string
}

func (app *AppSchema) initializeFirebase() {
	// initial firebase
	var dbconn interface{}
	utils.GetDecData("firebase", "", &dbconn)
	data, _ := json.Marshal(dbconn)
	opt := option.WithCredentialsJSON(data)
	Conn, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	app.Firebase = Conn
}
func (app *AppSchema) loggingMiddleWare(c *gin.Context, description string) {
	Token, Type := app.getToken(c)
	client, _ := app.Firebase.Firestore(ctx)
	var logging models.Logging
	loc, _ := time.LoadLocation("Asia/Jakarta")
	logging.MacAddress, _ = utils.GetMacAddr()
	logging.IPAddress = c.ClientIP()
	logging.Message = description
	logging.URL = c.Request.RequestURI
	logging.UserAgent = c.Request.UserAgent()
	logging.CreatedAt = time.Now().In(loc)
	if config.MODE == "PROD" {
		logging.URL = strings.ReplaceAll(c.Request.RequestURI, "/", "")
		logging.URL = strings.ReplaceAll(logging.URL, "."+strings.ToLower(config.CREATOR), "")
		logging.URL = utils.Ed.BNE(6, 1).Dec(logging.URL)
	}
	client.Collection(Type).Doc(Token).Collection("logging").Add(ctx, logging)

}
func (app *AppSchema) sendNotificationFCM(fcmData models.FCM) int {
	var FCMKey models.FCMKey
	utils.GetDecData("fcm", "", &FCMKey)
	requestBody, _ := json.Marshal(fcmData)
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	request, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(requestBody))
	if err != nil {
		return 1
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "key="+FCMKey.Key)

	resp, err := client.Do(request)
	if err != nil {
		return 1
	}
	var result models.FCMResponse
	json.NewDecoder(resp.Body).Decode(&result)

	// defer resp.Body.Close()
	if result.Success > 0 {
		return 0
	} else if result.Failure > 0 {
		return 1
	} else {
		return 1
	}

}

// func (app *AppSchema) sendNotifToAdmin(fmcRequest models.FCMRequest) []int {
// 	var adminList []models.Admin
// 	var status []int
// 	loc, _ := time.LoadLocation("Asia/Jakarta")
// 	app.firestoreByCollection("admins", &adminList)
// 	for i := 0; i < len(adminList); i++ {
// 		fmcRequest.Data["notificationID"] = utils.Ed.BNE(6, 2).Enc(time.Now().In(loc).String())
// 		admin := adminList[i]
// 		stat := app.sendNotificationFCM(models.FCM{
// 			Notification: models.NotificationSchema{
// 				Title:       fmcRequest.Title,
// 				Body:        fmcRequest.Body,
// 				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
// 			},
// 			Data: fmcRequest.Data,
// 			To:   admin.NotificationToken,
// 		})
// 		status = append(status, stat)
// 	}
// 	return status
// }

func (app *AppSchema) mappingDataFirestore(result *firestore.DocumentIterator, bind interface{}) {
	var data []interface{}
	var d interface{}
	for {
		doc, err := result.Next()
		if err != nil {
			break
		}
		doc.DataTo(&d)
		data = append(data, d)
	}

	jsonData, _ := json.Marshal(data)
	json.Unmarshal(jsonData, &bind)
}

func (app *AppSchema) firestoreGetDocument(collection string, documentID string, bind interface{}) {
	client, _ := app.Firebase.Firestore(ctx)
	result, _ := client.Collection(collection).Doc(documentID).Get(ctx)

	result.DataTo(&bind)
}

func (app *AppSchema) firestoreFilter(collection string, params Filter, bind interface{}) {
	client, _ := app.Firebase.Firestore(ctx)
	result := client.Collection(collection).Where(params.Key, params.Op, params.Value).Documents(ctx)

	app.mappingDataFirestore(result, &bind)
}

func (app *AppSchema) firestoreByCollection(collection string, bind interface{}) {
	client, _ := app.Firebase.Firestore(ctx)
	result := client.Collection(collection).Documents(ctx)
	app.mappingDataFirestore(result, &bind)
}

func (app *AppSchema) firestoreUpdate(collection string, documentID string, update []firestore.Update) (*firestore.WriteResult, error) {
	client, _ := app.Firebase.Firestore(ctx)
	// result, err := client.Collection(collection).Doc(documentID).Update(ctx, update)
	return client.Collection(collection).Doc(documentID).Update(ctx, update)
}
