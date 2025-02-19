package utils

import "strconv"

type ErrorCodeStruct struct {
	Titre      string
	Code       int
	ErrorTitle string
	Message    string
	IsLogin    bool
}

func (e *ErrorCodeStruct) Update(code int, errorTitle string, message string, ErrorToSend *bool) {
	e.Titre = strconv.Itoa(code) + " - " + errorTitle
	e.Code = code
	e.ErrorTitle = errorTitle
	e.Message = message
	*ErrorToSend = true
}
