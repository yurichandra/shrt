package handler

import (
	"net/http"
	"time"

	"github.com/yurichandra/shrt/object"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/shrt/service"
)

// ShortenerHandler represent handler of shortener.
type ShortenerHandler struct {
	url service.ShortenerServiceContract
}

// GetRoutes return all routes of URL.
func (h *ShortenerHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", h.Store)
	router.Get("/{shortURL}", h.Find)

	return router
}

// Store saves and return new or existing URL.
func (h *ShortenerHandler) Store(w http.ResponseWriter, r *http.Request) {
	var auth bool

	request := object.URLRequest{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))

		return
	}

	apiKey := r.Header.Get("api_key")
	if apiKey != "" {
		auth = true
	}

	url, err := h.url.Shorten(request.OriginalURL, apiKey, time.Now(), auth)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}

	err = render.Render(w, r, object.CreateURLResponse(url))
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))

		return
	}
}

// Find finds url by short url.
func (h *ShortenerHandler) Find(w http.ResponseWriter, r *http.Request) {
	shortURLParam := chi.URLParam(r, "shortURL")
	url, err := h.url.Find(shortURLParam)
	if err != nil {
		render.Render(w, r, sendNotFoundResponse(err.Error()))
		return
	}

	err = render.Render(w, r, object.CreateURLResponse(url))
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))
		return
	}
}
