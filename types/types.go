package types

import "time"

type OrderStore interface {
	GetAllOrders() ([]Order, error)
	SaveOrder(newOrder *Order) error
}

type Order struct {
	OrderUID          string      `json:"order_uid" validate:"required"`
	TrackNumber       string      `json:"track_number" validate:"required"`
	Entry             string      `json:"entry" validate:"required"`
	Delivery          Delivery    `json:"delivery" validate:"required"`
	Payment           Payment     `json:"payment" validate:"required"`
	Items             []OrderItem `json:"items" validate:"required"`
	Locale            string      `json:"locale" validate:"required"`
	InternalSignature string      `json:"internal_signature" validate:"omitempty"`
	CustomerID        string      `json:"customer_id" validate:"required"`
	DeliveryService   string      `json:"delivery_service" validate:"required"`
	ShardKey          string      `json:"shard_key" validate:"required"`
	SmID              int         `json:"sm_id" validate:"required"`
	DateCreated       time.Time   `json:"date_created" validate:"required"`
	OofShard          string      `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id" validate:"omitempty"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
	PaymentDT    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee" validate:"omitempty"`
}

type OrderItem struct {
	ChrtID      int    `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"required"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required"`
	NmID        int    `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}
