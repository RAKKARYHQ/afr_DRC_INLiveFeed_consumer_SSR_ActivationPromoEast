package afr_kafka

//reference: https://github.com/segmentio/kafka-go
//reference: https://pkg.go.dev/github.com/segmentio/kafka-go#section-readme
import (
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	//Writer *kafka.Writer
	Reader_SSR_4   *kafka.Reader
	Reader_SSR_26 *kafka.Reader

}

func NewKafkaClient() (conn *KafkaClient) {
	conn = new(KafkaClient)
	BrokerUrls := Configuration.KafkaBrokerUrls
	ClientId := Configuration.KafkaClientId
	conn.createKafkaReader(BrokerUrls, ClientId)
	return
}

func prepareKafkaReaderConfig(BrokerUrls, ClientId, Topic string) (config kafka.ReaderConfig) {
	config = kafka.ReaderConfig{
		Brokers: strings.Split(BrokerUrls, ","),
		GroupID: ClientId,
		Topic:   Topic,
		//MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: 1,
		StartOffset:     kafka.LastOffset, //kafka.FirstOffset,
		CommitInterval:  time.Second,      // flushes commits to Kafka every second
	}
	return config
}

func (c *KafkaClient) createKafkaReader(BrokerUrls, ClientId string) (err error) {
	log.Println("creating kafka readers")
	// config := kafka.ReaderConfig{
	// 	Brokers: strings.Split(BrokerUrls, ","),
	// 	GroupID: ClientId,
	// 	Topic:   Topic,
	// 	//MinBytes:        10e3,            // 10KB
	// 	MaxBytes:        10e6,            // 10MB
	// 	MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
	// 	ReadLagInterval: 1,
	// 	StartOffset:     kafka.FirstOffset,
	// 	CommitInterval:  time.Second, // flushes commits to Kafka every second
	// }
	// c.Reader = kafka.NewReader(config)
	// log.Println("kafka reader created successfuly ")

	config_SSR_4 := prepareKafkaReaderConfig(BrokerUrls, ClientId, "SSR_4")
	c.Reader_SSR_4 = kafka.NewReader(config_SSR_4)

	config_SSR_26 := prepareKafkaReaderConfig(BrokerUrls, ClientId, "SSR_26")
	c.Reader_SSR_26 = kafka.NewReader(config_SSR_26)

	log.Println("kafka readers created successfuly ")
	return
}
