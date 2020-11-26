package api

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
)

type Filter struct {
	Key   string
	Op    string
	Value string
}

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
