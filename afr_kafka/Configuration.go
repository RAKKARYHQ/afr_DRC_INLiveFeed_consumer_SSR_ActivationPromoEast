package afr_kafka

var Configuration ConfigType

type ConfigType struct {
	HttpServicePort string
	HostId          string

	KafkaBrokerUrls string
	KafkaTopics     string
	KafkaClientId   string

	TargetProtocol string
	TargetIP       string
	TargetPort     string
	TargetURI      string
}

func GetDefaultConfiguration() (err error) {
	//Configuration = setDefaultConfiguration_Dev_local()
	Configuration = setDefaultConfiguration()
	return nil
}

func setDefaultConfiguration() (Configuration ConfigType) {
	Configuration.HttpServicePort = "9662"
	Configuration.HostId = "SSR_ActivatioPromo Host"

	//Configuration.KafkaBrokerUrls = "localhost:9092,localhost:9192,localhost:9292"
	Configuration.KafkaBrokerUrls = "kafka1:9092,kafka2:9092,kafka3:9092"
	Configuration.KafkaClientId = "SR_ActivatioPromo"

	// http://10.100.11.21:50000/LiveFeedActivationPromo/ActivationEastPromo
	Configuration.TargetProtocol = "http"
	Configuration.TargetIP = "10.100.11.21"
	Configuration.TargetPort = "50000"
	Configuration.TargetURI = "/LiveFeedActivationPromo/ActivationEastPromo"
	return
}
