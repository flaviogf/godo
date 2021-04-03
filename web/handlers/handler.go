package handlers

import "html/template"

var Tmpl *template.Template

type Response struct {
	Data   interface{}
	Errors []string
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
