package controllers

import (
  "webmvc/base"
)

type Hello struct {}

func (h *Hello) Get() *base.HttpResponse {
  return &base.HttpResponse{
    Body: "Hello World!",
    StatusCode: 200,
  }
}
