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

	s.app.Static("/", "./public")

	s.app.Get("/swagger/*", swagger.HandlerDefault)

	s.registerTodoRoutes()

	// All other routes, send index.html
	// This is useful for using single page applications like React or Vue
	s.app.Get("/*", clientAppHandler)

	return s
}

func (s *Server) Run() error {
	return s.app.Listen(":" + strconv.Itoa(s.port))
}

func clientAppHandler(ctx *fiber.Ctx) error {
	return ctx.SendFile("public/index.html")
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*model.Error); ok {
		return ctx.Status(e.Code).JSON(e)
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}
