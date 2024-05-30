package auth

import (
	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
)

// just return the error code
func getAndCheck(input, condition any) (eCode string) {
	result := dg.I.Where(condition).Find(input)
	if result.Error != nil {
		return "fetch-fail"
	} else if result.RowsAffected == 0 {
		return "auth-login-incorrect"
	}

	return ""
}
