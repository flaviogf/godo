package handlers

type Response struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}

func Success(data interface{}) *Response {
	return &Response{
		Data:   data,
		Errors: []string{},
	}
}

func Failure(errors ...string) *Response {
	return &Response{
		Data:   struct{}{},
		Errors: errors,
	}
}
