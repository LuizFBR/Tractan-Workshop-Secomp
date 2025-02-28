package main
import (
    "github.com/segmentio/kafka-go"
    "fmt"
    "context"
    "log"
)

func main(){
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"54.160.9.28"},
        GroupID:   "secomp",
        Topic:     "secomp",
        MaxBytes:  10e6, // 10MB
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
    }

    if err := r.Close(); err != nil {
        log.Fatal("failed to close reader:", err)
    }
}