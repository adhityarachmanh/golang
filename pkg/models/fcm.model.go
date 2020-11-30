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
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Data  map[string]string `json:"data"`
}

type NotificationSchema struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	ClickAction string `json:"click_action"`
}

type FCMResponse struct {
	Success int `json:"success"`
	Failure int `json:"failure"`
}
