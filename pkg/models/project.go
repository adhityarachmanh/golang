package models

type VendorSchema struct {
	ID    string `json:"id" firestore:"id"`
	Name  string `json:"name" firestore:"name"`
	Slug  string `json:"slug" firestore:"slug"`
	Image string `json:"image" firestore:"image"`
}

type ProjectSchema struct {
	Title     string        `json:"title" firestore:"title"`
	CreatedAt string        `json:"createdAt" firestore:"createdAt"`
	Image     string        `json:"image" firestore:"image"`
	Url       string        `json:"url" firestore:"url"`
	Source    string        `json:"source" firestore:"source"`
	Type      string        `json:"type" firestore:"type"`
	Tools     []interface{} `json:"tools" firestore:"tools"`
	Vendors   []interface{} `json:"vendors" firestore:"vendors"`
}
