package order

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prok05/wb-level-0/types"
	"log"
)

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(dbpool *pgxpool.Pool) *Store {
	return &Store{
		pool: dbpool,
	}
}

func (s *Store) SaveOrder(newOrder *types.Order) error {
	query := "INSERT INTO orders (order_uid, track_number, entry, delivery_info, payment_info, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
	_, err := s.pool.Exec(context.Background(), query,
		newOrder.OrderUID, newOrder.TrackNumber, newOrder.Entry, newOrder.Delivery, newOrder.Payment, newOrder.Items, newOrder.Locale, newOrder.InternalSignature,
		newOrder.CustomerID, newOrder.DeliveryService, newOrder.ShardKey, newOrder.SmID, newOrder.DateCreated, newOrder.OofShard)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Store) GetAllOrders() ([]types.Order, error) {
	query := "SELECT * FROM orders"
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var orders []types.Order

	for rows.Next() {
		var order types.Order
		if err := rows.Scan(&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Delivery,
			&order.Payment,
			&order.Items,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
