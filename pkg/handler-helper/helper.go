package handlerhelper

import (
	"net/http"

	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
)

// Process error required for string
func requiredString(w http.ResponseWriter, fieldName, input string) bool {
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, ds.II{"errors": es.XErrors{
			fieldName: es.XError{
				Code:    "val-required",
				Message: l.I.Msg("val-required"),
			},
		}}, nil)
		return false
	}
	return true
}
