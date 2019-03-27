package webmvc

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "webmvc/base"
  "webmvc/controllers"
)

func TestServerHello(t *testing.T) {
  routes := Routes{
    // map[string]string{
    //   "/index": "Hello Webmvc from Go",
    //   "/hello": "Hello World!",
    // },
    map[string]base.ControllerInterface{
      "/index": &controllers.Index{},
      "/hello": &controllers.Hello{},
    },
  }

  server := &NewServer{&routes}

  t.Run("return Hello Webmvc from index", func( t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/index", nil)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Body.String()
    want := "Hello Webmvc from Go"

    if got != want {
      t.Errorf("got '%s', want '%s'", got, want)
    }
  })

  t.Run("return Hello World from hello", func( t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Body.String()
    want := "Hello World!"

    if got != want {
      t.Errorf("got '%s', want '%s'", got, want)
    }
  })
}
