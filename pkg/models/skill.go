package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseSkillSchema struct {
	Kategori string `json:"kategori"`
	Link     string `json:"link"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}
type SkillSchema struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Kategori string             `json:"kategori"`
	Link     string             `json:"link"`
	Name     string             `json:"name"`
	Image    string             `json:"image"`
}
