package controllers

import (
	"webmvc/base"
)

type Hello struct {
	base.Controller
}

func (h *Hello) Get() *base.HttpResponse {
	return &base.HttpResponse{
		Body:       "Hello World!",
		StatusCode: 200,
	}
}

func (h *Hello) Post() *base.HttpResponse {
	return &base.HttpResponse{
		Body:       "Post to Hello World!",
		StatusCode: 200,
	}
}
