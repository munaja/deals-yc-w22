package view

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	ms "github.com/munaja/exam-deals-yc-w22/pkg/api-core/ms-redis"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"

	mp "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	mss "github.com/munaja/exam-deals-yc-w22/internal/entities/main/subscription"
	m "github.com/munaja/exam-deals-yc-w22/internal/entities/main/view"
	ss "github.com/munaja/exam-deals-yc-w22/internal/use-cases/main/subscription"
)

const source = "view"

var limit = 10 // available for modification, by app settings for example

// Get detail of the profile that is currently being viewed
// any activity view and like will be stored in redis for its performance
// but the limitation mark will be saved in the database
func GetDetail(input m.GetDetailDto) (*ds.Data, error) {
	data := &mp.Profile{}
	// get the data from reids
	listVal, _ := ms.I.Get(GenerateListRedisKey(input.Viewer_Profile_Id)).Result()
	list := []int{}

	// if not availbale
	if listVal == "" {
		list := GenerateList(m.GenerateListDto{Viewer_Profile_Id: input.Viewer_Profile_Id})
		if len(list) > 0 {
			dg.I.Where("Id = ?", list[0]).Find(&data)
		} else {
			data = nil
		}
	} else {
		// if there is list simply get it by converting the value back to its original form
		listItems := strings.Split(listVal, " ")
		for i := range listItems {
			n, _ := strconv.Atoi(listItems[i])
			list = append(list, n)
		}

		// get the last idx
		idxKey := GenerateIdxRedisKey(input.Viewer_Profile_Id)
		idxVal, _ := ms.I.Get(idxKey).Int()

		// reached end of the list
		if idxVal > len(list)-1 {
			// suscriber can always generate new list
			if ss.Check(input.Viewer_User_Id, mss.TPremium) {
				list := GenerateList(m.GenerateListDto{Viewer_Profile_Id: input.Viewer_Profile_Id})
				if len(list) > 0 {
					dg.I.Where("Id = ?", list[0]).Find(&data)
				} else {
					data = nil
				}
			} else {
				// reached limit for non subscriber
				data = nil
			}
		} else {
			dg.I.Where("Id = ?", list[idxVal]).Find(&data)
		}
	}

	return &ds.Data{
		Meta: ds.IS{
			"source":    "view",
			"structure": "single-data",
			"status":    "fetched",
		},
		Data: data,
	}, nil
}

// generate list of candidates for view
func GenerateList(input m.GenerateListDto) []int {
	// get profile info
	user := struct{ Gender int }{Gender: 0}
	dg.I.Select("Gender").Table("profile").Where("Id = ?", input.Viewer_Profile_Id).Scan(&user)

	// get list profile
	data := []struct{ Id int }{}
	dg.I.Select("`profile`.`Id`").
		Table("`Profile`").
		Joins("LEFT JOIN `viewresult` ON `profile`.`Id`=`viewresult`.`Target_Profile_Id` AND `viewresult`.`Viewer_Profile_Id` = "+strconv.Itoa(input.Viewer_Profile_Id)).
		Where("`profile`.`id` != ? AND `profile`.`Gender` != ? AND (ISNULL(`viewresult`.`UpdatedAt`) OR `viewresult`.`UpdatedAt`<CURRENT_DATE())", input.Viewer_Profile_Id, user.Gender).
		Order("`viewresult`.`UpdatedAt`").
		Limit(limit).
		Scan(&data)

	// get id only
	result := []int{}
	for i := range data {
		result = append(result, data[i].Id)
	}

	// set data in redis exp time just in case
	currentTime := time.Now()
	hourLeft := 23 - currentTime.Hour()
	minLeft := 59 - currentTime.Minute()
	secLeft := 60 - currentTime.Second()
	expTime := currentTime.Add(time.Duration(hourLeft)*time.Hour + time.Duration(minLeft)*time.Minute + time.Duration(secLeft)*time.Second)
	expDuration := expTime.Sub(currentTime)
	ms.I.Set(GenerateListRedisKey(input.Viewer_Profile_Id), strings.Trim(fmt.Sprint(result), "[]"), expDuration)
	ms.I.Set(GenerateIdxRedisKey(input.Viewer_Profile_Id), 0, expDuration)

	return result
}

func GenerateListRedisKey(User_Id int) string {
	return "viewList_" + strconv.Itoa(User_Id)
}

func GenerateIdxRedisKey(User_Id int) string {
	return "viewIdx_" + strconv.Itoa(User_Id)
}
