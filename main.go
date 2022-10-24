package main

import (
	// route "github.com/jademirf/imersao-fs-fc2/simulator/app/route"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jademirf/imersao-fs-fc2/simulator/infra/kafka"
	kafka2 "github.com/jademirf/imersao-fs-fc2/simulator/kafka"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file" + err.Error())
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}


	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Hi", "readtest", producer)

	// for {
	// 	_ = 1
	// }

	// route := route.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[0])
}