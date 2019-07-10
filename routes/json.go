package routes

type json struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func errorJson() *json {
	j := &json{Code: 0}
	return j
}
func errorJsonWithMessage(message string) *json {
	j := &json{Code: 0, Message: message}
	return j
}

func successJsonWithData(data interface{}) *json {
	j := &json{Code: 1, Data: data}
	return j
}

func successJson() *json {
	j := &json{Code: 1}
	return j
}
