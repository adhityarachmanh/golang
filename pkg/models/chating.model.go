package models

type ChatingSchema struct {
	Id                string            `json:"id" firestore:"id"`
	Text              string            `json:"text" firestore:"text"`
	User              ChatingUser       `json:"user" firestore:"user"`
	CostumeProperties CustomeProperties `json:"customProperties" firestore:"customProperties"`
	CreatedAt         string            `json:"createdAt" firestore:"createdAt"`
}

type ChatingUser struct {
	Firstname string `json:"firstName" firestore:"firstName"`
	Lastname  string `json:"lastName" firestore:"lastName"`
	Username  string `json:"name" firestore:"name"`
	Uid       string `json:"uid" firestore:"uid"`
	Read      bool   `json:"read" firestore:"read"`
}

type CustomeProperties struct {
	Fullname  string `json:"name" firestore:"name"`
	Read      bool   `json:"read" firestore:"read"`
	Uid       string `json:"uid" firestore:"uid"`
	MessageID string `json:"message_id" firestore:"message_id"`
}
