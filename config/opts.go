package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type (
	Opts struct {
		// logger
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		// kafka
		Kafka struct {
			// Kafka settings
			Host  string `long:"kafka.host"                 env:"KAFKA_HOST"                        description:"Kafka host, eg. kafka-0:9092" required:"true"`
			Topic string `long:"kafka.topic"                env:"KAFKA_TOPIC"                       description:"Kafka topic, eg. alertmanager" required:"true"`
		}

		// general options
		ServerBind string `long:"bind"     env:"SERVER_BIND"   description:"Server address"     default:":9097"`
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
