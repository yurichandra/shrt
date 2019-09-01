package handler

import (
	"net/http"

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

	return router
}

// Store saves and return new or existing URL.
func (h *ShortenerHandler) Store(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
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

	data["originalURL"] = request.OriginalURL
	data["apiKey"] = apiKey

	url, err := h.url.Shorten(data, auth)
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
