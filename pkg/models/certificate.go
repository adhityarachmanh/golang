package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseCertificateSchema struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Kategori string `json:"kategori"`
}
type CertificateSchema struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name"`
	Image    string             `json:"image"`
	Kategori string             `json:"kategori"`
}
