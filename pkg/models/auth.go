package models

import "time"

type Chatting struct {
	Message string `json:"msg" firestore:"msg"`
}
type ChattingAdd struct {
	Message   string `json:"msg" firestore:"msg"`
	Arh       bool   `json:"arh" firestore:"arh"`
	CreatedAt string `json:"createdAt" firestore:"createdAt"`
	Read      bool   `json:"read" firestore:"read"`
	ChatID    string `json:"chatID" firestore:"chatID"`
}

type Visitor struct {
	Uid  string `json:"uid" firestore:"uid"`
	Name string `json:"name" firestore:"name"`
	Chat bool   `json:"chat" firestore:"chat"`
}
type VisitorRequest struct {
	Name string `json:"name" firestore:"name"`
}

type Logging struct {
	URL        string    `json:"url" firestore:"url"`
	IPAddress  string    `json:"ip" firestore:"ip"`
	MacAddress []string  `json:"macaddress" firestore:"macaddress"`
	Message    string    `json:"message" firestore:"message"`
	CreatedAt  time.Time `json:"createdAt" firestore:"createdAt"`
	UserAgent  string    `json:"useragent" firestore:"useragent"`
}
