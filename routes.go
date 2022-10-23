package GoMWF

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (c *GoMWF) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if c.DEBUG {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	return mux
}
