package langown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"maps"
	"os"

	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
	l "github.com/munaja/exam-deals-yc-w22/pkg/language"
)

type langOwn struct{}

var O langOwn = langOwn{}
var I l.LangData = l.LangData{}

func (o *langOwn) Init(c *t.LangConf) {
	I.Active = c.Active
	I.List = map[string]l.LangItem{"en": l.DefaultList}

	jsonFile, err := os.Open(fmt.Sprintf("%v/%v/%v", c.Path, c.Active, c.FileName))
	if err != nil {
		log.Fatal("failed to load source file. " + err.Error())
	}
	defer jsonFile.Close()

	var myLI l.LangItem = l.LangItem{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myLI)
	maps.Copy(I.List["en"], myLI)
	log.Println("Instantiation for language using lang-own, status: DONE!!")
}
