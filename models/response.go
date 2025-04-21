package models

type Response struct {
	Status bool `json:"status"`
	Body interface{} `json:"body"`
	Message string `json:"message"`
}