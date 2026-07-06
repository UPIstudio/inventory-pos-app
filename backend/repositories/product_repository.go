package repositories

import (
	"github.com/luthfi/inventory-pos-app/backend/config"
	"github.com/luthfi/inventory-pos-app/backend/models"
)

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}
