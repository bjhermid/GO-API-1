package cart

import (
	"fmt"
	"net/http"

	"github.com/bjhermid/go-api-1/types"
	"github.com/bjhermid/go-api-1/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	//sort object from front-End
	var cart types.CartCheckoutPayload
	if err := utils.ParseJson(r,cart); err != nil{
		utils.WriteError(w,http.StatusBadRequest,err)
	}

	if err:= utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	//get product
}
