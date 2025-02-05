package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

// Producer хранит писатели для разных топиков
type Producer struct {
	writers map[string]*kafka.Writer
	mu      sync.RWMutex
}

var (
	KafkaProducer *Producer
)

// Создаём новый продюсер
func NewProducer(brokers []string, topics []string) *Producer {
	producer := &Producer{
		writers: make(map[string]*kafka.Writer),
	}

	// Создаём writer'ов для каждого топика
	for _, topic := range topics {
		producer.writers[topic] = &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}
	}

	return producer
}

// Метод для отправки сообщений
func (p *Producer) Produce(topic string, msg kafka.Message) {
	p.mu.RLock()
	writer, exists := p.writers[topic]
	p.mu.RUnlock()

	if !exists {
		log.Printf("Ошибка: продюсер для топика '%s' не создан\n", topic)
		return
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println("Ошибка при отправке:", err)
	} else {
		fmt.Println("Сообщение отправлено в топик:", topic)
	}
}

// Закрываем всех writer'ов при завершении работы
func (p *Producer) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for topic, writer := range p.writers {
		writer.Close()
		delete(p.writers, topic)
		fmt.Println("Закрыт продюсер для топика:", topic)
	}
}

func Init() {
	brokers := []string{"localhost:19092", "localhost:19094", "localhost:19096"}
	topics := []string{"registrations", "get_authorizations", "check_authorizations"}

	KafkaProducer = NewProducer(brokers, topics)
}
