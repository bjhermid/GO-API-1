package order

import (
	"database/sql"

	"github.com/bjhermid/go-api-1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders(userID, total, status, address) VALUES (?,?,?,?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

func (s *Store) CreateOrderItem(order types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items(id, orderID, productID, quantity, price) VALUES (?,?,?,?,?)", order.ID, order.OrderID, order.ProductID, order.Quantity, order.Price)
	
	return err
}

