//Executar no Kafka para teste : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
package main
import (
		"github.com/patrickcaloriocarvalho/GeoCarMaps/infra/kafka"
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
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola","readtest",producer)
	for{
		_=1
	}
}