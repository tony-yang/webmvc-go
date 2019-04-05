package webmvc

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"webmvc/base"
	"webmvc/tester"
)

func TestServerHello(t *testing.T) {
	server := CreateNewServer()

	t.Run("controller with non-pre-defined HTTP code should return method not allowed", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/base", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Method Not Allowed"
		tester.AssertStringEqual(t, got, want)

		gotCode := response.Code
		wantCode := 405
		tester.AssertIntEqual(t, gotCode, wantCode)
	})

	t.Run("base controller should return 405 with method not allowed", func(t *testing.T) {
		server.Routes.RegisterRoute("/base", &base.Controller{})

		request, _ := http.NewRequest(http.MethodGet, "/base", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Method Not Allowed"
		tester.AssertStringEqual(t, got, want)

		gotCode := response.Code
		wantCode := 405
		tester.AssertIntEqual(t, gotCode, wantCode)
	})
}
