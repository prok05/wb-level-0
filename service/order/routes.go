package order

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prok05/wb-level-0/cache"
	"github.com/prok05/wb-level-0/types"
	"github.com/prok05/wb-level-0/utils"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	store types.OrderStore
	cache *cache.OrderCache
}

func NewHandler(store types.OrderStore, cache *cache.OrderCache) *Handler {
	return &Handler{store: store, cache: cache}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders/{orderID}", h.handleGetOrderByID).Methods("GET")
	router.HandleFunc("/", h.handleIndex)
}

func (h *Handler) handleGetOrderByID(w http.ResponseWriter, r *http.Request) {
	// получение orderID
	vars := mux.Vars(r)
	orderID, ok := vars["orderID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("order ID is required"))
		return
	}

	// получение заказа из кэша
	order, found := h.cache.Get(orderID)
	if !found {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("order not found"))
		return
	}

	// отправление заказа
	utils.WriteJSON(w, http.StatusOK, order)
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	temlp := template.Must(template.ParseFiles("./templates/index.html"))
	if err := temlp.Execute(w, nil); err != nil {
		log.Printf("Unable to render template")
	}
}
