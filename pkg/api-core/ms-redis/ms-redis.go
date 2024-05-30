package msredis

import (
	"log"

	"github.com/go-redis/redis"
	t "github.com/munaja/exam-deals-yc-w22/pkg/api-core/all-types"
)

type ms struct{}

var O ms            // instance
var I *redis.Client // instance

func (o *ms) Init(c *t.MsConf) {
	I = redis.NewClient(&redis.Options{
		Addr: c.Dsn,
	})
	_, err := I.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println("Instantiation for memory-storage using ms-redis, status: DONE!!")
}
