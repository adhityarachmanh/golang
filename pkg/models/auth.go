package models

import "time"

type Chatting struct {
	Message string `json:"msg" firestore:"msg"`
}

type Visitor struct {
	Name *string `json:"name" firestore:"name"`
	Chat bool    `json:"chat" firestore:"chat"`
}

type Logging struct {
	URL        string    `json:"url" firestore:"url"`
	IPAddress  string    `json:"ip" firestore:"ip"`
	MacAddress []string  `json:"macaddress" firestore:"macaddress"`
	Message    string    `json:"message" firestore:"message"`
	CreatedAt  time.Time `json:"cratedAt" firestore:"cratedAt"`
}
