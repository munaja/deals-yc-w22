package account

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	ac "github.com/munaja/exam-deals-yc-w22/pkg/api-core"
	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	p "github.com/munaja/exam-deals-yc-w22/pkg/password"
	sh "github.com/munaja/exam-deals-yc-w22/pkg/use-case-helper"
	"gorm.io/gorm"

	mp "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	mut "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user-token"
	pf "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/profile"
	sut "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/user-token"
)

const source = "user"

// Self create account
func Register(input m.RegisterDto) (*ds.Data, error) {
	var user m.User
	if err := sc.Copy(&user, &input); err != nil {
		return nil, es.XErrors{"struct": es.XError{Code: "copy-fail", Message: l.I.Msg("data-copy-fail")}}
	}

	password, err := p.Hash(*user.Password)
	if err != nil {
		return nil, es.XErrors{"data": es.XError{Code: "process-fail", Message: l.I.Msg("data-process-fail")}}
	} else {
		user.Password = &password
	}

	err = dg.I.Where(&m.RegisterDto{Email: input.Email}).Or(&m.RegisterDto{Name: input.Name}).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, es.XErrors{"data": es.XError{Code: "update-fail", Message: l.I.Msg("data-update-fail")}}
	}

	xerr := es.XErrors{}
	if user.Id > 0 {
		if user.Name == input.Name {
			xerr["name"] = es.XError{Code: "registered", Message: l.I.Msg("registered")}
		}
		if user.Email == input.Email {
			xerr["email"] = es.XError{Code: "registered", Message: l.I.Msg("registered")}
		}
		if len(xerr) > 0 {
			return nil, xerr
		}
	}

	if err = dg.I.Transaction(func(tx *gorm.DB) error {
		StatusNew := m.USNew
		user.Status = &StatusNew
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		var userToken mut.UserToken
		if err := sut.Request(&userToken, input.Email, mut.UTConfirmByEmail, 60*24*3, 5); err != nil {
			return err
		} else if err := tx.Save(&userToken).Error; err != nil {
			return err
		}

		if ac.App.Env == "development" {
			user.Note = "For Dev Only: use this path '/account/confirm-by-email?email=" + user.Email + "&token=" + userToken.Token.String() + "' to activate account"
		}

		return nil
	}); err != nil {
		ed := sh.Event{
			Feature: "register",
			Action:  "create-data",
			Source:  source,
			Status:  "failed",
			ECode:   "data-create-fail",
			EDetail: err.Error(),
		}
		return nil, sh.SetError(ed, nil)
	}

	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "created",
		},
		Data: user,
	}, nil
}

// confirm account using email uses the link for activation sent by email
func ConfirmByEmail(input m.ResendConfirmationEmailDto) (*ds.Data, error) {
	// check user token
	var userToken mut.UserToken
	if localECode, myEDetail := sut.CheckByEmail(&userToken, input.Email, input.Token, mut.UTConfirmByEmail); localECode != "" {
		return nil, es.XErrors{"email-confirmation": es.XError{Code: localECode, Message: l.I.Msg(localECode) + ", " + myEDetail}}
	}

	// check user
	var user m.User
	if err := checkUserStatus(&user, input.Email, m.USNew); err != nil {
		return nil, es.XErrors{"email-confirmation": es.XError{Code: "emailConfirm-fail", Message: l.I.Msg("emailConfirm-fail") + ", " + err.Error()}}
	}

	// go
	if err := dg.I.Transaction(func(tx *gorm.DB) error {
		// update user
		UserStatus := m.USActive
		user.Status = &UserStatus
		result := tx.Save(&user)
		if result.RowsAffected == 0 {
			return fmt.Errorf(l.I.Msg("save-fail") + ": user")
		}

		// delete token since it has been used
		result = tx.Unscoped().Delete(&userToken)
		if result.RowsAffected == 0 {
			return fmt.Errorf(l.I.Msg("save-fail") + ": token")
		}

		// create profile upon activation because the ID will be added in the JWT
		if _, err := pf.Create(mp.CreateDto{User_Id: user.Id}, tx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		ed := sh.Event{
			Feature: "confirm-by-email",
			Action:  "process-data",
			Source:  source,
			Status:  "failed",
			ECode:   "save-fail",
			EDetail: err.Error(),
		}
		return nil, sh.SetError(ed, nil)
	}

	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "confirmed",
		},
		Data: user,
	}, nil
}

func ResendConfirmationEmail(input m.ResendEmailConfirmDto) (*ds.Data, error) {
	// check user
	var user m.User
	if err := checkUserStatus(&user, input.Email, m.USNew); err != nil {
		return nil, es.XErrors{"user": es.XError{Code: "emailConfirm-resend-fail", Message: l.I.Msg("emailConfirm-resend-fail") + ", " + err.Error()}}
	}

	// check user token
	const tokenType = mut.UTConfirmByEmail
	var userToken mut.UserToken
	if err := sut.Request(&userToken, input.Email, tokenType, 60*24*3, 5); err != nil {
		return nil, es.XErrors{"user": es.XError{Code: "emailConfirm-resend-fail", Message: l.I.Msg("emailConfirm-resend-fail") + ", " + err.Error()}}
	}

	// next go
	ed := sh.Event{
		Feature: "confirm-by-email",
		Action:  "create-data",
		Source:  source,
		Status:  "failed",
		ECode:   "generate-fail",
	}
	if err := sut.Request(&userToken, input.Email, tokenType, 3*24*60, 5); err != nil {
		ed.EDetail = "token"
		return nil, sh.SetError(ed, userToken)
	} else if err := dg.I.Save(&userToken).Error; err != nil {
		ed.EDetail = "confirmation email"
		return nil, sh.SetError(ed, userToken)
	}

	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "request-is-made",
		},
		Data: user,
	}, nil
}

