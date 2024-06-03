package serverengine

import (
	"net/http"
	"web-engine-go/middleware"
	"web-engine-go/utils"
)
type Route struct {
	WithLogger bool
	Handler http.Handler
}

type Server struct {
	Routes map[string] *Route
}

func New(routes map[string]*Route) *Server {
	return &Server{Routes: routes}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	route, ok := s.Routes[head]
	if !ok {
	utils.Respond(w, r, http.StatusBadRequest, "root route not found")
		return
	}
	next := route.Handler
	if (route.WithLogger) {
		next = middleware.Logger(next)
	}
	next.ServeHTTP(w, r)

}

