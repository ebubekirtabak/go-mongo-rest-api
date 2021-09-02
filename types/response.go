package types

type Response struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

