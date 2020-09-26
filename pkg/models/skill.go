package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseSkillSchema struct {
	Kategori   string `json:"kategori"`
	Link       string `json:"link"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Color      string `json:"color"`
	Background string `json:"background"`
	Progress   int32  `json:"progress"`
}
type SkillSchema struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Kategori   string             `json:"kategori"`
	Link       string             `json:"link"`
	Name       string             `json:"name"`
	Image      string             `json:"image"`
	Color      string             `json:"color"`
	Background string             `json:"background"`
	Progress   int32              `json:"progress"`
}
