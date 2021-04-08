package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnatasr/fake99-simulator/application/route"
	"github.com/johnatasr/fake99-simulator/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		fmt.Println("Publish created:  ", p)
		time.Sleep(time.Millisecond * 500)
	}
}
