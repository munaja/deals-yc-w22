package apicore

import (
	"flag"
	"log"
	"os"

	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
	"gopkg.in/yaml.v3"
)

type extCall func()

type app struct {
	CodeName        string
	FullName        string `yaml:"fullName"`
	Env             string
	Version         string
	LangConf        *t.LangConf        `yaml:"langConf"`
	DbConf          *t.DbConf          `yaml:"dbConf"`
	MsConf          *t.MsConf          `yaml:"msConf"`
	HttpConf        *t.HttpConf        `yaml:"httpConf"`
	RateLimiterConf *t.RateLimiterConf `yaml:"rateLimiterConf"`
	extCalls        []extCall
}

func (a *app) initConfig() {
	CfgFile = "./config.yml"
	flag.StringVar(&CfgFile, "config-file", "./config.yml", "Configuration path (default=./config.yaml)")
	flag.Parse()

	yamlFile, err := os.ReadFile(CfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Print("config is loaded successfully")
}

func (a *app) RegisterExtrCall(e extCall) {
	a.extCalls = append(a.extCalls, e)
}

func (a *app) initExtCall() {
	for _, init := range a.extCalls {
		init()
	}
}
