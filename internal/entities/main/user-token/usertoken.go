package usertoken

import (
	"time"

	"github.com/google/uuid"
	"github.com/munaja/exam-deals-yc-w22/internal/entities/helper/base"

	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
)

type Type string

type UserToken struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	User_Email   string     `json:"user_email" gorm:"index:unique"`
	User         user.User  `gorm:"foreignKey:User_Email;references:Email"`
	Type         Type       `json:"tpye" gorm:"size:20"`
	Token        uuid.UUID  `json:"token,omitempty" gorm:"size:255"`
	ExpiredAt    *time.Time `json:"expiredAt"`
	AttemptCount byte       `json:"attemptCount"`
}

type CreateDto struct {
	User_Email string `json:"user_id" validate:"required;email"`
}

type UpdateDto struct {
	User_Email string `json:"user_email" validate:"required;email"`
	Type       string `json:"type" validate:"required"`
	Token      int16  `json:"position" validate:"required"`
}

const UTConfirmByEmail Type = "ConfirmByEmail"
const UTResetPass Type = "ResetPass"

var usList map[Type]string = map[Type]string{
	"ConfirmByEmail": "Confirmation By Email",
	"ResetPass":      "Reset Password",
}

func GetTypeText(code Type) string {
	status := usList[code]
	return status
}
