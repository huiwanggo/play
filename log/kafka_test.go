package log

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

var address = "127.0.0.1:19092"

func TestTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		fmt.Printf("partition: %d, topic: %s \n", p.ID, p.Topic)
	}
}

func TestWrite(t *testing.T) {
	w := kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    "test",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     "test",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		panic(err)
	}
}
