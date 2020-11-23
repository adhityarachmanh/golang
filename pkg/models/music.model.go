package models

type MusicSchema struct {
	Title  string `json:"title" firestore:"title"`
	Image  string `json:"image" firestore:"image"`
	Artist string `json:"artist" firestore:"artist"`

	Song  interface{} `json:"song" firestore:"song"`
	Tempo int         `json:"tempo" firestore:"tempo"`
}
