package cat

import (
	"context"

	"github.com/hsxflowers/cat-api/cat/domain"
	"github.com/labstack/gommon/log"
)

type Service struct {
	repo domain.CatStorage
}

func NewCatService(repo domain.CatStorage) domain.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(ctx context.Context, catId string) (*domain.CatResponse, error) {
	log.Debug("[GET - get] - Request processed by service ")

	cat, err := s.repo.Get(ctx, catId)
	if err != nil {
		log.Errorf("get_cat_service: error on get cat, %s", err)
		return nil, err
	}

	catResponse := s.toCatResponse(cat)

	return catResponse, nil
}

func (s *Service) Create(ctx context.Context, req *domain.CatRequest) error {
	log.Debug("[POST - Create] - Request processed by service ")

	cat := req.ToCatsDomain()

	err := s.repo.Create(ctx, cat)
	if err != nil {
		log.Errorf("create_cat_service: error on create cat, %s", err)
		return err
	}

	return nil
}

func (s *Service) toCatResponse(cat *domain.Cat) *domain.CatResponse {
	return &domain.CatResponse{
		CatId: cat.CatId,
		Url:   cat.Url,
		Tag:   cat.Tag,
	}
}
