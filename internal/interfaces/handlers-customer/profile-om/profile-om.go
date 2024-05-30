package profileom

import (
	"net/http"

	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	sau "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
	s "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/profile"
)

func GetDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	res, err := s.GetDetail(authInfo.User_Name, "of-mine")
	hh.DataResponse(w, res, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	var payload m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &payload) == false {
		return
	}

	res, err := s.UpdateByUserName(authInfo.User_Name, payload)
	hh.DataResponse(w, res, err)
}
