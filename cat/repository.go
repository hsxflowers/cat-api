package cat

import (
	"context"
	_ "log"

	"github.com/hsxflowers/cat-api/cat/domain"
	"github.com/hsxflowers/cat-api/exceptions"
	"github.com/labstack/gommon/log"
)

type CatRepository struct {
	database domain.CatStorage
}

func NewCatRepository(database domain.CatDatabase) *CatRepository {
	return &CatRepository{
		database: database,
	}
}

func (repo *CatRepository) Get(ctx context.Context, catId string) (*domain.Cat, error) {
	response, err := repo.database.Get(ctx, catId)
	if err != nil {
		if err.Error() == "not found" {
			log.Error("cat_repo: cat_id not found", err)
			return nil, exceptions.New(exceptions.ErrCatNotFound, err)
		}
		log.Error("cat_repo: error on get cat", err)
		return nil, exceptions.New(exceptions.ErrInternalServer, err)
	}
	return response, err
}

func (repo *CatRepository) Create(ctx context.Context, cat *domain.Cat) error {
	err := repo.database.Create(ctx, cat)
	if err != nil {
		log.Error("cat_repo: error on create cat in the database", err)
		return exceptions.New(exceptions.ErrInternalServer, err)
	}

	return nil
}
