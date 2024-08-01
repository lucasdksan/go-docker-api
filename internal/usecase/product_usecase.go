package usecase

import (
	"go-docker-api/internal/model"
	"go-docker-api/internal/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

func (puc *ProductUsecase) GetProducts() ([]model.Product, error) {
	return puc.repository.GetProducts()
}

func (puc *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := puc.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, nil
	}

	product.ID = productId

	return product, nil
}

func (puc *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := puc.repository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (puc *ProductUsecase) DeleteProduct(id_product int) error {
	if err := puc.repository.DeleteProduct(id_product); err != nil {
		return err
	}

	return nil
}
