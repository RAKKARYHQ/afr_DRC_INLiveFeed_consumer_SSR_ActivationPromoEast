package afr_kafka

type UserControl struct {
	//MongoDB     *MongoDB
	KafkaClient *KafkaClient
}

func NewUserControl() *UserControl {
	UC := &UserControl{
		//MongoDB:     NewMongoDBClient(),
		KafkaClient: NewKafkaClient(),
	}
	return UC
}
