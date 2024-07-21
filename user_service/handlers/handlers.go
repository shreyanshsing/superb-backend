package handlers

import (
	"log"

	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"superb.com/v0/user_service/db-config"
	models "superb.com/v0/user_service/models"
	"superb.com/v0/user_service/utils"
)

type CreateUserInput struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
}
 
func CreateUser(c *fiber.Ctx) error {
	// Parse request body
	var input CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate input
	validate := validate.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create user
	username := utils.GenerateRandomUserName(input.FullName)
	user := models.UserModel{
		Password: input.Password,
		Email:    input.Email,
		FullName: input.FullName,
		Username: username,
		Intersets: pq.Int32Array{},
		Communities: pq.Int32Array{},
		Posts: pq.Int32Array{},
		Rating: 0,
		Level: 0,
		AboutMe: "",
		Links: pq.StringArray{},
		PublicProfileURL: "",
	}

	// Save to database
	db := database.Database

	result := db.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	log.Print("User created successfully")

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"data": user.ToUserDto(),
	})
}

func GetUser(c *fiber.Ctx) error {
	// Get user ID from URL
	id := c.Params("id")

	// Find user
	db := database.Database
	var user models.UserModel
	result := db.First(&user, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": user.ToUserDto(),
		"message": "User found",
	})
}

