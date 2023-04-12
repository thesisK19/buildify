package service

import (
	"context"

	"github.com/thesisK19/buildify/app/gen-code/config"
	"github.com/thesisK19/buildify/app/gen-code/internal/store"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
}

func NewService(cfg *config.Config, repository store.Repository) *Service {
	return &Service{
		config:     cfg,
		repository: repository,
		adapters:   serviceAdapters{},
	}
}

func (s *Service) Close(ctx context.Context) {
	s.repository.Close()
}

func (s *Service) Ping() error {
	err := s.repository.Ping() // TODO:
	return err
}

func (s *Service) GetRepository() store.Repository {
	return s.repository
}
