package models

type MusicSchema struct {
	Title      string        `json:"title" firestore:"title"`
	Artist     string        `json:"artist" firestore:"artist"`
	SongString string        `firestore:"song"`
	Song       []interface{} `json:"song"`
	Tempo      int           `json:"tempo" firestore:"tempo"`
}
