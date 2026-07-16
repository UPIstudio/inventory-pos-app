package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/inventory-pos-app/backend/models"
	"github.com/luthfi/inventory-pos-app/backend/repositories"
)

type productHandler struct{}

func GetProducts(c *fiber.Ctx) error {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var products models.Product
	if err := c.BodyParser(&products); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format JSON salah"})
	}

	if err := repositories.SaveProduct(&products); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal Menyimpan Produk"})
	}

	return c.Status(201).JSON(products)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var products models.Product
	if err := c.BodyParser(&products); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format JSON salah"})
	}

	products.ID = uint(id)
	if err := repositories.UpdateProduct(&products); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupdate data"})
	}

	return c.JSON(products)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID Tidak Valid"})
	}

	if err := repositories.DeleteProduct(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data"})
	}

	return c.JSON(fiber.Map{"message": "Produk berhasil dihapus"})
}
