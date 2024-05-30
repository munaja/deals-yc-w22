package viewresult

import (
	"github.com/munaja/exam-deals-yc-w22/internal/entities/helper/base"
	p "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
)

type Result string

const RLike = "like"
const RPass = "pass"

type ViewResult struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	base.DateModel
	Viewer_Profile_Id int       `json:"viewer_profile_id"`
	Viewer_Profile    p.Profile `gorm:"foreignKey:Viewer_Profile_Id"`
	Target_Profile_Id int       `json:"target_profile_id"`
	Target_Profile    p.Profile `gorm:"foreignKey:Target_Profile_Id"`
	Result            Result    `json:"result" gorm:"size:10"`
}

type CreateDto struct {
	Viewer_Profile_Id int    `json:"-"`
	Target_Profile_Id int    `json:"target_profile_id" validate:"required;numeric"`
	Result            Result `json:"result" validate:"required;alpha"`
}

type GetListDto struct {
}

type ListDto struct {
}

type UpdateDto struct {
	Viewer_Profile_Id int    `json:"viewer_profile_id"`
	Target_Profile_Id int    `json:"target_profile_id"`
	Result            Result `json:"result" gorm:"size:10"`
}

var resultList map[Result]string = map[Result]string{
	RLike: "Like",
	RPass: "Pass",
}

func GetTypeText(code Result) string {
	status := resultList[code]
	return status
}
