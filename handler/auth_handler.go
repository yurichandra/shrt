package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/shrt/object"
	"github.com/yurichandra/shrt/service"
)

// AuthHandler represent handler of auth.
type AuthHandler struct {
	auth service.AuthServiceContract
}

// GetRoutes available for auth handler.
func (h *AuthHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/authorize", h.Authorize)
	router.Post("/authenticate", h.Authenticate)

	return router
}

// Authorize the request.
func (h *AuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	req := object.UserRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))
		return
	}

	user, err := h.auth.Authorize(req.Email, req.Password)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}

	render.Status(r, http.StatusCreated)
	err = render.Render(w, r, object.CreateUserResponse(user))
}

// Authenticate the request.
func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	req := object.UserRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))
		return
	}

	user, err := h.auth.Authenticate(req.Email, req.Password)
	if err != nil {
		render.Render(w, r, sendUnauthorizedResponse(err.Error()))
		return
	}

	render.Render(w, r, object.CreateUserResponse(user))
}
