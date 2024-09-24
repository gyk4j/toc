package web

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

const (
	contentTypeHeader = "Content-Type"
	textHTML          = "text/html"
	textPlain         = "text/plain"
)

// PassthroughBuilder returns the handler, aka the builder identity function
func PassthroughBuilder(handler http.Handler) http.Handler { return handler }

func NewTocWeb() *TocWeb {
	return &TocWeb{
		handlers: make(map[string]map[string]http.Handler),
	}
}

type TocWeb struct {
	handlers   map[string]map[string]http.Handler
	Middleware func(middleware.Builder) http.Handler
}

func (o *TocWeb) initHandlerCache() {
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
}

func (o *TocWeb) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Println(r.URL)

		rw.Header().Set(contentTypeHeader, textHTML)
		rw.WriteHeader(http.StatusNotFound)
	})
}

func (o *TocWeb) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
