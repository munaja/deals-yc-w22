package languagehelper

import (
	"fmt"

	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
)

func ErrorMsgGen(errCode string, errDetail ...string) string {
	errMsg := ""
	if len(errDetail) == 0 || errDetail[0] == "" {
		errMsg = l.I.Msg(errCode)
	} else if len(errDetail) == 1 && errDetail[0] != "" { // manual
		errMsg = fmt.Sprintf(l.I.Msg(errCode), errDetail[0])
	} else if len(errDetail) == 2 && errDetail[0] != "" && errDetail[1] != "" { // manual
		errMsg = fmt.Sprintf(l.I.Msg(errCode), errDetail[0], errDetail[1])
	}
	return errMsg
}

func ErrorBundler(errCode string, errDetail ...string) es.XError {
	if len(errDetail) == 0 {
		return es.XError{Code: errCode, Message: ErrorMsgGen(errCode)}
	} else {
		return es.XError{Code: errCode, Message: ErrorMsgGen(errCode, errDetail[0])}
	}
}
