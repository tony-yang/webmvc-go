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
    map[string]base.ControllerInterface{
      "/index": &controllers.Index{},
      "/hello": &controllers.Hello{},
      "/base": &base.Controller{},
    },
  }

  server := &NewServer{&routes}

  t.Run("return Hello Webmvc from index", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/index", nil)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Body.String()
    want := "Hello Webmvc from Go"
    AssertStringEqual(t, got, want)

    gotCode := response.Code
    wantCode := http.StatusOK
    AssertIntEqual(t, gotCode, wantCode)
  })

  t.Run("return Hello World from hello", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Body.String()
    want := "Hello World!"
    AssertStringEqual(t, got, want)

    gotCode := response.Code
    wantCode := http.StatusOK
    AssertIntEqual(t, gotCode, wantCode)
  })

  t.Run("base controller should return 405 with Method not allowed", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/base", nil)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Body.String()
    want := "Method not allowed\n"
    AssertStringEqual(t, got, want)

    gotCode := response.Code
    wantCode := 405
    AssertIntEqual(t, gotCode, wantCode)
  })
}

func AssertStringEqual(t *testing.T, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got string '%s', want '%s'", got, want)
  }
}

func AssertIntEqual(t *testing.T, got, want int) {
  t.Helper()
  if got != want {
    t.Errorf("got int '%d', want '%d'", got, want)
  }
}
