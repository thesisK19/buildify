package service

import (
	"context"
	"fmt"

	dynamicDataAdapter "github.com/thesisK19/buildify/app/dynamic_data/pkg/adapter"
	genCodeAdapter "github.com/thesisK19/buildify/app/gen_code/pkg/adapter"
	"github.com/thesisK19/buildify/app/user/config"
	"github.com/thesisK19/buildify/app/user/internal/store"
)

type Service struct {
	config     *config.Config
	repository store.Repository
	adapters   serviceAdapters
}

type serviceAdapters struct {
	genCode     genCodeAdapter.GenCodeAdapter
	dynamicData dynamicDataAdapter.DynamicDataAdapter
}

func NewService(cfg *config.Config, repository store.Repository) *Service {
	dynamicDataAdapter, err := dynamicDataAdapter.NewDynamicDataAdapter(cfg.DynamicDataGRPCAddr)
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
			genCode:     genCodeAdapter,
			dynamicData: dynamicDataAdapter,
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
