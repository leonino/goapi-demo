package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

type ProductRepositoryInterface interface {
	GetProducts() ([]model.Product, error)
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	products := []model.Product{}
	const query = "SELECT * FROM product"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query := "SELECT * FROM product WHERE id = $1"
	row := pr.connection.QueryRow(query, id)

	product := model.Product{}
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"

	err := pr.connection.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (model.Product, error) {
	query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"

	_, err := pr.connection.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM product WHERE id = $1"

	_, err := pr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (pr *ProductRepository) GetProductByName(name string) (model.Product, error) {
	query := "SELECT * FROM product WHERE product_name = $1"
	row := pr.connection.QueryRow(query, name)

	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}
