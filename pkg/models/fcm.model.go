package models

type FCMKey struct {
	Key string `json:"key"`
}
type FCM struct {
	Notification NotificationSchema `json:"notification"`
	To           string             `json:"to"`
	Data         interface{}        `json:"data"`
}

type FCMRequest struct {
	Message string `json:"message"`
}

type NotificationSchema struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type FCMResponse struct {
	Success int `json:"success"`
	Failure int `json:"failure"`
}
