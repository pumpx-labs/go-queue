package main

import (
	"fmt"

	"github.com/pumpx-labs/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: "localhost:11300",
				Tube:     "tube",
			},
			{
				Endpoint: "localhost:11300",
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: "localhost:6379",
			Type: redis.NodeType,
		},
	})
	consumer.Consume(func(body []byte) {
		fmt.Println(string(body))
	})
}
