package models

import "time"

type BannedVisitor struct {
	DocumentID string `json:"document_id" firestore:"document_id"`
	Uid        string `json:"uid" firestore:"uid"`
	IPAddress  string `json:"ip_address" firestore:"ip_address"`
	UserAgent  string `json:"user_agent" firestore:"user_agent"`
}

type Visitor struct {
	Uid       string `json:"uid" firestore:"uid"`
	Name      string `json:"name" firestore:"name"`
	Chat      bool   `json:"chat" firestore:"chat"`
	Token     string `json:"token" firestore:"token"`
	IPAddress string `json:"ip_address" firestore:"ip_address"`
	UserAgent string `json:"user_agent" firestore:"user_agent"`
	TimeVisit string `json:"time_visit" firestore:"time_visit"`
}
type VisitorRequest struct {
	Uid       string `json:"uid" firestore:"uid"`
	Name      string `json:"name" firestore:"name"`
	Chat      bool   `json:"chat" firestore:"chat"`
	Token     string `json:"token" firestore:"token"`
	IPAddress string `json:"ip_address" firestore:"ip_address"`
	UserAgent string `json:"user_agent" firestore:"user_agent"`
	Banned    bool   `json:"banned" firestore:"banned"`
}

type Logging struct {
	URL        string    `json:"url" firestore:"url"`
	IPAddress  string    `json:"ip" firestore:"ip"`
	MacAddress []string  `json:"macaddress" firestore:"macaddress"`
	Message    string    `json:"message" firestore:"message"`
	CreatedAt  time.Time `json:"createdAt" firestore:"createdAt"`
	UserAgent  string    `json:"useragent" firestore:"useragent"`
}

type AdminRequest struct {
	Username          string `json:"username" firestore:"username"`
	Password          string `json:"password" firestore:"password"`
	NotificationToken string `json:"notificationToken" firestore:"notificationToken"`
}

type Admin struct {
	Uid               string `json:"uid" firestore:"uid"`
	Name              string `json:"name" firestore:"name"`
	Username          string `json:"username" firestore:"username"`
	Password          string `json:"password" firestore:"password"`
	Token             string `json:"token" firestore:"token"`
	Image             string `json:"image" firestore:"image"`
	Pin               int64  `json:"pin" firestore:"pin"`
	NotificationToken string `json:"notificationToken" firestore:"notificationToken"`
}
