package repositories

import (
	"github.com/luthfi/inventory-pos-app/backend/config"
	"github.com/luthfi/inventory-pos-app/backend/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}
func SaveProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
