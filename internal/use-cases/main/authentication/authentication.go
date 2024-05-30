package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	dg "github.com/munaja/exam-deals-yc-w22/pkg/api-core/db-gorm-mysql"
	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ms "github.com/munaja/exam-deals-yc-w22/pkg/api-core/ms-redis"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	lh "github.com/munaja/exam-deals-yc-w22/pkg/language-helper"
	p "github.com/munaja/exam-deals-yc-w22/pkg/password"

	mp "github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	mu "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	c "github.com/munaja/exam-deals-yc-w22/internal/use-cases/helper/config"
)

//	type TokenDetails struct {
//		AccessToken  string
//		RefreshToken string
//		AccessUuid   string
//		RefreshUuid  string
//		AtExpires    int64
//		RtExpires    int64
//	}

// Generates token and store in redis at one place
// just return the error code
func GenToken(input mu.LoginDto) (*ds.Data, error) {
	// Get User
	var user mu.User
	if errCode := getAndCheck(&user, mu.User{Name: input.Name}); errCode != "" {
		return nil, es.XErrors{"authentication": es.XError{Code: errCode, Message: lh.ErrorMsgGen(errCode)}}
	}

	// manually get profile due to circular dependency
	var profile mp.Profile
	if result := dg.I.Where("User_Id  = ?", user.Id).Find(&profile); result.Error != nil {
		return nil, es.XError{Code: "data-fetch-fail", Message: "profile is corrupt"}
	}

	if user.LoginAttemptCount > 5 {
		if user.LastSuccessLogin != nil {
			now := time.Now()
			lastAllowdLogin := user.LastAllowdLogin
			if lastAllowdLogin.After(now.Add(-time.Hour * 1)) {
				return nil, es.XErrors{"authentication": es.XError{Code: "auth-login-tooMany", Message: lh.ErrorMsgGen("auth-login-tooMany")}}
			} else {
				tn := time.Now()
				user.LastAllowdLogin = &tn
				user.LoginAttemptCount = 0
				dg.I.Save(&user)
			}
		} else {
			tn := time.Now()
			user.LastAllowdLogin = &tn
			dg.I.Save(&user)
			return nil, es.XErrors{"authentication": es.XError{Code: "auth-login-tooMany", Message: lh.ErrorMsgGen("auth-login-tooMany")}}
		}
	}

	if !p.Check(input.Password, *user.Password) {
		user.LoginAttemptCount = user.LoginAttemptCount + 1
		dg.I.Save(&user)
		return nil, es.XErrors{"authentication": es.XError{Code: "auth-login-incorrect", Message: lh.ErrorMsgGen("auth-login-incorrect")}}
	} else if *user.Status == mu.USBlocked {
		return nil, es.XErrors{"authentication": es.XError{Code: "auth-login-blocked", Message: lh.ErrorMsgGen("auth-login-blocked")}}
	} else if *user.Status == mu.USNew {
		return nil, es.XErrors{"authentication": es.XError{Code: "auth-login-unverified", Message: lh.ErrorMsgGen("auth-login-unverified")}}
	}

	// Access token prep
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf(l.I.Msg("uuid-gen-fail"), err))
	}
	duration := time.Hour * 24
	if input.LongTerm {
		duration = time.Hour * 24 * 30
	}
	aUuid := id.String()
	atExpires := time.Now().Add(duration).Unix()
	atSecretKey := c.AuthConf.AtSecretKey

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["user_name"] = user.Name
	atClaims["user_email"] = user.Email
	atClaims["profile_id"] = profile.Id
	atClaims["exp"] = atExpires
	atClaims["uuid"] = aUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	ats, err := at.SignedString([]byte(atSecretKey))
	if err != nil {
		return nil, es.XErrors{"user": es.XError{Code: "token-sign-err", Message: lh.ErrorMsgGen("token-sign-err")}}
	}
	// Save to redis
	now := time.Now()
	atx := time.Unix(atExpires, 0) //converting Unix to UTC(to Time object)
	err = ms.I.Set(aUuid, strconv.Itoa(user.Id), atx.Sub(now)).Err()
	if err != nil {
		panic(fmt.Sprintf(l.I.Msg("redis-store-fail"), err.Error()))
	}

	tn := time.Now()
	user.LoginAttemptCount = 0
	user.LastSuccessLogin = &tn
	user.LastAllowdLogin = &tn
	dg.I.Save(&user)

	// Current data
	return &ds.Data{
		Meta: ds.IS{
			"source":    "authentication",
			"structure": "single-data",
			"status":    "verified",
		},
		Data: ds.II{
			"user_id":     strconv.Itoa(user.Id),
			"user_name":   user.Name,
			"user_email":  user.Email,
			"accessToken": ats,
		},
	}, nil
}

func RevokeToken(uuid string) {
	ms.I.Del(uuid)
}

func VerifyToken(r *http.Request, tokenType TokenType) (data *jwt.Token, errCode, errDetail string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, "auth-missingHeader", ""
	}
	authArr := strings.Split(auth, " ")
	if len(authArr) == 2 {
		auth = authArr[1]
	}

	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(l.I.Msg("token-sign-unexcpeted"), token.Header["alg"])
		}
		if tokenType == AccessToken {
			return []byte(c.AuthConf.AtSecretKey), nil
		} else {
			return []byte(c.AuthConf.RtSecretKey), nil
		}
	})
	if err != nil {
		return nil, "token-parse-fail", err.Error()
	}
	return token, "", ""
}

func ExtractToken(r *http.Request, tokenType TokenType) (data *AuthInfo, err error) {
	token, errCode, errDetail := VerifyToken(r, tokenType)
	if errCode != "" {
		return nil, es.XError{Code: errCode, Message: lh.ErrorMsgGen(errCode, errDetail)}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["uuid"].(string)
		if !ok {
			return nil, es.XError{Code: "token-invalid", Message: lh.ErrorMsgGen("token-invalid", "uuid not available")}
		}
		user_id, myErr := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
		if myErr != nil {
			return nil, es.XError{Code: "token-invalid", Message: lh.ErrorMsgGen("token-invalid", "user_id is not available")}
		}
		profile_id, myErr := strconv.Atoi(fmt.Sprintf("%v", claims["profile_id"]))
		if myErr != nil {
			return nil, es.XError{Code: "token-invalid", Message: lh.ErrorMsgGen("token-invalid", "profile_id is not available")}
		}
		accessUuidRedis := ms.I.Get(accessUuid)
		if accessUuidRedis.String() == "" {
			return nil, es.XError{Code: "token-unidentified", Message: lh.ErrorMsgGen("token-unidentified")}
		}
		data = &AuthInfo{
			Uuid:       accessUuid,
			User_Id:    int(user_id),
			User_Name:  fmt.Sprintf("%v", claims["user_name"]),
			User_Email: fmt.Sprintf("%.f", claims["user_email"]),
			Profile_Id: int(profile_id),
		}
		return
	}
	return nil, es.XError{Code: "token", Message: "token-invalid"}
}
