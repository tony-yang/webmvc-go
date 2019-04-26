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
	if s.Routes.Exists(requestPath) {
		controller, params := s.Routes.GetController(requestPath)
		base.Debug("Request path = ", requestPath)
		base.Debug("Controller = ", controller)
		base.Debug("Params = ", params)
		var response *base.HttpResponse
		switch r.Method {
		case http.MethodGet:
			response = controller.Get()
		case http.MethodPost:
			response = controller.Post()
		case http.MethodPut:
			response = controller.Put()
		case http.MethodDelete:
			response = controller.Delete()
		default:
			response = controller.Get()
		}

		if response.StatusCode == 200 {
			fmt.Fprint(w, response.Body)
		} else {
			http.Error(w, response.Body, response.StatusCode)
		}
	} else {
		http.Error(w, "Not Found", 404)
	}
}

func CreateNewServer() *NewServer {
	routes := base.CreateNewRouter()
	server := &NewServer{routes}
	return server
}
