package frontend

import (
	"strconv"

	"github.com/bozkurtemre/go-service-weaver-template/internal/example"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s Server) HelloWorld(c *fiber.Ctx) (err error) {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello, World!"})
}

func (s Server) GetAllExamples(c *fiber.Ctx) (err error) {
	out, err := s.example.Get().AllExamples(c.Context())
	if err != nil {
		log.Errorf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (s Server) GetExampleById(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameter:"+c.Params("id"))
	}
	out, err := s.example.Get().GetOneExampleById(c.Context(), int64(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (s Server) CreateExample(c *fiber.Ctx) (err error) {
	var body example.ExampleIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	_, err = s.example.Get().CreateExample(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}
