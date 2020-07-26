package model

type ApiError struct {
	Code int `json:"-"`
	Message string `json:"message"`
	Error error `json:"-"`
}
