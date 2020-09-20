// CREATOR : Adhitya Rachman H
package models

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
}
