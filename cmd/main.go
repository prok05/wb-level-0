package main

import (
	"context"
	"github.com/prok05/wb-level-0/cache"
	"github.com/prok05/wb-level-0/cmd/api"
	"github.com/prok05/wb-level-0/config"
	"github.com/prok05/wb-level-0/db"
	"github.com/prok05/wb-level-0/service/consumer"
	"github.com/prok05/wb-level-0/service/order"
	"log"
)

func main() {
	// подключение к БД
	pgStorage := db.NewPgStorage()
	pool, err := pgStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// инициализация кэша
	orderCache := cache.New(-1, -1, order.NewStore(pool))

	// восстановление кэша из БД
	if err := orderCache.RestoreCache(); err != nil {
		log.Fatalf("error while restoring cache: %v", err)
	}

	// инициализация consumer
	orderConsumer := consumer.New(order.NewStore(pool), orderCache)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := orderConsumer.Run(ctx); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// запуск сервера
	server := api.NewAPIServer(config.Envs.ServerPort, pool, orderCache)
	if err := server.Run(); err != nil {
		log.Fatalf("error while starting server %v", err)
	}
}
