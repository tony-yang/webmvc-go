package webmvc

import (
	"fmt"
	"net/http"
	"webmvc/base"
)

// NewServer creates a new server instance and is the entry serving point
// of the framework
type NewServer struct {
	Routes *base.Routes
}

func (s *NewServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	controller := s.Routes.GetController(requestPath)

	if controller != nil {
		base.Debug("Request path = ", requestPath)
		base.Debug("Controller = ", controller)
		response := controller.Get()
		if response.StatusCode == 200 {
			fmt.Fprint(w, response.Body)
		} else {
			http.Error(w, "Method Not Allowed", response.StatusCode)
		}
	} else {
		http.Error(w, "Method Not Allowed", 405)
	}
}

func CreateNewServer() *NewServer {
	routes := base.CreateNewRouter()
	server := &NewServer{routes}
	return server
}
