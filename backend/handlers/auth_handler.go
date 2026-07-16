package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/inventory-pos-app/backend/models"
	"github.com/luthfi/inventory-pos-app/backend/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format data salah"})
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format data salah"})
	}

	if err := h.authService.Register(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal registrasi akun"})
	}

	return c.JSON(fiber.Map{"message": "User berhasil dibuat, silahkan login"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	val := c.Locals("userId")

	if val == nil {
		return c.Status(401).JSON(fiber.Map{"error": "User ID tidak ditemukan di sesi"})
	}

	userID, ok := val.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Format User ID tidak valid"})
	}

	err := h.authService.Logout(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal logout"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Logout berhasil"})
}
