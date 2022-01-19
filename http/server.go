package http

import (
	"strconv"

	_ "github.com/MrTj458/fiber-api-todo/docs"
	"github.com/MrTj458/fiber-api-todo/model"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app  *fiber.App
	port int

	TodoService model.TodoService
}

func NewServer(port int) *Server {
	s := &Server{
		app: fiber.New(fiber.Config{
			ErrorHandler: errorHandler,
		}),
		port: port,
	}

	s.app.Use(logger.New())

	s.app.Get("/swagger/*", swagger.HandlerDefault)

	s.registerTodoRoutes()

	return s
}

func (s *Server) Run() error {
	return s.app.Listen(":" + strconv.Itoa(s.port))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*model.Error); ok {
		ctx.Status(e.Code).JSON(e)
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}
