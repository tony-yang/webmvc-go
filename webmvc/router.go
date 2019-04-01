package webmvc

import (
	"webmvc/base"
)

// Routes create the route controller mapping and routes the traffic to
// the corresponding controller for further processing
type Routes struct {
	routes map[string]base.ControllerInterface
}

func (r *Routes) GetController(path string) base.ControllerInterface {
	return r.routes[path]
}
