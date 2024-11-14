package afr_kafka

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var chan_SSR = make(chan []byte, 500000)
var chan_SSR_controler = make(chan int, 50)

func (Uc *UserControl) SSR_ReadersStats() {
	log.Println("SSR readers stats process started")
	for range time.Tick(time.Second * 15) {
		stats := Uc.KafkaClient.Reader_SSR_4.Stats()
		KReaderStatus.With(prometheus.Labels{"Reader": "SSR_4", "StatType": "Offset"}).Set(float64(stats.Offset))
		KReaderStatus.With(prometheus.Labels{"Reader": "SSR_4", "StatType": "Lag"}).Set(float64(stats.Lag))

		stats = Uc.KafkaClient.Reader_SSR_26.Stats()
		KReaderStatus.With(prometheus.Labels{"Reader": "SSR_26", "StatType": "Offset"}).Set(float64(stats.Offset))
		KReaderStatus.With(prometheus.Labels{"Reader": "SSR_26", "StatType": "Lag"}).Set(float64(stats.Lag))

	}
}

func (Uc *UserControl) SSR_Process() {
	log.Println("SSR process started")
	for {
		select {
		case msg := <-chan_SSR:
			go Uc.Post_SSR_totarget(msg)
			chan_SSR_controler <- 1
		default:
			<-time.After(10 * time.Millisecond)
		}
	}
}

func (c *KafkaClient) RunKafkaReader_SSR_4() {
	log.Println("Starting kafka reader SSR_4")
	for {
		if IsTarget_Up == 1 {
			kafkaMsg, err := c.Reader_SSR_4.ReadMessage(context.Background())
			if err != nil {
				log.Println("error while receiving message from SSR_4: ", err.Error())
				time.Sleep(30 * time.Second)
				continue
			}
			chan_SSR <- kafkaMsg.Value
			TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "4", "Description": "Kafka read stream"}).Inc()
		} else {
			log.Println("SSR_4 target host is not reachable. Sleep 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}
}

func (c *KafkaClient) RunKafkaReader_SSR_26() {
	log.Println("Starting kafka reader SSR_26")
	for {
		if IsTarget_Up == 1 {
			kafkaMsg, err := c.Reader_SSR_26.ReadMessage(context.Background())
			if err != nil {
				log.Println("error while receiving message from SSR_26: ", err.Error())
				time.Sleep(30 * time.Second)
				continue
			}
			chan_SSR <- kafkaMsg.Value
			TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "26", "Description": "Kafka read stream"}).Inc()
		} else {
			log.Println("SSR_26 target host is not reachable. Sleep 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}
}

// func (c *KafkaClient) RunKafkaReader_SSR_4() {
// 	log.Println("Starting kafka reader SSR_4")
// 	for {
// 		if IsTarget_Up == 1 {
// 			kafkaMsg, err := c.Reader_SSR_4.ReadMessage(context.Background())
// 			if err != nil {
// 				log.Println("error while receiving message from SSR_4: ", err.Error())
// 				time.Sleep(30 * time.Second)
// 				continue
// 			}
// 			if ssr.SSR.Cos == "340" && ssr.SSR.Cos == "331" {
// 				chan_SSR <- kafkaMsg.Value
// 				TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "4", "Description": "Kafka read stream"}).Inc()			
// 			}
// 			} else {
// 			log.Println("SSR_4 target host is not reachable. Sleep 5 seconds...")
// 			time.Sleep(5 * time.Second)
// 		}
// 	}
// }

// func (c *KafkaClient) RunKafkaReader_SSR_26() {
// 	log.Println("Starting kafka reader SSR_26")
// 	for {
// 		if IsTarget_Up == 1 {
// 			kafkaMsg, err := c.Reader_SSR_26.ReadMessage(context.Background())
// 			if err != nil {
// 				log.Println("error while receiving message from SSR_26: ", err.Error())
// 				time.Sleep(30 * time.Second)
// 				continue
// 			}
// 			var ssr SSR_Message
// 			err = xml.Unmarshal(kafkaMsg.Value, &ssr)
// 			if err != nil {
// 				log.Println("error in PostSSR - Unmarshal body: ", err)
// 				return
// 			}
// 			if ssr.SSR.Cos == "340" && ssr.SSR.Cos == "331" {
// 				chan_SSR <- kafkaMsg.Value
// 				TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "26", "Description": "Kafka read stream"}).Inc()
		
// 			}
// 		} else {
// 			log.Println("SSR_26 target host is not reachable. Sleep 5 seconds...")
// 			time.Sleep(5 * time.Second)
// 		}
// 	}
// }


