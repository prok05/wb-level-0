package api

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prok05/wb-level-0/cache"
	"github.com/prok05/wb-level-0/service/order"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	pool  *pgxpool.Pool
	cache *cache.OrderCache
}

func NewAPIServer(addr string, pool *pgxpool.Pool, cache *cache.OrderCache) *APIServer {
	return &APIServer{
		addr:  addr,
		pool:  pool,
		cache: cache,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	orderStore := order.NewStore(s.pool)
	orderHandler := order.NewHandler(orderStore, s.cache)
	orderHandler.RegisterRoutes(subrouter)

	log.Println("Listening on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}
