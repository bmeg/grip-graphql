package middleware

import ( 
    "net/http"
    "fmt"
    "github.com/bmeg/grip/gripql"
    "github.com/graphql-go/handler"
)

type GraphQLJS struct {
	client    gripql.Client
	gjHandler *handler.Handler
    gen3      bool
    graphName string
    config    string
}

func setup(gh *GraphQLJS, writer http.ResponseWriter, request *http.Request){
	//var handler *GraphQLJS
	//var ok bool
    fmt.Println("HANDLER", gh.gjHandler)

    /*
	if handler, ok = gh.gjHandler; ok {
		//Call the setup function. If nothing has changed it will return without doing anything
		err := handler.setup(request.Header)
        if err != nil{
            handleError(err, writer)
            return
        }
	} else {
        tokenCache := NewTokenCache()
		//Graph handler was not found, so we'll need to set it up
		var err error
		handler, err = newGraphHandler(graphName, gh.client, request.Header, tokenCache)
        if err != nil{
            handleError(err, writer)
            return
        }
		gh.handlers[graphName] = handler
	}
	if handler != nil && handler.gqlHandler != nil {
		handler.gqlHandler.ServeHTTP(writer, request)
	} else {
         response := ServerError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("General error occured while setting up graphql handler")}
         jsonResponse, _ := json.Marshal(response)
         writer.Write(jsonResponse)
	}
    */
}
