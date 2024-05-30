package httpi

import (
	"net/http"

	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
)

type Interface interface {
	Init(*t.HttpConf, *http.Handler)
}
