package models

import "time"

type Visitor struct {
	Uid   string `json:"uid" firestore:"uid"`
	Name  string `json:"name" firestore:"name"`
	Chat  bool   `json:"chat" firestore:"chat"`
	Token string `json:"token" firestore:"token"`
}
type VisitorRequest struct {
	Name  string `json:"name" firestore:"name"`
	Chat  bool   `json:"chat" firestore:"chat"`
	Token string `json:"token" firestore:"token"`
}

type Logging struct {
	URL        string    `json:"url" firestore:"url"`
	IPAddress  string    `json:"ip" firestore:"ip"`
	MacAddress []string  `json:"macaddress" firestore:"macaddress"`
	Message    string    `json:"message" firestore:"message"`
	CreatedAt  time.Time `json:"createdAt" firestore:"createdAt"`
	UserAgent  string    `json:"useragent" firestore:"useragent"`
}
