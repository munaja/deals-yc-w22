package account

import (
	"errors"
	"fmt"

	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	"gorm.io/gorm"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
)

func checkUserStatus(user *m.User, email string, status m.Status) error {
	err := dg.I.Where("Email = ?", email).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(fmt.Sprintf(l.I.Msg("data-notFound-condition"), "user", "email", email))
	} else if err != nil {
		return fmt.Errorf(l.I.Msg("data-fetch-fail"))
	} else if user.Status != nil && *(user.Status) != status {
		return fmt.Errorf(fmt.Sprintf(l.I.Msg("data-state-mismatch"), "user", m.GetUSText(status)))
	}
	return nil
}

func sterilizeUser(u *m.User) {
	u.Password = nil
	u.Status = nil
	u.LastSuccessLogin = nil
	u.LastAllowdLogin = nil
}
