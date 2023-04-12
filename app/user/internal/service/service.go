package service

import (
	"context"
	"fmt"

	"github.com/thesisK19/buildify/app/user/config"
	"github.com/thesisK19/buildify/app/user/internal/adapter"
	"github.com/thesisK19/buildify/app/user/internal/store"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
	genCode adapter.GenCodeClient
}

func NewService(cfg *config.Config, repository store.Repository) *Service {
	genCode, err := adapter.NewGenCodeClient(cfg.GenCodeHost)
	if err != nil {
		// should not return err, we will re-connect later
		fmt.Println("Init NewGenCodeClient fail...")
	}

	return &Service{
		config:     cfg,
		repository: repository,
		adapters: serviceAdapters{
			genCode: genCode,
		},
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
