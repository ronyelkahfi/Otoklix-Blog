package helpers

import "net/http"



func GenerateErrorCode(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNoContent
	case ErrDbServer:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
