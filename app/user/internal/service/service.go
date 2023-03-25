package service

import (
	"context"

	"github.com/thesisK19/buildify/app/user/config"
	"github.com/thesisK19/buildify/app/user/internal/store"

	"github.com/sirupsen/logrus"
)

type Service struct {
	config     *config.Config
	log        *logrus.Logger
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
}

func NewService(cfg *config.Config, logger *logrus.Logger, repository store.Repository) *Service {
	return &Service{
		config:     cfg,
		log:        logger,
		repository: repository,
		adapters:   serviceAdapters{},
	}
}

func (s *Service) Close(ctx context.Context) {
	s.repository.Close()
}

func (s *Service) Ping() error {
	err := s.repository.Ping()
	return err
}

func (s *Service) GetRepository() store.Repository {
	return s.repository
}
