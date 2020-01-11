package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

//NewChiRouter is a Chi HTTP router constructor
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(http.ResponseWriter, *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(http.ResponseWriter, *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP Server running on port: %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
