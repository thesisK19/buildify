package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/thesisK19/buildify/app/file-management/config"
	"github.com/thesisK19/buildify/app/file-management/internal/handler"
	"github.com/thesisK19/buildify/app/file-management/internal/store"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
	router     *mux.Router
}

type serviceAdapters struct {
}

func NewService(cfg *config.Config, repository store.Repository, router *mux.Router) *Service {
	return &Service{
		config:     cfg,
		repository: repository,
		adapters:   serviceAdapters{},
		router:     router,
	}
}

func (s *Service) setRouter() {
	s.Post("/upload/image", handler.UploadImageHandler)
	// s.Get("/", handler.HelloWorld)
}

// Run will start the http server on host that you pass in. host:<ip:port>
func (s *Service) Serve() error {
	errch := make(chan error)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	address := fmt.Sprintf("%s:%d", s.config.HTTP.Host, s.config.HTTP.Port)
	srv := &http.Server{
		Addr: address,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      cors.AllowAll().Handler(s.router), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Error starting http server, ", err)
			errch <- err
		}
		log.Printf("Server is listening on %s\n", address)
	}()

	// shutdown
	for {
		select {
		case <-stop:
			log.Println("Shutting down server")
			//nolint:gomnd
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
			return nil
		case err := <-errch:
			return err
		}
	}
}

func (s *Service) Close() {
	s.repository.Close()
}

func (s *Service) Ping() error {
	err := s.repository.Ping()
	return err
}

func (s *Service) GetRepository() store.Repository {
	return s.repository
}

// Get wraps the router for GET method
func (s *Service) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (s *Service) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (s *Service) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (s *Service) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("DELETE")
}

// UseMiddleware will add global middleware in router
func (s *Service) UseMiddleware(middleware mux.MiddlewareFunc) {
	s.router.Use(middleware)
}

// func corsHandler() *cors.Cors {
// 	c := cors.New(cors.Options{
// 		AllowedOrigins: []string{"*"},
// 		AllowedMethods: []string{
// 			http.MethodHead,
// 			http.MethodGet,
// 			http.MethodPost,
// 			http.MethodPut,
// 			http.MethodPatch,
// 			http.MethodDelete,
// 		},
// 		AllowedHeaders:   []string{"*"},
// 		AllowCredentials: false,
// 	})

// 	return c
// }
