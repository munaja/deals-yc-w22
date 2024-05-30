package error

import (
	"encoding/json"
	"net/http"

	lh "github.com/munaja/exam-deals-yc-w22/pkg/language-helper"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	js, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(500)
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusNotFound, lh.ErrorBundler("data-notFound"))
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusMethodNotAllowed, lh.ErrorBundler("request-methodNotAllowed"))
}
