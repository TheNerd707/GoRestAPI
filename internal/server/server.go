package server

import (
	"net/http"
)


type Route struct {
	method string
	route string
	function http.HandlerFunc
}


type Application struct {
	routes []Route
	port string
}



func New() *Application {
	return &Application{}
}

func (app *Application) Route(method string, route string, function http.HandlerFunc) {
	var temp Route
	temp.method = method
	temp.route = route
	temp.function = function
	app.routes = append(app.routes, temp)
}

type MuxHandler func(*http.ServeMux)

func (app *Application) Listen(port string, funcion MuxHandler) error {
	app.port = port
	mux := http.NewServeMux()
	funcion(mux)
	for i := range app.routes {
		route := app.routes[i]
		mux.HandleFunc(route.method + " " + route.route, route.function)
	}
	return http.ListenAndServe(app.port, mux)
}