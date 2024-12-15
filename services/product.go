package services

import "product-management-system/models"

type ProductInput struct {
	UserID        uint
	Name          string
	Description   string
	Price         float64
	ProductImages []string
}

func CreateProduct(input ProductInput) (models.Product, error) {
	product := models.Product{
		UserID:        input.UserID,
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		ProductImages: input.ProductImages,
	}
	// Save product logic goes here (e.g., db.Create(&product))
	return product, nil
}
