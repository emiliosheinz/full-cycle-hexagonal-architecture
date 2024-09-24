package application

type ProductService struct {
	Persistence ProductPersistanceInterface
}

func (s *ProductService) Ger(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
