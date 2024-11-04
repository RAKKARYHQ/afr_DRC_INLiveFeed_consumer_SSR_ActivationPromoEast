package afr_kafka

import (
	"log"
	"net"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var IsTarget_Up float64

type destinationHost struct {
	HostName string
	HostIP   string
	HostPort string
}

var destinationHots []destinationHost

func Init_DestinationHosts() {
	//Kafka Cluster port monitoring
	host := destinationHost{
		HostName: "kafka1",
		HostIP:   "kafka1",
		HostPort: "9092",
	}
	destinationHots = append(destinationHots, host)

	host = destinationHost{
		HostName: "kafka2",
		HostIP:   "kafka2",
		HostPort: "9092",
	}
	destinationHots = append(destinationHots, host)

	host = destinationHost{
		HostName: "kafka3",
		HostIP:   "kafka3",
		HostPort: "9092",
	}
	destinationHots = append(destinationHots, host)

	//Target Host
	host = destinationHost{
		HostName: "TargetHost",
		HostIP:   Configuration.TargetIP,
		HostPort: Configuration.TargetPort,
	}
	destinationHots = append(destinationHots, host)

}

func PortlinkInquiry() {
	log.Println("link inquiry started")
	Init_DestinationHosts()
	for range time.Tick(time.Second * 15) {
		for _, host := range destinationHots {
			hoststat := linkInquiry(host.HostIP, host.HostPort)
			PortStatus.With(prometheus.Labels{"DestinationHost": host.HostName}).Set(hoststat)
			if host.HostIP == Configuration.TargetIP {
				if IsTarget_Up != hoststat {
					IsTarget_Up = hoststat
				}
			}
		}
	}
}

func linkInquiry(host, port string) (linkStatus float64) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		//fmt.Println("Connecting error:", err)
		linkStatus = 2 //down
		return
	}
	if conn != nil {
		defer conn.Close()
		//fmt.Println("Opened", net.JoinHostPort(host, port))
		linkStatus = 1 //up
		return
	} else {
		linkStatus = 2 //down
		return
	}
}
