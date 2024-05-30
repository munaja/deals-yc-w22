package apicore

import (
	"net/http"

	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
	"github.com/munaja/exam-deals-yc-w22/pkg/api-core/dbi"
	"github.com/munaja/exam-deals-yc-w22/pkg/api-core/httpi"
	"github.com/munaja/exam-deals-yc-w22/pkg/api-core/langi"
	"github.com/munaja/exam-deals-yc-w22/pkg/api-core/msi"
)

var App *app
var CfgFile string

// init
func init() {
	App = &app{
		LangConf: &t.LangConf{},
		DbConf:   &t.DbConf{},
		MsConf:   &t.MsConf{},
		HttpConf: &t.HttpConf{},
	}
	App.initConfig()
}

// app start the App
func Run(h http.Handler, m ...any) {
	for i := range m {
		if myModule, ok := m[i].(langi.Interface); ok {
			myModule.Init(App.LangConf)
		} else if myModule, ok := m[i].(dbi.Interface); ok {
			myModule.Init(App.DbConf)
		} else if myModule, ok := m[i].(msi.Interface); ok {
			myModule.Init(App.MsConf)
		} else if myModule, ok := m[i].(httpi.Interface); ok {
			myModule.Init(App.HttpConf, &h)
		}
	}
}
