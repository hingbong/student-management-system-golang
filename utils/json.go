package utils

type json struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorJsonWithMessage(message string) *json {
	j := &json{Code: 0, Message: message}
	return j
}

func SuccessJsonWithData(data interface{}) *json {
	j := &json{Code: 1, Data: data}
	return j
}

func SuccessJson() *json {
	j := &json{Code: 1}
	return j
}
