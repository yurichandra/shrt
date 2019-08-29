package handler

import "github.com/yurichandra/shrt/service"

// NewURLHandler return new handler of URL.
func NewURLHandler(srv service.URLServiceContract) *URLHandler {
	return &URLHandler{
		url: srv,
	}
}

// NewAuthHandler return new handler of auth.
func NewAuthHandler(srv service.AuthServiceContract) *AuthHandler {
	return &AuthHandler{
		auth: srv,
	}
}
