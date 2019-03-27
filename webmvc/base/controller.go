package base

type HttpResponse struct {
  Body string
  StatusCode int
}

type ControllerInterface interface {
  Get() *HttpResponse
}

type Controller struct {}

func (c *Controller) Get() *HttpResponse {
  return &HttpResponse{
    Body: "Method not allowed",
    StatusCode: 405,
  }
}
