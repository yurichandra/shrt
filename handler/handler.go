package handler

import "github.com/yurichandra/shrt/service"

// NewShortenerHandler return new handler of shortener.
func NewShortenerHandler(srv service.ShortenerServiceContract) *ShortenerHandler {
	return &ShortenerHandler{
		url: srv,
	}
}

// NewAuthHandler return new handler of auth.
func NewAuthHandler(srv service.AuthServiceContract) *AuthHandler {
	return &AuthHandler{
		auth: srv,
	}
}
