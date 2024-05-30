package account

import (
	"net/http"

	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	sac "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/account"
	sau "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input m.RegisterDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	res, err := sac.Register(input)
	hh.DataResponse(w, res, err)
}

func ConfirmByEmail(w http.ResponseWriter, r *http.Request) {
	var input m.ResendConfirmationEmailDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	res, err := sac.ConfirmByEmail(input)
	hh.DataResponse(w, res, err)
}

func RequestConfirmationEmail(w http.ResponseWriter, r *http.Request) {
	var input m.ResendEmailConfirmDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	res, err := sac.ResendConfirmationEmail(input)
	hh.DataResponse(w, res, err)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var input m.ChangePassDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	res, err := sac.ChangePassword(authInfo.User_Id, input)
	hh.DataResponse(w, res, err)
}

func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	var input m.RequestResetPassDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	res, err := sac.RequestResetPass(input)
	hh.DataResponse(w, res, err)
}

func CheckResetPassword(w http.ResponseWriter, r *http.Request) {
	var input m.CheckResetPassDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	res, err := sac.CheckResetPass(input)
	hh.DataResponse(w, res, err)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var input1 m.CheckResetPassDto
	if hh.ValidateStructByURL(w, *r.URL, &input1) == false {
		return
	}
	var input2 m.ResetPassDto
	if hh.ValidateStructByIOR(w, r.Body, &input2) == false {
		return
	}

	res, err := sac.ResetPass(input1, input2)
	hh.DataResponse(w, res, err)
}

func Check(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	data := ds.II{
		"data": authInfo,
	}
	hh.WriteJSON(w, http.StatusOK, data, nil)
}
