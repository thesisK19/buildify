package service

import (
	"context"
	"fmt"

	"github.com/thesisK19/buildify/app/dynamic_data/config"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/store"
	genCodeAdapter "github.com/thesisK19/buildify/app/gen_code/pkg/adapter"
	userAdapter "github.com/thesisK19/buildify/app/user/pkg/adapter"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
	genCode genCodeAdapter.GenCodeAdapter
	user    userAdapter.UserAdapter
}

func NewService(cfg *config.Config, repository store.Repository) *Service {
	userAdapter, err := userAdapter.NewUserAdapter(cfg.UserGRPCAddr)
	if err != nil {
		fmt.Printf("Can't connect to user service, err=%v", err)
	}
	genCodeAdapter, err := genCodeAdapter.NewGenCodeAdapter(cfg.GenCodeGRPCAddr)
	if err != nil {
		fmt.Printf("Can't connect to user service, err=%v", err)
	}

	return &Service{
		config:     cfg,
		repository: repository,
		adapters: serviceAdapters{
			user:    userAdapter,
			genCode: genCodeAdapter,
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
