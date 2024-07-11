package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bmeg/grip/log"
)

type ServerError struct {
	StatusCode int
	Message    string
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("StatusCode: %d, Message: %s", e.StatusCode, e.Message)
}

func HandleError(err error, writer http.ResponseWriter) error {
	if ae, ok := err.(*ServerError); ok {
		response := ServerError{StatusCode: ae.StatusCode, Message: ae.Message}
		log.Infof(ae.Error())
		jsonResponse, _ := json.Marshal(response)
		writer.WriteHeader(ae.StatusCode)
		writer.Write(jsonResponse)
		return err
	}
	log.Infof("Uncaught Error: %s", err)
	return nil
}
