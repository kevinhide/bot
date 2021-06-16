package routes

import (
	"bot/handlers"
	m "bot/middlewares"
	"bot/redis"
	"bot/shared"
	"net/http"

	"github.com/gorilla/mux"
)

//Route : ""
type Route struct {
	Handler *handlers.Handler
	Shared  *shared.Shared
	Redis   *redis.RedisCli
}

//GetRoute : ""
func GetRoute(handler *handlers.Handler, s *shared.Shared, Redis *redis.RedisCli) *Route {
	return &Route{handler, s, Redis}
}

// Adapt :
func Adapt(h http.Handler, adapters ...m.Adapter) http.Handler {
	if len(adapters) == 0 {
		return h
	}
	return adapters[0](Adapt(h, adapters[1:cap(adapters)]...))
}

//Category : ""
func (rout *Route) Category(r *mux.Router) {
	r.Handle("/category/getallcategory/{isParent}", Adapt(http.HandlerFunc(rout.Handler.Category))).Methods("GET")
}
