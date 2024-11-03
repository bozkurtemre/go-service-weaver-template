package frontend

import (
	"context"

	"github.com/bozkurtemre/go-service-weaver-template/internal/example"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	weaver.Implements[weaver.Main]
	example  weaver.Ref[example.Service]
	frontend weaver.Listener `weaver:"frontend"`

	f *fiber.App
}

func Serve(ctx context.Context, s *Server) (err error) {
	s.f = fiber.New()

	router := s.f.Use(logger.New(logger.ConfigDefault))

	router.Get("/", s.HelloWorld)

	grpExamples := router.Group("/examples")
	grpExamples.Get("/", s.GetAllExamples)
	grpExamples.Get("/:id", s.GetExampleById)
	grpExamples.Post("/", s.CreateExample)

	return s.f.Listener(s.frontend)
}
