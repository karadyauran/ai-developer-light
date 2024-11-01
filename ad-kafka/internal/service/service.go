package service

type Service struct {
	KafkaProducer *KafkaProducer
	// kafkaConsumer *KafkaConsumer
}

func NewService(brokers string) *Service {
	return &Service{
		KafkaProducer: NewKafkaProducer(brokers),
		// kafkaConsumer: NewKafkaConsumer(brokers, groupID),
	}
}
