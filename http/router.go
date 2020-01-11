package router

import "net/http"

//Router is the interface to be implemented by the HTTP routers
type Router interface {
	GET(uri string, f func(http.ResponseWriter, *http.Request))
	POST(uri string, f func(http.ResponseWriter, *http.Request))
	SERVE(port string)
}
