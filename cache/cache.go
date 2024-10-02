package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/prok05/wb-level-0/types"
	"log"
	"time"
)

type OrderCache struct {
	c     *cache.Cache
	store types.OrderStore
}

func New(defaultExpiration, cleanupInterval time.Duration, store types.OrderStore) *OrderCache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &OrderCache{c: c, store: store}
}

func (c *OrderCache) Set(key string, value types.Order) {
	c.c.Set(key, value, cache.DefaultExpiration)
}

func (c *OrderCache) Get(key string) (*types.Order, bool) {
	value, found := c.c.Get(key)
	if !found {
		return nil, false
	}
	order, ok := value.(types.Order)
	if !ok {
		return nil, false
	}
	return &order, true
}

func (c *OrderCache) RestoreCache() error {
	orders, err := c.store.GetAllOrders()
	if err != nil {
		log.Println("error while getting orders")
		return err
	}

	if len(orders) == 0 {
		log.Println("Cache: database is empty, nothing to restore")
		return nil
	} else {
		for _, order := range orders {
			c.Set(order.OrderUID, order)
		}
		log.Println("Cache: cache restored, items count:", c.c.ItemCount())

		return nil
	}
}
