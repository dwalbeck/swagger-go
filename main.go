package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	midware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"swagger-go/api"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil
	apiRoutes := api.NewDemoApi()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(midware.OapiRequestValidator(swagger))

	// ------ start UI method 1 ---------------------------------
	fs := http.FileServer(http.Dir("swaggerui"))
	r.Handle("/swaggerui/*", http.StripPrefix("/swaggerui/", fs))

	dir, _ := os.Getwd()
	fileDir := http.Dir(filepath.Join(dir, "static"))
	FileServer(r, "/swaggerui", fileDir)

	// ------ start UI method 2 ---------------------------------
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "swaggerui"))
	FileServer(r, "/api-docs", filesDir)

	// ------ enc UI --------------------------------------------

	api.HandlerFromMux(apiRoutes, r)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("localhost", "8080"),
	}

	log.Println("starting server")
	log.Fatal(s.ListenAndServe())
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
