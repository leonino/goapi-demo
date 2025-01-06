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
	const sql = "SELECT * FROM product"

	rows, err := pr.connection.Query(sql)
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

func (pr *ProductRepository) GetProduct(id int) (model.Product, error) {
	sql := "SELECT * FROM product WHERE id = $1"
	row := pr.connection.QueryRow(sql, id)

	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	sql := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"

	err := pr.connection.QueryRow(sql, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (model.Product, error) {
	sql := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"

	_, err := pr.connection.Exec(sql, product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	sql := "DELETE FROM product WHERE id = $1"

	_, err := pr.connection.Exec(sql, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (pr *ProductRepository) GetProductByName(name string) (model.Product, error) {
	sql := "SELECT * FROM product WHERE product_name = $1"
	row := pr.connection.QueryRow(sql, name)

	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}
