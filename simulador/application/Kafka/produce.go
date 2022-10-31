package kafka
import (
	route2 "github.com/patrickcaloriocarvalho/GeoCarMaps/application/route"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patrickcaloriocarvalho/GeoCarMaps/infra/kafka"
	"encoding/json"
	"log"
	"os"
	"time"
)
//{"routeId":"1","clientId":"1"}
//{"routeId":"2","clientId":"2"}
//{"routeId":"3","clientId":"3"}
func Produce(msg *ckafka.Message){
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value,&route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _,p := range positions{
		kafka.Publish(p,os.Getenv("KafkaProduceTopic"),producer)
		time.Sleep(time.Millisecond * 500)
	}
}