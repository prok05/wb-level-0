package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/prok05/wb-level-0/types"
	"net/http"
	"time"
)

var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func GenerateOrder() types.Order {
	return types.Order{
		OrderUID:    uuid.NewString(),
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: types.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: types.Payment{
			Transaction:  uuid.NewString(),
			RequestID:    "",
			Currency:     "RUB",
			Provider:     "wbpay",
			Amount:       534,
			PaymentDT:    9239123,
			DeliveryCost: 1500,
			Bank:         "alpha",
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []types.OrderItem{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Maracas",
				Sale:        30,
				Size:        "0",
				TotalPrice:  332,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}
