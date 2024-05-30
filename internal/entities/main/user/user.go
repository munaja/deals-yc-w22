package user

import (
	"fmt"
	"time"

	"github.com/munaja/exam-deals-yc-w22/internal/entities/helper/base"
)

type Status byte
type Position int16

type User struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	Name              string     `json:"name" gorm:"size:100;unique"`
	Email             string     `json:"email" gorm:"size:100;unique"`
	Password          *string    `json:"password,omitempty" gorm:"size:255"`
	Status            *Status    `json:"status,omitempty"`
	LoginAttemptCount int        `json:"-"`
	LastSuccessLogin  *time.Time `json:"lastSuccessLogin,omitempty"`
	LastAllowdLogin   *time.Time `json:"lastAllowdLogin,omitempty"`
	Note              string     `json:"note,omitempty" gorm:"-"`
}

type CreateDto struct {
	Name        string    `json:"name" validate:"required;maxLength=100"`
	Email       string    `json:"email" validate:"email;maxLength=100"`
	Password    string    `json:"password" validate:"required;minLength=8;maxLength=50"`
	Status      int16     `json:"status"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type UpdateDto struct {
	Email string `json:"email"`
}

type RegisterDto struct {
	Name       string `json:"name" validate:"required;maxLength=100"`
	Email      string `json:"email" validate:"required;email;maxLength=100"`
	Password   string `json:"password" validate:"required;minLength=8;maxLength=50"`
	RePassword string `json:"repassword" validate:"required;eqField=Password"`
}

type ResendConfirmationEmailDto struct {
	Email string `json:"email" validate:"required;email;maxLength=100"`
	Token string `json:"token" validate:"required"`
}

type ResendEmailConfirmDto struct {
	Email string `json:"email" validate:"required;email"`
}

type LoginDto struct {
	Name     string   `json:"name" validate:"required;maxLength=100"`
	Password string   `json:"password" validate:"required;maxLength=50"`
	LongTerm bool     `json:"longTerm"`
	Position Position `json:"-"`
}

type FilterDto struct {
	Name   *string `json:"name"`
	Type   *int16  `json:"type"`
	Email  *string `json:"email"`
	Status *int16  `json:"status"`
	// fixed fields
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type ChangePassDto struct {
	OldPassword string `json:"oldPassword" validate:"required;minLength=8"`
	NewPassword string `json:"newPassword" validate:"required;minLength=8"`
	RePassword  string `json:"rePassword" validate:"required;minLength=8;eqField=NewPassword"`
}

type RequestResetPassDto struct {
	Email string `json:"email" validate:"required;email"`
}

type CheckResetPassDto struct {
	Email string `json:"email" validate:"required;email"`
	Token string `json:"token" validate:"required"`
}

type ResetPassDto struct {
	NewPassword string `json:"newPassword" validate:"required;minLength=8"`
	RePassword  string `json:"rePassword" validate:"required;minLength=8"`
}

const (
	UPCustomer Position = 1
	UPOperator Position = 2
	UPAdmin    Position = 3
	UPOwner    Position = 4

	USNew       Status = 0
	USActive    Status = 1
	USBlocked   Status = 2
	USSuspended Status = 3
)

var usList map[Status]string = map[Status]string{
	0: "New",
	1: "Active",
	2: "Blocked",
	3: "Suspended",
}

func GetUSText(code Status) string {
	status, _ := usList[code]
	fmt.Println(code)
	fmt.Println(status)
	return status
}
