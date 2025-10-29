package products

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/bjhermid/go-api-1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil

}

func (s *Store) GetProductByIDs(productIds []int) ([]types.Product, error) {
	placeHolder := strings.Repeat(",?", len(productIds)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeHolder)

	//convert ProductID to []interface{}
	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)

	}
	return products, nil

}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec(
		"UPDATE products SET name = ?, price = ?, image = ?, description = ?, quantity = ? WHERE id = ?",
		product.Name, product.Price, product.Image, product.Descriptions, product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil

}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Descriptions,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
