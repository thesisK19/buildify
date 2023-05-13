package service

import (
	"context"
	"fmt"

	dynamicDataAdapter "github.com/thesisK19/buildify/app/dynamic_data/pkg/adapter"
	userAdapter "github.com/thesisK19/buildify/app/user/pkg/adapter"

	"github.com/thesisK19/buildify/app/gen_code/config"
	"github.com/thesisK19/buildify/app/gen_code/internal/store"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
	user        userAdapter.UserAdapter
	dynamicData dynamicDataAdapter.DynamicDataAdapter
}

func NewService(cfg *config.Config, repository store.Repository) *Service {
	userAdapter, err := userAdapter.NewUserAdapter(cfg.UserGRPCAddr)
	if err != nil {
		fmt.Printf("Can't connect to user service, err=%v", err)
	}
	dynamicDataAdapter, err := dynamicDataAdapter.NewDynamicDataAdapter(cfg.DynamicDataGRPCAddr)
	if err != nil {
		fmt.Printf("Can't connect to user service, err=%v", err)
	}
	return &Service{
		config:     cfg,
		repository: repository,
		adapters: serviceAdapters{
			user:        userAdapter,
			dynamicData: dynamicDataAdapter,
		},
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
