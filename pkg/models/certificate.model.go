package models

import ()

type CertificateSchema struct {
	ID       string `json:"id" firestore:"id"`
	Name     string `json:"name" firestore:"name"`
	Image    string `json:"image" firestore:"image"`
	Kategori string `json:"kategori" firestore:"kategori"`
}
