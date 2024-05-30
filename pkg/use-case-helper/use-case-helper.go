package usecasehelper

import (
	"encoding/json"
	"log"

	l "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
)

// To standardize the error logging format FOR database related processes
type Event struct {
	Feature string
	Action  string
	Source  string
	Status  string
	ECode   string
	EDetail string
}

func SetError(d Event, data any) es.XError {
	dataString, _ := json.Marshal(data)
	msg := l.I.Msg(d.ECode)
	log.Printf(
		"code: %v, message: %v, feature: %v, source: %v, action: %v, status: %v, detail: %v, data: %v",
		d.ECode, msg, d.Feature, d.Source, d.Action, d.Status, d.EDetail, string(dataString),
	)
	return es.XError{Code: d.ECode, Message: l.I.Msg(d.ECode)}
}
