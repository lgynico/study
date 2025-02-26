package eventdriven

// 基于 kafka 的事件驱动
func kafkaProducer() {
	// producer, err := kafka.NewProducer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// })

	// if err != nil {
	// 	panic(err)
	// }
	// defer producer.Close()

	// topic := "user-events"
	// err = producer.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
	// 	Value:          []byte("UserRegistered:Nico"),
	// }, nil)

	// if err != nil {
	// 	panic(err)
	// }
}

func kafkaComsumer() {
	// consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// 	"group.id":          "user-event-group",
	// 	"auto.offset.reset": "earliest",
	// })
	// defer consumer.Close()

	// consumer.Subscribe("user-events", nil)

	// for {
	// 	msg, _ := consumer.ReadMessage(-1)
	// 	fmt.Println("[Consumer] Received event: ", string(msg.Value))
	// }
}
