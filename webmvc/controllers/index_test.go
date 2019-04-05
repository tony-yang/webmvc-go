package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "webmvc"
	"webmvc/tester"
)

func TestIndex(t *testing.T) {
  server := webmvc.CreateNewServer()
  server.Routes.RegisterRoute("/index", &Index{})

	t.Run("return Hello Webmvc from index", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/index", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Hello Webmvc from Go"
		tester.AssertStringEqual(t, got, want)

		gotCode := response.Code
		wantCode := http.StatusOK
		tester.AssertIntEqual(t, gotCode, wantCode)
	})
}
