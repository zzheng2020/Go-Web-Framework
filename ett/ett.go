package ett

import (
	"log"
	"net/http"
	"strings"
)

// define handler to handle request
type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix       string
		middleweares []HandlerFunc
		parent       *RouterGroup
		engine       *Engine
	}
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middleweares = append(group.middleweares, middlewares...)
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}

	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}

	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// handle received requests
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleweares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
