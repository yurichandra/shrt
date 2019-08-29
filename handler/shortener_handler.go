package handler

import (
	"net/http"
	"strconv"

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

	router.Get("/", h.Get)
	router.Get("/{id}", h.Find)
	router.Post("/", h.Store)

	return router
}

// Get return all available urls.
func (h *ShortenerHandler) Get(w http.ResponseWriter, r *http.Request) {
	urls := h.url.Get()

	urlListResponse := object.CreateURLListResponse(urls)
	err := render.RenderList(w, r, urlListResponse)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))

		return
	}
}

// Find return single url.
func (h *ShortenerHandler) Find(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	url := h.url.Find(uint(id))
	if url.ID == 0 {
		render.Render(w, r, sendNotFoundResponse("URL with certain ID is not found"))

		return
	}

	urlResponse := object.CreateURLResponse(url)

	err := render.Render(w, r, urlResponse)
	if err != nil {
		render.Render(w, r, sendInternalServerErrorResponse(err.Error()))

		return
	}
}

// Store saves and return new or existing URL.
func (h *ShortenerHandler) Store(w http.ResponseWriter, r *http.Request) {
	request := object.URLRequest{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, sendUnprocessableEntityResponse(err.Error()))

		return
	}

	url, err := h.url.Create(request.OriginalURL)
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
