package cart

import (
	"fmt"

	"github.com/bjhermid/go-api-1/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error){
	productsIDs := make([]int,len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d,", item.ProductID)
		}

		productsIDs[i] = item.ProductID
	}
	return productsIDs, nil
}

//create order in access to repository to make table association
func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error){
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	//check if product actually available (as a stock)
	if  err := checkIfCartInStock(items, productMap); err!= nil {
		return 0,0, nil
	}
	//1. calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	//reduce quanitity products in our db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	
	// create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID: userID,
		Total: totalPrice,
		Status: "pending",
		Address: "some address",
	})
	if err != nil {
		return 0,0,err
	}

	// create the order_items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID: orderID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			Price: productMap[item.ProductID].Price,
		})
	}

	return orderID,totalPrice, nil
}

func checkIfCartInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity request", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice (cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}

