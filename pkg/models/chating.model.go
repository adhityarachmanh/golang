package models

type ChatingSchema struct {
	Id               string                 `json:"id" firestore:"id"`
	Text             string                 `json:"text" firestore:"text"`
	CreatedAt        int                    `json:"createdAt" firestore:"createdAt"`
	CustomProperties map[string]interface{} `json:"customProperties" firestore:"customProperties"`
	User             ChatingUserSchema      `json:"user" firestore:"user"`
	// Image            string                 `json:"image" firestore:"image"`
	// Video            string                 `json:"video" firestore:"video"`
}

type ChatingUserSchema struct {
	Avatar           string                 `json:"avatar" firestore:"avatar"`
	Color            int                    `json:"color" firestore:"color"`
	ContainerColor   int                    `json:"containerColor" firestore:"containerColor"`
	CustomProperties CustomPropertiesSchema `json:"customProperties" firestore:"customProperties"`
	FirstName        string                 `json:"firstName" firestore:"firstName"`
	LastName         string                 `json:"lastName" firestore:"lastName"`
	Name             string                 `json:"name" firestore:"name"`
	Uid              string                 `json:"uid" firestore:"uid"`
}

type CustomPropertiesSchema struct {
	Read bool `json:"read" firestore:"read"`
}

type ActiveChat struct {
	Name string `json:"name" firestore:"name"`
	Time int    `json:"time" firestore:"time"`
}
