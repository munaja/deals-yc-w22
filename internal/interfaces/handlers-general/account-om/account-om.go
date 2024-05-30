package accountom

import (
	"fmt"
	"net/http"

	td "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	sac "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/account"
	sau "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var input m.ChangePassDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	ctx := r.Context()
	fmt.Println(ctx)
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	res, err := sac.ChangePassword(authInfo.User_Id, input)
	hh.DataResponse(w, res, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	if !ok {
		hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	data := td.II{
		"data": authInfo,
	}
	hh.WriteJSON(w, http.StatusOK, data, nil)
}
