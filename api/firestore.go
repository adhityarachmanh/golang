package api

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
)

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

func (app *AppSchema) firestoreByCollection(collection string, bind interface{}) {
	var data []interface{}
	var d interface{}
	client, _ := app.Firebase.Firestore(ctx)
	result := client.Collection(collection).Documents(ctx)
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
