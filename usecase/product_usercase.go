package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{productRepository: repository}
}

func (pu ProductUsecase) GetProducts() ([]model.Product, error) {
	products, err := pu.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pu ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.productRepository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	product, err := pu.productRepository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (pu ProductUsecase) UpdateProduct(product model.Product) (model.Product, error) {
	product, err := pu.productRepository.UpdateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (pu ProductUsecase) DeleteProduct(id int) error {
	err := pu.productRepository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (pu ProductUsecase) GetProductByName(name string) (model.Product, error) {
	product, err := pu.productRepository.GetProductByName(name)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
