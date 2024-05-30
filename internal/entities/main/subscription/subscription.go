package subscription

import (
	"time"

	"github.com/munaja/exam-deals-yc-w22/internal/entities/helper/base"
	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
)

type Type string

const TPremium Type = "premium"

type Subscription struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	User_Id          int       `json:"user_id"`
	User             user.User `json:"-" gorm:"foreignKey:User_Id;PRELOAD:false"`
	Type             Type      `json:"name" gorm:"size:100"`
	ExpiredDate      time.Time `json:"expiredDate"`
	PaymentMethod_Id int       `json:"paymentMethod_id"`
}

type SubscriptionLog struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	User_Id          int       `json:"user_id"`
	User             user.User `json:"user" gorm:"foreignKey:User_Id;PRELOAD:false"`
	Type             Type      `json:"name" gorm:"size:100"`
	ExpiredDate      time.Time `json:"expiredDate"`
	PaymentMethod_Id int       `json:"paymentMethod_id"`
}

type CreateDto struct {
	User_Id          int  `json:"-"`
	Type             Type `json:"type" validate:"required;alpha"`
	Months           int  `json:"months" validate:"required;numeric"`
	PaymentMethod_Id int  `json:"paymentMethod_id" validate:"required;numeric"`
}

var typeList map[Type]string = map[Type]string{
	TPremium: "Premium",
}

func GetTypeText(code Type) string {
	status := typeList[code]
	return status
}
