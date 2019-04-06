package controllers

import (
	"webmvc/base"
)

type Index struct {
	base.Controller
}

func (index *Index) Get() *base.HttpResponse {
	return &base.HttpResponse{
		Body:       "Hello Webmvc from Go",
		StatusCode: 200,
	}
}
