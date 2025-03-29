package config

import "github.com/pumpx-labs/go-queue/rabbitmq"

type Config struct {
	ListenerConf rabbitmq.RabbitListenerConf
}
