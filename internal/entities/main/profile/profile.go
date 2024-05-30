package profile

import (
	"time"

	gdt "gorm.io/datatypes"

	"github.com/munaja/exam-deals-yc-w22/internal/entities/helper/base"
	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
)

type Profile struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	User_Id          int       `json:"user_id"`
	User             user.User `json:"-" gorm:"foreignKey:User_Id"`
	Name             string    `json:"name" gorm:"size:100"`
	Birthdate        *gdt.Date `json:"birthDate"`
	Gender           *byte     `json:"gender"`
	Address          string    `json:"address" gorm:"size:100"`
	BuildingNumber   *string   `json:"buildingNumber" gorm:"size:5"`
	Regency_Id       int       `json:"regency_id"`
	Postalcode       *string   `json:"postalCode" gorm:"size:6"`
	PhoneNumber      *string   `json:"phoneNumber" gorm:"size:20"`
	WhatsappNumber   *string   `json:"whatsappNumber" gorm:"size:20"`
	ProfileImgSrc    *string   `json:"profileImg" gorm:"size:1024"`
	LastGenerateView time.Time `json:"lastGenerateView"`
}

// minimalis DTO since the creation is part of account activation
type CreateDto struct {
	User_Id int
}

type UpdateDto struct {
	Name           string    `json:"name" validate:"required;alphaSpace;maxLength=100"`
	Birthdate      *gdt.Date `json:"birthDate" validate:"required"`
	Gender         *byte     `json:"gender" validate:"required"`
	Address        string    `json:"address" validate:"required;maxLength=200"`
	BuildingNumber *string   `json:"buildingNumber" validate:"maxLength=5"`
	Regency_Id     int       `json:"regency_id" validate:"required"`
	Postalcode     *string   `json:"postalCode" validate:"maxLength=6"`
	PhoneNumber    *string   `json:"phoneNumber" validate:"required;maxLength=20"`
	WhatsappNumber *string   `json:"whatsappNumber" validate:"maxLength=20"`
}

type ProfilDto struct {
	// user
	User_Id        int       `json:"user_id"`
	User_Name      string    `json:"user_name"`
	User_CreatedAt time.Time `json:"user_createdAt"`
	// person
	Name           string   `json:"name"`
	Birthdate      gdt.Date `json:"birthDate"` // TODO: there should be option to show / hide this field
	Gender         *byte    `json:"gender"`
	Address        string   `json:"address"`        // TODO: there should be option to show / hide this field
	BuildingNumber *string  `json:"buildingNumber"` // TODO: there should be option to show / hide this field
	Regency_Id     int      `json:"regency_id"`
	PhoneNumber    *string  `json:"phoneNumber"`    // TODO: there should be option to show / hide this field
	WhatsappNumber *string  `json:"whatsappNumber"` // TODO: there should be option to show / hide this field
}

type ProfilOMDto struct {
	// user
	User_Id        int       `json:"user_id"`
	User_Name      string    `json:"user_name"`
	User_CreatedAt time.Time `json:"user_createdAt"`
	// person
	Name           string   `json:"name"`
	Birthdate      gdt.Date `json:"birthDate"`
	Gender         *byte    `json:"gender"`
	Address        string   `json:"address"`
	BuildingNumber *string  `json:"buildingNumber"`
	Regency_Id     int      `json:"regency_id"`
	PhoneNumber    *string  `json:"phoneNumber"`
	WhatsappNumber *string  `json:"whatsappNumber"`
}
