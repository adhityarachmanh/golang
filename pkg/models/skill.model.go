package models

import ()

type SkillSchema struct {
	ID         string `json:"id" firestore:"id"`
	Kategori   string `json:"kategori" firestore:"kategori"`
	Link       string `json:"link" firestore:"link"`
	Name       string `json:"name" firestore:"name"`
	Image      string `json:"image" firestore:"image"`
	Color      string `json:"color" firestore:"color"`
	Background string `json:"background" firestore:"background"`
	Progress   int32  `json:"progress" firestore:"progress"`
}
