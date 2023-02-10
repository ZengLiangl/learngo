package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func err(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("customize error")
}
func noerr(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "阿斯顿"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "user error"},
	{errNoPermission, 403, "Forbidden"},
	{err, 500, "Internal Server Error"},
	{noerr, 200, "no error"},
}

func TestTestWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://imooc.com", nil)
		f(response, request)

		verifyResponse(t, response.Result(), tt.code, tt.message)
	}
}

func TestWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(t, resp, tt.code, tt.message)
	}
}

func verifyResponse(t *testing.T, resp *http.Response, expectCode int, expectMessage string) {
	b, _ := io.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectCode || body != expectMessage {
		t.Errorf("expect (%d, %s) got (%d, %s)", expectCode, expectMessage, resp.StatusCode, body)
	}
}
