package handlers

type Envelope struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}

func Success(data interface{}) *Envelope {
	return &Envelope{
		Data:   data,
		Errors: []string{},
	}
}

func Failure(errors ...string) *Envelope {
	return &Envelope{
		Data:   struct{}{},
		Errors: errors,
	}
}
