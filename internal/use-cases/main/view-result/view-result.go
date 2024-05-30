package viewresult

import (
	"errors"
	"strconv"
	"strings"
	"time"

	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	ms "github.com/munaja/exam-deals-yc-w22/pkg/api-core/ms-redis"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	"gorm.io/gorm"

	mp "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/view-result"
	sv "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/view"
	uh "github.com/munaja/exam-deals-yc-w22/pkg/use-case-helper"
)

const source = "view-result"

var limit = 10 // available for modification, by app settings for example

// Set result of the profile that is currently being viewed
func Create(input m.CreateDto) (*ds.Data, error) {
	event := uh.Event{
		Source:  source,
		Feature: "create",
		Status:  "fail",
	}

	profileData := &mp.Profile{}
	// get the data from reids
	listVal, _ := ms.I.Get(sv.GenerateListRedisKey(input.Viewer_Profile_Id)).Result()
	list := []int{}

	// if not availbale, cancel because there is target to be validated
	if listVal == "" {
		event.Action = "getting list"
		event.ECode = "data-fetch-fail"
		event.EDetail = "failed to get list"
		return nil, uh.SetError(event, nil)
	}

	// if there is list simply get it by converting the value back to its original form
	listItems := strings.Split(listVal, " ")
	for i := range listItems {
		n, _ := strconv.Atoi(listItems[i])
		list = append(list, n)
	}

	// get the last idx
	idxKey := sv.GenerateIdxRedisKey(input.Viewer_Profile_Id)
	idxVal, _ := ms.I.Get(idxKey).Int()

	// reached end of the list
	if idxVal > len(list)-1 {
		event.Action = "getting detail"
		event.ECode = "data-fetch-fail"
		event.EDetail = "failed to get detail, out of bound index"
		return nil, uh.SetError(event, nil)
	}
	if list[idxVal] != input.Target_Profile_Id {
		event.Action = "getting detail"
		event.ECode = "data-state-mismatch"
		event.EDetail = "data being viewed is not the same with data being scored"
		return nil, uh.SetError(event, nil)
	}

	if err := dg.I.Where("Id = ?", list[idxVal]).Find(&profileData).Error; err != nil {
		event.Action = "getting profile detail"
		event.ECode = "data-fetch-fail"
		event.EDetail = "data being viewed can not be found"
		return nil, uh.SetError(event, nil)
	}

	data := m.ViewResult{}
	if err := dg.I.Where("Viewer_Profile_Id = ? AND Target_Profile_Id = ?", input.Viewer_Profile_Id, input.Target_Profile_Id).
		Find(&data).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		event.Action = "getting result detail"
		event.ECode = "data-fetch-fail"
		event.EDetail = "data being viewed can not be found"
		return nil, uh.SetError(event, nil)
	}

	status := "created"
	if err := dg.I.Transaction(func(tx *gorm.DB) error {
		if data.Id < 1 {
			data.Viewer_Profile_Id = input.Viewer_Profile_Id
			data.Target_Profile_Id = input.Target_Profile_Id
			data.Result = input.Result
			if err := tx.Create(&data).Error; err != nil {
				return err
			}
		} else {
			data.Result = input.Result
			if err := tx.Save(&data).Error; err != nil {
				return err
			}
			status = "updated"
		}

		currentTime := time.Now()
		hourLeft := 23 - currentTime.Hour()
		minLeft := 59 - currentTime.Minute()
		secLeft := 60 - currentTime.Second()
		expTime := currentTime.Add(time.Duration(hourLeft)*time.Hour + time.Duration(minLeft)*time.Minute + time.Duration(secLeft)*time.Second)
		expDuration := expTime.Sub(currentTime)
		ms.I.Set(idxKey, idxVal+1, expDuration)

		return nil
	}); err != nil {
		event.Action = "saving data"
		event.ECode = "data-save-fail"
		event.EDetail = "can't save data"
		return nil, uh.SetError(event, nil)
	}

	return &ds.Data{
		Meta: ds.IS{
			"source":    "view",
			"structure": "single-data",
			"status":    status,
		},
		Data: data,
	}, nil

}
