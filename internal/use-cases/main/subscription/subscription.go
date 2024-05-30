package subscription

import (
	"errors"
	"time"

	sc "github.com/jinzhu/copier"
	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	ms "github.com/munaja/exam-deals-yc-w22/pkg/api-core/ms-redis"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	lh "github.com/munaja/exam-deals-yc-w22/pkg/language-helper"
	"gorm.io/gorm"

	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	sh "github.com/munaja/exam-deals-yc-w22/pkg/use-case-helper"

	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/subscription"
)

const source = "subscription"

// create data which is stored in both database and memory-cache
func Create(input m.CreateDto) (*ds.Data, error) {
	event := sh.Event{
		Feature: "create",
		Source:  source,
		Status:  "failed",
	}

	if m.GetTypeText(input.Type) == "" {
		return nil, es.XErrors{"type": es.XError{Code: "invalid", Message: lh.ErrorMsgGen("invalid")}}
	}

	data := &m.Subscription{}
	if err := dg.I.Where("User_Id = ? AND Type = ?", input.User_Id, input.Type).
		Find(&data).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		event.ECode = "data-fetch-fail"
		event.Action = "fetching currect data"
		event.EDetail = err.Error()
		return nil, sh.SetError(event, nil)
	}

	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		event.ECode = "data-copy-fail"
		event.Action = "copy input to data"
		event.EDetail = err.Error()
		return nil, sh.SetError(event, data)
	}

	currentTime := time.Now()
	nextTime := time.Now()
	if data.Id < 1 {
		data.ExpiredDate = nextTime.AddDate(0, input.Months, 0)
		if res := dg.I.Create(&data); res.Error != nil {
			event.ECode = "data-create-fail"
			event.Action = "create data"
			event.EDetail = res.Error.Error()
		}
	} else {
		if currentTime.Before(data.ExpiredDate) {
			data.ExpiredDate = data.ExpiredDate.AddDate(0, input.Months, 0)
		} else {
			data.ExpiredDate = nextTime.AddDate(0, input.Months, 0)
		}
		if err := dg.I.Save(&data).Error; err != nil {
			event.ECode = "data-update-fail"
			event.Action = "create data"
			event.EDetail = err.Error()
			return nil, sh.SetError(event, data)
		}
	}

	// no need to check error due the usage of double check for subscripttion validation
	// redis is the secondary source, but it will be checked first and if it fails in the checking
	// system will check the database directly
	ms.I.Set(generateRedisKey(input.User_Id, string(input.Type)), true, nextTime.Sub(currentTime))

	return &ds.Data{
		Meta: ds.IS{
			"source":    "subscription",
			"structure": "single-data",
			"status":    "created",
		},
		Data: data,
	}, nil
}

// check subscription status both in redis and database
func Check(User_Id int, theType m.Type) bool {
	// redis first
	msData := ms.I.Get(generateRedisKey(User_Id, string(theType)))
	msDataVal := msData.Val()
	if msDataVal == "" {
		// database second
		data := &m.Subscription{}
		if err := dg.I.Where("User_Id = ? AND Type = ? AND ExpiredDate < CURRENT_DATE()?", User_Id, theType).
			Find(&data).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			sh.SetError(sh.Event{
				Feature: "check",
				Source:  source,
				Status:  "failed",
				ECode:   "data-fetch-fail",
				Action:  "fetching currect data",
				EDetail: err.Error(),
			}, nil)
		}
		if data.Id < 1 {
			return false
		}
		return true
	}
	return true
}

//
