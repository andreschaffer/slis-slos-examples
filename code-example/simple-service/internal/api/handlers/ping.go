package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fail, err := parseFailParam(req)
		if err != nil {
			http.Error(w, "invalid query param 'fail' value", http.StatusBadRequest)
			return
		}

		sleep, err := parseSleepParam(req)
		if err != nil {
			http.Error(w, "invalid query param 'sleepMs' value", http.StatusBadRequest)
			return
		}

		time.Sleep(sleep)

		if fail {
			http.Error(w, "no pong, failed as requested!", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "pong!")
		}
	})
}

func parseFailParam(req *http.Request) (bool, error) {
	failParam := req.URL.Query().Get("fail")

	if failParam == "" {
		return false, nil
	}

	return strconv.ParseBool(failParam)
}

func parseSleepParam(req *http.Request) (time.Duration, error) {
	sleepParam := req.URL.Query().Get("sleep")

	if sleepParam == "" {
		return 0, nil
	}

	return time.ParseDuration(sleepParam)
}
