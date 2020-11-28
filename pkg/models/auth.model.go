package models

import "time"

type BannedVisitor struct {
	IPAddress string `json:"ip_address" firestore:"ip_address"`
}

type Visitor struct {
	Uid       string `json:"uid" firestore:"uid"`
	Name      string `json:"name" firestore:"name"`
	Chat      bool   `json:"chat" firestore:"chat"`
	Token     string `json:"token" firestore:"token"`
	IPAddress string `json:"ip_address" firestore:"ip_address"`
}
type VisitorRequest struct {
	Name      string `json:"name" firestore:"name"`
	Chat      bool   `json:"chat" firestore:"chat"`
	Token     string `json:"token" firestore:"token"`
	IPAddress string `json:"ip_address" firestore:"ip_address"`
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
	Username string `json:"username" firestore:"username"`
	Password string `json:"password" firestore:"password"`
}

type Admin struct {
	Uid      string `json:"uid" firestore:"uid"`
	Name     string `json:"name" firestore:"name"`
	Username string `json:"username" firestore:"username"`
	Password string `json:"password" firestore:"password"`
	Token    string `json:"token" firestore:"token"`
	Image    string `json:"image" firestore:"image"`
}
