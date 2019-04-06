package base

// HttpResponse defines the response data structure
type HttpResponse struct {
	Body       string
	StatusCode int
}

// ControllerInterface defines the methods allowed for each controller
type ControllerInterface interface {
	Get() *HttpResponse
	Post() *HttpResponse
	Put() *HttpResponse
	Delete() *HttpResponse
}

// Controller is the base class that implements the ControllerInterface
// and acts as the template for all inherited sub-controllers
type Controller struct{}

func (c *Controller) Get() *HttpResponse {
	return &HttpResponse{
		Body:       "GET Method Not Allowed",
		StatusCode: 405,
	}
}

func (c *Controller) Post() *HttpResponse {
	return &HttpResponse{
		Body:       "POST Method Not Allowed",
		StatusCode: 405,
	}
}

func (c *Controller) Put() *HttpResponse {
	return &HttpResponse{
		Body:       "PUT Method Not Allowed",
		StatusCode: 405,
	}
}

func (c *Controller) Delete() *HttpResponse {
	return &HttpResponse{
		Body:       "DELETE Method Not Allowed",
		StatusCode: 405,
	}
}
