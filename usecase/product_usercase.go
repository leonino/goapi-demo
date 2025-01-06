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

func (pu ProductUsecase) GetProduct(id int) (model.Product, error) {
	product, err := pu.productRepository.GetProduct(id)
	if err != nil {
		return model.Product{}, err
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
