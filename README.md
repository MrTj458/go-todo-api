# Go Todo Rest API

This is an example REST API made with [Go](https://go.dev/) and the [Fiber](https://github.com/gofiber/fiber) framework.

It includes validation using the [Validator](https://github.com/go-playground/validator) package and returns developer-friendly error messages:

```json
{
  "code": 422,
  "detail": "Invalid Todo object received",
  "fields": [
    {
      "location": "description",
      "type": "string",
      "detail": "required"
    }
  ]
}
```

It also uses the [fiber-swagger](https://github.com/arsmn/fiber-swagger) package along with [swag](https://github.com/swaggo/swag) to auto-generate Swagger documentation:

```go
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
```
