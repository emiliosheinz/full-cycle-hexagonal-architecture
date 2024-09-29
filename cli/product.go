package cli

import (
	"fmt"

	"github.com/emiliosheinz/exagonal-architecture/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	productPrice float64,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s, price %f, and status %s has been created",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product %s has been enabled",
			product.GetId(),
		)

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product %s has been disabled",
			product.GetId(),
		)
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetId(),
			res.GetName(),
			res.GetPrice(),
			res.GetStatus(),
		)
	}

	return result, nil
}
