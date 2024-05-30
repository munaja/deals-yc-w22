package config

import (
	"log"
	"os"

	a "github.com/munaja/exam-deals-yc-w22/pkg/api-core"
	"gopkg.in/yaml.v3"
)

type authConf struct {
	AtSecretKey string `yaml:"atSecretKey"`
	RtSecretKey string `yaml:"rTSecretKey"`
}

var AuthConf authConf = authConf{}

func SetConfig() {
	yamlFile, err := os.ReadFile(a.CfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = yaml.Unmarshal(yamlFile, AuthConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
