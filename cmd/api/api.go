package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bjhermid/go-api-1/services/cart"
	"github.com/bjhermid/go-api-1/services/order"
	"github.com/bjhermid/go-api-1/services/products"
	"github.com/bjhermid/go-api-1/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// run
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//user
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	//products
	productStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	//cart
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore,productStore,userStore)
	cartHandler.RegisterRoutes(subrouter)

	//debug
	log.Println("Listen on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
