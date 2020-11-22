package models

import ()

type ContactSchema struct {
	Icon string `json:"icon" firestore:"icon"`
	Name string `json:"name" firestore:"name"`
	Text string `json:"text" firestore:"text"`
}
