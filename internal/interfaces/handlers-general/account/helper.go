package account

import (
	"net/http"

	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"
)

func checkRepassword(password, repassword string, w http.ResponseWriter, r *http.Request) bool {
	if password != repassword {
		hh.WriteJSON(w, http.StatusUnauthorized, ds.II{
			"Meta":   ds.IS{"count": "1"},
			"Errors": es.XErrors{"repassword": es.XError{Code: "equalToField", Message: l.I.Msg("equalToField") + ": password"}},
		}, nil)
		return false
	}
	return true
}
