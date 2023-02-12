package lerrors

import "errors"

var (
	ErrNoRecord        = errors.New("no such record was found in database")
	ErrMarshallingJson = errors.New("error marshalling json")
	ErrInQuery         = errors.New("error in database query")
	ErrInParam         = errors.New("error wrong param when calling the endpoint")
	ErrFile            = errors.New("error working with file")
)
