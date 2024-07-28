package presenter

import (
	"github.com/felipeganho/to-do-list/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

// TodoSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func TodoSuccessResponse(data []entities.Todo) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
