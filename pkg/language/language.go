package language

import (
	"golang.org/x/exp/maps"
)

// Configuration type that is used by the core
type LangItem map[string]string
type LangData struct {
	Active string
	List   map[string]LangItem
}

func New() LangData {
	return LangData{}
}

func (o *LangData) SetList(ln string, l LangItem) {
	o.List = map[string]LangItem{ln: l}
}

func (o *LangData) Add(name string) {
	_, ok := o.List[name]
	if !ok {
		o.List[name] = LangItem{}
	}
}

func (o *LangData) AddMsg(code string, message string, opt ...string) {
	lang := o.Active
	if len(opt) > 0 {
		o.Add(opt[1])
		lang = opt[1]
	}
	o.List[lang][code] = message
}

func (o *LangData) AddMsgList(list LangItem, opt ...string) {
	lang := o.Active
	if len(opt) > 0 {
		o.Add(opt[1])
		lang = opt[1]
	}

	maps.Copy(o.List[lang], list)
}

func (o *LangData) Msg(k string, opt ...string) string {
	lang := o.Active
	if len(opt) > 0 {
		o.Add(opt[1])
		lang = opt[1]
	}

	if msg, ok := o.List[lang][k]; !ok {
		return "** warning: usage of unlisted code:" + k + " **"
	} else {
		return msg
	}
}
