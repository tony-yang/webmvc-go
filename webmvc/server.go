package webmvc

import (
  "fmt"
  "net/http"
)

// NewServer creates a new server instance and is the entry serving point
// of the framework
type NewServer struct {
  routes *Routes
}

func (s *NewServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  requestPath := r.URL.Path
  controller := s.routes.GetController(requestPath)
  response := *controller.Get()
  if response.StatusCode == 200 {
    fmt.Fprint(w, response.Body)
  } else {
    http.Error(w, response.Body, response.StatusCode)
  }
}
