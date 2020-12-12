// CREATOR : Adhitya Rachman H
package models

import "time"

type DatabaseSchema struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Db       string `json:"db"`
	Port     string `json:"port"`
	Uri      string `json:"uri"`
}

type KeySchema struct {
	Enc map[int]int
	Dec []int
}
type ResponseSchema struct {
	Status  int         `json:"s"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Creator string      `json:"creator"`
}

type RequestSchema struct {
	Data    interface{} `json:"data"`
	Creator string      `json:"creator"`
}
type RequestProdSchema struct {
	Data    string `json:"data"`
	Creator string `json:"creator"`
}
type Logging struct {
	URL        string    `json:"url" firestore:"url"`
	IPAddress  string    `json:"ip" firestore:"ip"`
	MacAddress []string  `json:"macaddress" firestore:"macaddress"`
	Message    string    `json:"message" firestore:"message"`
	CreatedAt  time.Time `json:"createdAt" firestore:"createdAt"`
	UserAgent  string    `json:"useragent" firestore:"useragent"`
}
