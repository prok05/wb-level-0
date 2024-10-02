package consumer

import (
	"context"
	"encoding/json"
	"github.com/prok05/wb-level-0/cache"
	"github.com/prok05/wb-level-0/config"
	"github.com/prok05/wb-level-0/service/order"
	"github.com/prok05/wb-level-0/types"
	"github.com/prok05/wb-level-0/utils"
	"github.com/segmentio/kafka-go"
	"log"
)

type OrderConsumer struct {
	reader *kafka.Reader
	store  *order.Store
	cache  *cache.OrderCache
}

func New(store *order.Store, cache *cache.OrderCache) *OrderConsumer {
	return &OrderConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{config.Envs.KafkaURL},
			Topic:       config.Envs.KafkaTopic,
			GroupID:     "order-group",
			StartOffset: kafka.LastOffset,
		}),
		store: store,
		cache: cache,
	}
}

func (oc *OrderConsumer) Run(ctx context.Context) error {
	for {
		log.Println("Listening to messages...")
		msg, err := oc.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error while reading messages: %v", err)
			return err
		}

		log.Printf("Received message: %s", msg.Value)

		// парсинг сообщения
		var newOrder types.Order
		if err := json.Unmarshal(msg.Value, &newOrder); err != nil {
			log.Printf("JSON unmarshall error: %v", err)
			continue
		}

		// валидация
		if err := utils.Validate.Struct(newOrder); err != nil {
			log.Printf("Order validation error: %v", err)
			continue
		}

		// сохранение заказа в БД
		err = oc.store.SaveOrder(&newOrder)
		if err != nil {
			log.Printf("error while saving order to database: %v", err)
			continue
		}

		// сохранение заказа к кэш
		oc.cache.Set(newOrder.OrderUID, newOrder)
	}
}
