package authentication

import (
	"context"
	"net/http"

	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	hh "github.com/munaja/exam-deals-yc-w22/pkg/handler-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	s "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/authentication"
)

var Position m.Position

func Login(w http.ResponseWriter, r *http.Request) {
	var input m.LoginDto
	if !(hh.ValidateStructByIOR(w, r.Body, &input)) {
		return
	}

	input.Position = Position
	res, err := s.GenToken(input)
	if err != nil {
		hh.WriteJSON(w, http.StatusUnauthorized, ds.II{"errors": err}, nil)
	} else {
		hh.DataResponse(w, res, err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authInfo := context.Context.Value(r.Context(), "authInfo").(*s.AuthInfo)
	s.RevokeToken(authInfo.Uuid)
	hh.WriteJSON(w, http.StatusOK, ds.IS{"message": "logged out"}, nil)
}

func GuardMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessDetail, err := s.ExtractToken(r, s.AccessToken)
		if err != nil {
			hh.WriteJSON(w, http.StatusUnauthorized, err.(es.XError), nil)
			return
		}
		ctx := context.WithValue(r.Context(), "authInfo", accessDetail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