func ChangePassword(id int, input m.ChangePassDto) (*ds.Data, error) {
	var user m.User
	result := dg.I.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return nil, sh.SetError(sh.Event{
			Feature: "change-password",
			Action:  "fetch-data",
			Source:  source,
			Status:  "failed",
			ECode:   "fetch-data-fail",
		}, user)
	} else if !p.Check(input.OldPassword, *user.Password) {
		return nil, es.XErrors{"oldPassword": es.XError{Code: "invalid", Message: l.I.Msg("invalid")}}
	}

	password, err := p.Hash(input.NewPassword)
	if err != nil {
		return nil, es.XErrors{"password": es.XError{Code: "data-generate-fail", Message: l.I.Msg("data-generate-fail")}}
	}

	user.Password = &password
	if result := dg.I.Save(&user); result.Error != nil {
		return nil, sh.SetError(sh.Event{
			Feature: "change-password",
			Action:  "save-data",
			Source:  source,
			Status:  "failed",
			ECode:   "save-data-fail",
		}, user)
	}

	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "updated",
		},
		Data: user,
	}, nil
}

func RequestResetPass(input m.RequestResetPassDto) (*ds.Data, error) {
	// check user
	var user m.User
	if err := checkUserStatus(&user, input.Email, m.USActive); err != nil {
		return nil, es.XErrors{"email": es.XError{Code: "data-process-fail", Message: l.I.Msg("data-process-fail") + ", " + err.Error()}}
	}

	// chek user token
	var userToken mut.UserToken
	const tokenType = mut.UTResetPass
	if err := sut.Request(&userToken, input.Email, tokenType, 30, 5); err != nil {
		return nil, es.XErrors{"email": es.XError{Code: "data-process-fail", Message: l.I.Msg("data-process-fail") + ", " + err.Error()}}
	} else if err := dg.I.Save(&userToken).Error; err != nil {
		ed := sh.Event{
			Feature: "reset-password",
			Action:  "request",
			Source:  source,
			Status:  "failed",
			ECode:   "save-fail",
			EDetail: err.Error(),
		}
		return nil, sh.SetError(ed, userToken)
	}

	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "request-is-made",
		},
		Data: user,
	}, nil
}

func CheckResetPass(input m.CheckResetPassDto) (any, error) {
	var userToken mut.UserToken
	if errCode, _ := sut.CheckByEmail(&userToken, input.Email, input.Token, mut.UTResetPass); errCode != "" {
		return nil, es.XErrors{"email": es.XError{Code: errCode, Message: l.I.Msg(errCode)}}
	}

	return &ds.Data{
		Meta: ds.IS{
			"source":    "usertoken",
			"structure": "single-data",
			"status":    "fetched",
		},
		Data: userToken,
	}, nil
}

func ResetPass(input1 m.CheckResetPassDto, input2 m.ResetPassDto) (*ds.Data, error) {
	var userToken mut.UserToken
	if errCode, _ := sut.CheckByEmail(&userToken, input1.Email, input1.Token, mut.UTResetPass); errCode != "" {
		return nil, es.XErrors{"email": es.XError{Code: errCode, Message: l.I.Msg(errCode)}}
	}

	// TODO: PINDAH KE VALIDATOR
	if input2.NewPassword != input2.RePassword {
		return nil, es.XErrors{"rePassword": es.XError{Code: "equalToField", Message: l.I.Msg("equalToField") + " new password"}}
	}

	ed := sh.Event{
		Feature: "reset-password",
		Action:  "request",
		Source:  source,
		Status:  "failed",
		ECode:   "save-fail",
	}

	var user m.User
	result := dg.I.Where(&m.User{Email: input1.Email}).First(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		ed.ECode = "data-fetch-fail"
		ed.EDetail = result.Error.Error()
		return nil, sh.SetError(ed, user)
	}

	password, err := p.Hash(input2.NewPassword)
	if err != nil {
		ed.ECode = "data-process-fail"
		ed.EDetail = err.Error()
		return nil, sh.SetError(ed, user)
	} else {
		user.Password = &password
	}

	if result := dg.I.Save(&user); result.Error != nil {
		ed.ECode = "data-update-fail"
		ed.EDetail = result.Error.Error()
		return nil, sh.SetError(ed, user)
	}

	dg.I.Unscoped().Delete(&userToken)
	sterilizeUser(&user)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "user",
			"structure": "single-data",
			"status":    "request-is-made",
		},
		Data: user,
	}, nil
}
