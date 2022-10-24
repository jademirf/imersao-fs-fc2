package kafka

import (
	"encoding/json"
	route "github.com/jademirf/imersao-fs-fc2/simulator/app/route"
	"github.com/jademirf/imersao-fs-fc2/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KAFKA_PRODUCE_TOPICS"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}