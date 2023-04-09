package main

import (
	"context"
	"flag"
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

	"github.com/sirupsen/logrus"
)

type Service struct {
	config     *config.Config
	log        *logrus.Logger
	repository store.Repository
	adapters   serviceAdapters
	router     *mux.Router
}

type serviceAdapters struct {
}

func NewService(cfg *config.Config, logger *logrus.Logger, repository store.Repository, router *mux.Router) *Service {
	return &Service{
		config:     cfg,
		log:        logger,
		repository: repository,
		adapters:   serviceAdapters{},
		router:     router,
	}
}

func (s *Service) setRouter() {
	s.Post("/upload/image", handler.UploadImageHandler)
	s.Get("/", handler.HelloWorld)
}

// Run will start the http server on host that you pass in. host:<ip:port>
func (s *Service) Serve() {

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
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	log.Printf("Server is listening on %s\n", address)

	// Block until we receive our signal.
	<-c

	log.Println("Signal: ", c)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Println("shutting down...")
	srv.Shutdown(ctx)
	s.Close()
	os.Exit(0)
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
