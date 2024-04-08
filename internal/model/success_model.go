package models

type SuccessJSON struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
