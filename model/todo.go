package model

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TodoService interface {
	FindTodoByID(id uuid.UUID) (*Todo, error)
	AllTodos() ([]*Todo, error)
	CreateTodo(todo TodoCreate) (*Todo, error)
	UpdateTodo(id uuid.UUID, upd TodoUpdate) (*Todo, error)
	DeleteTodo(id uuid.UUID) error
}

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Desc      string    `json:"description"`
	Completed bool      `json:"completed"`
}

func (t *Todo) Validate() ([]*ErrorField, bool) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(t)
	if err != nil {
		errors := []*ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorField{
				Location: err.Field(),
				Type:     err.Type().String(),
				Detail:   err.ActualTag(),
			})
		}

		return errors, false
	}

	return nil, true
}

type TodoCreate struct {
	Desc string `json:"description" validate:"required"`
}

func (t *TodoCreate) Validate() ([]*ErrorField, bool) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(t)
	if err != nil {
		errors := []*ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorField{
				Location: err.Field(),
				Type:     err.Type().String(),
				Detail:   err.ActualTag(),
			})
		}

		return errors, false
	}

	return nil, true
}

type TodoUpdate struct {
	Desc      string `json:"description"`
	Completed bool   `json:"completed"`
}

func (t *TodoUpdate) Validate() ([]*ErrorField, bool) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(t)
	if err != nil {
		errors := []*ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorField{
				Location: err.Field(),
				Type:     err.Type().String(),
				Detail:   err.ActualTag(),
			})
		}

		return errors, false
	}

	return nil, true
}
