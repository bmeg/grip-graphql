package middleware

import (
	    "encoding/json"
	    "net/http"
	    "fmt"

        "github.com/bmeg/grip/log"
)


type ServerError struct {
    StatusCode int
    Message string
}

func (e *ServerError) Error() string {
        return fmt.Sprintf("StatusCode: %d, Message: %s", e.StatusCode, e.Message)
}

func HandleError(err error, writer http.ResponseWriter) (error){
    if ae, ok := err.(*ServerError); ok {
        response := ServerError{StatusCode: ae.StatusCode, Message: ae.Message}
        log.Infof(ae.Error())
        jsonResponse, _ := json.Marshal(response)
        writer.WriteHeader(ae.StatusCode)
        writer.Write(jsonResponse)
        return err
    }
    log.Infof("Bad code ERROR, shouldn't have made it here, make sure to pass errors to error middleware data in form ServerError")
    return nil
}
