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
	nextId   uint64
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{
		products: make(map[uint64]recomendation.Product),
	}
}

func (service *DummyProductService) Describe(productID uint64) (*recomendation.Product, error) {
	if value, ok := service.products[productID]; ok {
		return &value, nil
	} else {
		return nil, fmt.Errorf("DummyProductService.Describe: cannot find product with %d id", productID)
	}
}

func (service *DummyProductService) List(cursor uint64, limit uint64) ([]recomendation.Product, error) {
	mapSize := len(service.products)

	if cursor > uint64(mapSize) {
		return nil, fmt.Errorf("DummyProductService.List: cursor out of range")
	}

	productsList := make([]recomendation.Product, 0, mapSize)

	var pos uint64

	for _, product := range service.products {
		if pos >= cursor && pos < cursor+limit {
			productsList = append(productsList, product)
		}
		pos++
		if pos >= cursor+limit {
			break
		}
	}

	return productsList, nil
}

func (service *DummyProductService) Create(r recomendation.Product) (uint64, error) {

	service.nextId++

	r.Id = service.nextId

	service.products[r.Id] = r

	return service.nextId, nil
}

func (service *DummyProductService) Update(productID uint64, product recomendation.Product) error {
	if _, ok := service.products[productID]; !ok {
		return fmt.Errorf("DummyProductService.Update: product %d does not exist", productID)
	}

	service.products[productID] = product

	return nil
}

func (service *DummyProductService) Remove(productID uint64) (bool, error) {
	var ok bool

	if _, ok = service.products[productID]; !ok {
		return false, fmt.Errorf("DummyProductService.Remove: product %d does not exist", productID)
	}

	delete(service.products, productID)

	return ok, nil
}
