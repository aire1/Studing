package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Produce(topic string, msg kafka.Message) {
	// Создаём продюсер
	writer := kafka.NewWriter(kafka.WriterConfig{
		// Brokers:  []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"}, // Адрес Kafka
		Brokers:  []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:    topic,               // Название топика
		Balancer: &kafka.LeastBytes{}, // Балансировка нагрузки
	})
	defer writer.Close()

	// msg := kafka.Message{
	// 	Key:   []byte("reg"),
	// 	Value: []byte("Hello, kafka!"),
	// }

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatal("Ошибка при отправке:", err)
	}
	fmt.Println("Сообщение отправлено")
}
