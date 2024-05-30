package profile

import (
	"net/http"

	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	s "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/profile"
)

func GetDetail(w http.ResponseWriter, r *http.Request) {
	user_name := r.PathValue("user_name")
	if user_name == "" {
		return
	}

	res, err := s.GetDetail(user_name)
	hh.DataResponse(w, res, err)
}
