package view

import (
	"net/http"

	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/view"
	sau "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
	s "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/view"
)

func GetDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	payload := m.GetDetailDto{
		Viewer_User_Id:    authInfo.User_Id,
		Viewer_Profile_Id: authInfo.Profile_Id,
	}
	res, err := s.GetDetail(payload)
	hh.DataResponse(w, res, err)
}
