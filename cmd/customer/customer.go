package main

import (
	a "github.com/munaja/exam-deals-yc-w22/pkg/api-core"
	d "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	h "github.com/munaja/exam-deals-yc-w22/pkg/api-core/http-std"
	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	r "github.com/munaja/exam-deals-yc-w22/pkg/api-core/ms-redis"

	ch "github.com/munaja/exam-deals-yc-w22/internal/interfaces/handlers-customer"
	uc "github.com/munaja/exam-deals-yc-w22/internal/use-cases/helper/config"
)

func main() {
	a.App.RegisterExtrCall(uc.SetConfig)
	a.Run(ch.SetRoutes(), &l.O, &d.O, &r.O, &h.O)
}
