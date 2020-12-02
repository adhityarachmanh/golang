package models

type InformationSchema struct {
	Fullname string `json:"fullname" firestore:"fullname"`
	Name     string `json:"name" firestore:"name"`
	Phone    int    `json:"phone" firestore:"phone"`
	Email    string `jsonn:"email" firestore:"email"`
	Address  string `json:"address" firestore:"address"`
}
