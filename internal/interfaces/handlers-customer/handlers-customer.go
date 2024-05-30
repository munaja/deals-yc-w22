package view

import (
	"net/http"

	"github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-general/home"

	ac "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-general/account"
	aco "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-general/account-om"
	au "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-general/authentication"

	pf "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer/profile"
	pfo "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer/profile-om"
	s "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer/subscription"
	v "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer/view"
	vr "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer/view-result"
)

func SetRoutes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/", home.Index)

	r.HandleFunc("POST /authentication/login", au.Login)
	r.HandleFunc("POST /authentication/logout", au.Logout)

	r.HandleFunc("POST /account/register", ac.Register)
	r.HandleFunc("POST /account/request-confirm-by-email", ac.RequestConfirmationEmail)
	r.HandleFunc("GET /account/confirm-by-email", ac.ConfirmByEmail)
	r.HandleFunc("POST /account/request-reset-password", ac.RequestResetPassword)
	r.HandleFunc("GET /account/check-reset-password", ac.CheckResetPassword)
	r.HandleFunc("PATCH /account/reset-password", ac.ResetPassword)

	r.Handle("PATCH /account/change-password", au.GuardMW(http.HandlerFunc(aco.ChangePassword)))

	r.HandleFunc("PATCH /profile/{user_name}", pf.GetDetail)

	r.Handle("GET /profile-om", au.GuardMW(http.HandlerFunc(pfo.GetDetail)))
	r.Handle("PATCH /profile-om", au.GuardMW(http.HandlerFunc(pfo.Update)))

	r.Handle("POST /subscription", au.GuardMW(http.HandlerFunc(s.Create)))

	r.Handle("GET /view", au.GuardMW(http.HandlerFunc(v.GetDetail)))
	r.Handle("POST /view-result", au.GuardMW(http.HandlerFunc(vr.Create)))

	return r
}
