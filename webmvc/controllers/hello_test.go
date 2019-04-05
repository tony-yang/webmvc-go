package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "webmvc"
  "webmvc/tester"
)

func TestHello(t *testing.T) {
  server := webmvc.CreateNewServer()
  server.Routes.RegisterRoute("/hello", &Hello{})

	t.Run("return Hello World from hello", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Hello World!"
		tester.AssertStringEqual(t, got, want)

		gotCode := response.Code
		wantCode := http.StatusOK
		tester.AssertIntEqual(t, gotCode, wantCode)
	})
}
