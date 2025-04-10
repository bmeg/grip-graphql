// pluginiface/handler.go
package pluginface

import "net/http"

type HTTPHandlerCreator func(client any, config map[string]string) (http.Handler, error)
