package models

type VendorSchema struct {
	ID    string `json:"id" firestore:"id"`
	Name  string `json:"name" firestore:"name"`
	Slug  string `json:"slug" firestore:"slug"`
	Image string `json:"image" firestore:"image"`
}

type AlbumProjectSchema struct {
	Title  string `json:"title" firestore:"title"`
	Source string `json:"source" firestore:"source"`
	Type   string `json:"type" firestore:"type"`
}

type ProjectSchema struct {
	ProjectId string               `json:"ProjectID" firestore:"ProjectID"`
	Logo      string               `json:"logo" firestore:"logo"`
	Detail    string               `json:"detail" firestore:"detail"`
	Title     string               `json:"title" firestore:"title"`
	StartAt   string               `json:"startAt" firestore:"startAt"`
	EndAt     string               `json:"endAt" firestore:"endAt"`
	CreatedAt string               `json:"createdAt" firestore:"createdAt"`
	Image     string               `json:"image" firestore:"image"`
	Url       string               `json:"url" firestore:"url"`
	Source    string               `json:"source" firestore:"source"`
	Type      string               `json:"type" firestore:"type"`
	Tools     []interface{}        `json:"tools" firestore:"tools"`
	Publish   interface{}          `json:"publish" firestore:"publish"`
	Vendors   []interface{}        `json:"vendors" firestore:"vendors"`
	Album     []AlbumProjectSchema `json:"album"`
}
