package middleware

import (
	    //"errors"
	    "encoding/json"
	    "net/http"
	    "fmt"
)



type ServerError struct {
    StatusCode int
    Message string
}


func (e *ServerError) Error() string {
    return e.Message
}

func handleError(err error, writer http.ResponseWriter) {
    if ae, ok := err.(*ServerError); ok {
        response := ServerError{StatusCode: ae.StatusCode, Message: ae.Message}
        jsonResponse, _ := json.Marshal(response)
        writer.WriteHeader(ae.StatusCode)
        writer.Write(jsonResponse)
    }else {
        response := ServerError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("General error occured while setting up graphql handler")}
        jsonResponse, _ := json.Marshal(response)
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write(jsonResponse)
    }
}

