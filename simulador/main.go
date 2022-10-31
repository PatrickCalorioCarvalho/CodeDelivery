package main
import (
		"fmt"
		"github.com/patrickcaloriocarvalho/GeoCarMaps/infra/kafka"
		kafka2 "github.com/patrickcaloriocarvalho/GeoCarMaps/application/kafka"
		ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
		"github.com/joho/godotenv"
		"log"
)
func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao ler o arquivo .env")
	}
}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan{
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg) 
	}
}