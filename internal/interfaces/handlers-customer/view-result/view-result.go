package viewresult

import (
	"net/http"

	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/view-result"
	sau "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
	s "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/view-result"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	var payload m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &payload) == false {
		return
	}

	payload.Viewer_Profile_Id = authInfo.Profile_Id
	res, err := s.Create(payload)
	hh.DataResponse(w, res, err)

}
