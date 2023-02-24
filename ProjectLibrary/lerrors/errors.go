package lerrors

import "errors"

var (
	// web? errors
	ErrInParam         = errors.New("error wrong param when calling the endpoint")
	ErrFile            = errors.New("error working with file")
	ErrMinio           = errors.New("error interacting with minio storage")
	ErrNoObject        = errors.New("no such object found in minio storage")
	ErrMarshallingJson = errors.New("error marshalling json")
	// db errors
	ErrNoRecord      = errors.New("no such record was found in database")
	ErrExecQuery     = errors.New("error executing query")
	ErrInQuery       = errors.New("error building query")
	ErrScanningQuery = errors.New("error scanning query results")
)
