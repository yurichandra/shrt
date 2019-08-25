package object

import (
	"github.com/go-chi/render"
	"github.com/yurichandra/shrt/model"
)

// CreateURLResponse creating object response for url model.
func CreateURLResponse(url model.URL) render.Renderer {
	return &URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
	}
}

// CreateURLListResponse creating list response of url model.
func CreateURLListResponse(url []model.URL) []render.Renderer {
	urls := make([]render.Renderer, 0)

	for _, item := range url {
		urls = append(urls, CreateURLResponse(item))
	}

	return urls
}
