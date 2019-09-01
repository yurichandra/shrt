package object

import (
	"errors"
	"net/http"
)

// URLRequest represent object request of URL model.
type URLRequest struct {
	OriginalURL string `json:"original_url"`
}

// URLResponse represent object response of URL model.
type URLResponse struct {
	OriginalURL string `json:"original_url"`
	Keys        string `json:"keys"`
}

// Bind doing validation for request object.
func (req *URLRequest) Bind(r *http.Request) error {
	if req.OriginalURL == "" {
		return errors.New("Original URL can't be empty")
	}

	return nil
}

// Render renders url model.
func (res *URLResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
