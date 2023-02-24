package product

import (
	"errors"
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(index int) (*Product, error) {
	if index < 0 || len(allProducts) < index {
		return nil, errors.New(fmt.Sprintf("Fail to find product with index: %d", index))
	}
	return &allProducts[index], nil
}
