package webmvc

import (
	"webmvc/base"
)

// Routes create the route controller mapping and routes the traffic to
// the corresponding controller for further processing
type Routes struct {
	routes map[string]base.ControllerInterface
}

func CreateNewRouter() *Routes {
	return &Routes{
		map[string]base.ControllerInterface{},
	}
}

func (r *Routes) GetController(path string) base.ControllerInterface {
	return r.routes[path]
}

func (r *Routes) RegisterRoute(path string, controller base.ControllerInterface) {
	r.routes[path] = controller
}
