//Executar no Kafka para teste : kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtes
package main
import (
		"fmt"
		"github.com/patrickcaloriocarvalho/GeoCarMaps/infra/kafka"
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
	}
}