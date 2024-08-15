package product

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/recomendation"
)

type ProductService interface {
	Describe(productID uint64) (*recomendation.Product, error)
	List(cursor uint64, limit uint64) ([]recomendation.Product, error)
	Create(recomendation.Product) (uint64, error)
	Update(productID uint64, product recomendation.Product) error
	Remove(productID uint64) (bool, error)
}

type DummyProductService struct {
	products map[uint64]recomendation.Product
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}

func (service *DummyProductService) Describe(productID uint64) (*recomendation.Product, error) {
	if value, ok := service.products[productID]; ok {
		return &value, nil
	} else {
		return nil, fmt.Errorf("DummyProductService.Describe: cannot find product with %d id", productID)
	}
}

func (service *DummyProductService) List(cursor uint64, limit uint64) ([]recomendation.Product, error) {
	return nil, nil
}

func (service *DummyProductService) Create(recomendation.Product) (uint64, error) {
	return 0, nil
}

func (service *DummyProductService) Update(productID uint64, product recomendation.Product) error {
	return nil
}

func (service *DummyProductService) Remove(productID uint64) (bool, error) {
	return false, nil
}
