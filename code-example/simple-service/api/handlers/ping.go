package handlers

import (
	"fmt"
	"net/http"
)

func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})
}
