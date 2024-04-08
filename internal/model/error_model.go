package models

type ErrorJSON struct {
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
