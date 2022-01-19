package http

import (
	"github.com/MrTj458/fiber-api-todo/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) registerTodoRoutes() {
	r := s.app.Group("/api/todos")

	r.Get("/", s.handleTodosIndex)
	r.Get("/:id", s.handleTodoById)
	r.Post("/", s.handleTodoCreate)
	r.Put("/:id", s.handleTodoUpdate)
	r.Delete("/:id", s.handleTodoDelete)
}

// @Summary      Get all Todos
// @Description  Get all Todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Todo
// @Failure      500  {object}  model.Error
// @Router       /api/todos [get]
func (s *Server) handleTodosIndex(c *fiber.Ctx) error {
	todos, err := s.TodoService.AllTodos()
	if err != nil {
		return model.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.JSON(todos)
}

// @Summary      Get Todo by ID
// @Description  Get Todo by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  model.Todo
// @Failure      400  {object}  model.Error
// @Failure      404  {object}  model.Error
// @Router       /api/todos/{id} [get]
func (s *Server) handleTodoById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		// return c.Status(fiber.StatusBadRequest).JSON(model.NewError("invalid UUID"))
		return model.NewError(fiber.StatusBadRequest, "invalid UUID")
	}

	todo, err := s.TodoService.FindTodoByID(id)
	if err != nil {
		// return c.Status(fiber.StatusNotFound).JSON(model.NewError("todo not found"))
		return model.NewError(fiber.StatusNotFound, "todo not found")
	}

	return c.JSON(todo)
}

// @Summary      Create Todo
// @Description  Create Todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      model.TodoCreate  true  "Add Todo"
// @Success      201   {object}  model.Todo
// @Failure      422   {object}  model.Error
// @Failure      500   {object}  model.Error
// @Router       /api/todos [post]
func (s *Server) handleTodoCreate(c *fiber.Ctx) error {
	todo := model.TodoCreate{}

	// Decode JSON
	if err := c.BodyParser(&todo); err != nil {
		return model.NewError(fiber.StatusUnprocessableEntity, "invalid JSON object received")
	}

	// Validate Todo
	if errors, ok := todo.Validate(); !ok {
		return model.NewErrorWithFields(fiber.StatusUnprocessableEntity, "Invalid Todo object received", errors)
	}

	// Create new todo
	newTodo, err := s.TodoService.CreateTodo(todo)
	if err != nil {
		return model.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

// @Summary      Update Todo
// @Description  Update Todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @param        id    path      string            true  "Todo ID"
// @Param        todo  body      model.TodoUpdate  true  "Update Todo"
// @Success      200   {object}  model.Todo
// @Failure      400   {object}  model.Error
// @Failure      422   {object}  model.Error
// @Failure      500   {object}  model.Error
// @Router       /api/todos/{id} [put]
func (s *Server) handleTodoUpdate(c *fiber.Ctx) error {
	// Get ID from URL param
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return model.NewError(fiber.StatusBadRequest, "invalid UUID")
	}

	// Decode JSON
	upd := model.TodoUpdate{}
	if err := c.BodyParser(&upd); err != nil {
		return model.NewError(fiber.StatusUnprocessableEntity, "invalid JSON object received")
	}

	// Validate TodoUpdate object
	if errors, ok := upd.Validate(); !ok {
		return model.NewErrorWithFields(fiber.StatusUnprocessableEntity, "invalid TodoUpdate object received", errors)
	}

	// Update Todo
	todo, err := s.TodoService.UpdateTodo(id, upd)
	if err != nil {
		return model.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.JSON(todo)
}

// @Summary      Delete Todo
// @Description  Delete Todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @param        id   path  string  true  "Todo ID"
// @Success      204  "Todo Deleted"
// @Failure      400  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Router       /api/todos/{id} [delete]
func (s *Server) handleTodoDelete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		// return c.Status(fiber.StatusBadRequest).JSON(model.NewError("invalid UUID received"))
		return model.NewError(fiber.StatusBadRequest, "invalid UUID received")
	}

	if err := s.TodoService.DeleteTodo(id); err != nil {
		// return c.Status(fiber.StatusInternalServerError).JSON(model.NewError("internal server error"))
		return model.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
