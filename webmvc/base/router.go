package base

import ()

// Routes create the route controller mapping and routes the traffic to
// the corresponding controller for further processing
type Routes struct {
	routes map[string]ControllerInterface
}

func CreateNewRouter() *Routes {
	return &Routes{
		map[string]ControllerInterface{},
	}
}

func (r *Routes) GetController(path string) ControllerInterface {
	return r.routes[path]
}

func (r *Routes) RegisterRoute(path string, controller ControllerInterface) {
	r.routes[path] = controller
}
