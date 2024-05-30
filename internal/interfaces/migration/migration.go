package migration

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/profile"
	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/subscription"
	"github.com/munaja/exam-deals-yc-w22/internal/entities/main/user"
	usertoken "github.com/munaja/exam-deals-yc-w22/internal/entities/main/user-token"
	viewresult "github.com/munaja/exam-deals-yc-w22/internal/entities/main/view-result"
)

type DbConf struct {
	Dsn     string
	Dialect string
}

// Migrate all tables at once, one time only for exam purpose
func Migrate() {
	// use default config file location or use flat
	cfgFile := "./config.yml"
	flag.StringVar(&cfgFile, "config-file", "./config.yml", "Configuration path (default=./config.yaml)")
	flag.Parse()

	// read the config file
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// parse into config struct
	var dbConf DbConf
	err = yaml.Unmarshal(yamlFile, &dbConf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("config is loaded successfully")

	// create database connection
	db, err := gorm.Open(mysql.Open(dbConf.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("database-connection is established successfully")

	// migrate all the tables
	modelList := []any{
		&profile.Profile{},
		&subscription.Subscription{},
		&subscription.SubscriptionLog{},
		&user.User{},
		&usertoken.UserToken{},
		&viewresult.ViewResult{},
	}
	db.AutoMigrate(modelList...)
	log.Printf("migration is complete")
}
