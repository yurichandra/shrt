package object

import (
	"github.com/go-chi/render"
	"github.com/yurichandra/shrt/model"
)

// CreateUserResponse creating object response for user model.
func CreateUserResponse(user model.User) render.Renderer {
	return &UserResponse{
		Key: user.Key,
	}
}
