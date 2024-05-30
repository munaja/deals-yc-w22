package home

import (
	"net/http"

	td "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"
	lh "github.com/munaja/exam-deals-yc-w22/pkg/language-helper"

	e "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-general/error"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		e.ErrorResponse(w, r, http.StatusNotFound, lh.ErrorBundler("data-notFound"))
		return
	}

	hh.WriteJSON(w, http.StatusOK, td.Message{
		Message: "Welcome Dating App API!!",
	}, nil)
}
