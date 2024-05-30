package dbi

import (
	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
)

type Interface interface {
	Init(*t.DbConf)
}
