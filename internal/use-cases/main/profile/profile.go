package profile

import (
	"github.com/jinzhu/copier"
	sc "github.com/jinzhu/copier"
	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	"gorm.io/gorm"

	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	su "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/user"
	sh "github.com/munaja/exam-deals-yc-w22/pkg/use-case-helper"
)

const source = "profile"

// get profile detail with option which mostly is to distinguis between the owner and strangers
func GetDetail(user_name string, opt ...string) (*ds.Data, error) {
	// get user first
	user, err := su.GetDetailByName(user_name)
	if err != nil {
		return nil, es.XErrors{"data-notFound": es.XError{Code: "data-fetch-fail", Message: l.I.Msg("data-fetch-fail") + ", " + err.Error()}}
	}

	var person *m.Profile
	err = dg.I.Where("User_Id = ?", user.Id).First(&person).Error
	// special case, profile will still showing empty data if no data found
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, es.XErrors{"data-notFound": es.XError{Code: "process-fail", Message: l.I.Msg("data-fetch-fail")}}
	}

	var profile m.ProfilOMDto = m.ProfilOMDto{}
	profile.User_Id = user.Id
	profile.User_Name = user.Name
	profile.User_CreatedAt = user.CreatedAt
	if len(opt) > 0 && opt[0] == "of-mine" {
		if err = sc.CopyWithOption(&profile, person, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, es.XErrors{"data": es.XError{Code: "data-copy-fail"}}
		}
	} else {
		if err = sc.CopyWithOption(&profile, person, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, es.XErrors{"data": es.XError{Code: "data-copy-fail"}}
		}
	}
	return &ds.Data{
		Meta: ds.IS{
			"source":    "profile",
			"status":    "fetched",
			"structure": "single",
		},
		Data: profile,
	}, nil
}

// this should only be called upon account activation, no endpoints utilize this and no complete data is needed
func Create(input m.CreateDto, tx *gorm.DB) (*m.Profile, error) {
	profile := m.Profile{
		User_Id: input.User_Id,
	}
	if res := tx.Create(&profile); res.Error != nil {
		return nil, res.Error
	}
	return &profile, nil
}

// update data by user name
func UpdateByUserName(user_name string, input m.UpdateDto) (*ds.Data, error) {
	// prepare event
	event := sh.Event{
		Feature: source,
		Action:  "get-resource",
		Source:  source,
		Status:  "failed",
		ECode:   "data-fetch-fail",
	}

	// get user first, in case someone trying to force update unavailable data
	user, err := su.GetIdByName(user_name)
	if err != nil {
		return nil, es.XErrors{"data": es.XError{Code: "process-fail", Message: l.I.Msg("data-fetch-fail") + ", " + err.Error()}}
	}

	// old data
	var data *m.Profile = &m.Profile{}
	err = dg.I.Where("User_Id = ?", user["Id"]).First(&data).Error
	// special case, profile will still showing empty data if no data found
	if err != nil && err != gorm.ErrRecordNotFound {
		event.EDetail = err.Error()
		return nil, es.XErrors{"data": es.XError{Code: "process-fail", Message: l.I.Msg("data-fetch-fail") + ", " + err.Error()}}
	}

	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return nil, sh.SetError(event, data)
	}

	dg.I.Save(&data)
	return &ds.Data{
		Meta: ds.IS{
			"source":    "profile",
			"status":    "created",
			"structure": "single",
		},
		Data: data,
	}, nil
}
