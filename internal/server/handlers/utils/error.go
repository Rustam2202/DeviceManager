package utils

type ErrorResponce struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
