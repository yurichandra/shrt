package object

import (
	"errors"
	"net/http"
)

// UserRequest represent object request of User model.
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse represent object response of User model.
type UserResponse struct {
	Key string `json:"api_key"`
}

// Bind doing validation for request object.
func (req *UserRequest) Bind(r *http.Request) error {
	if req.Email == "" {
		return errors.New("`email` can't be empty")
	}

	if req.Password == "" {
		return errors.New("`password` can't be empty")
	}

	return nil
}

// Render renders url model.
func (res *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
