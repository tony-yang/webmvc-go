package base

// HttpResponse defines the response data structure
type HttpResponse struct {
  Body string
  StatusCode int
}

// ControllerInterface defines the methods allowed for each controller
type ControllerInterface interface {
  Get() *HttpResponse
}

// Controller is the base class that implements the ControllerInterface
// and acts as the template for all inherited sub-controllers
type Controller struct {}

func (c *Controller) Get() *HttpResponse {
  return &HttpResponse{
    Body: "Method Not Allowed",
    StatusCode: 405,
  }
}
