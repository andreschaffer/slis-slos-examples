package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		queryParams := req.URL.Query()

		var err error
		fail := false

		failParam := queryParams.Get("fail")
		if failParam != "" {
			fail, err = strconv.ParseBool(failParam)
			if err != nil {
				http.Error(w, "invalid query param 'fail' value", http.StatusBadRequest)
				return
			}
		}

		if fail {
			http.Error(w, "failed as requested!", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "pong")
		}
	})
}
