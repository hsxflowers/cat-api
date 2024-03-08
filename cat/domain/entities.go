package domain

import (
	"context"
	"net/url"

	"github.com/hsxflowers/cat-api/exceptions"
)

const MIN_LENGTH_CAT_ID = 3
const MIN_LENGTH_TAG = 3

type CatStorage interface {
	Get(ctx context.Context, tag string) (*Cat, error)
	Create(ctx context.Context, channel *Cat) error
}

type CatDatabase interface {
	Create(ctx context.Context, cat *Cat) error
	Get(ctx context.Context, catId string) (*Cat, error)
}

type Service interface {
	Get(ctx context.Context, tag string) (*CatResponse, error)
	Create(ctx context.Context, cat *CatRequest) error
}

func (u *CatRequest) Validate() error {
	if u.CatId == "" || len(u.CatId) < MIN_LENGTH_CAT_ID {
		return exceptions.ErrCatIdIsRequired
	}

	if u.Tag == "" || len(u.Tag) < MIN_LENGTH_TAG {
		return exceptions.ErrTagIsNotValid
	}

	if u.Url != "" {
		_, err := url.ParseRequestURI(u.Url)
		if err != nil {
			return exceptions.ErrUrlIsNotValid
		}
	}

	return nil
}

func (u *CatRequest) ToCatsDomain() *Cat {
	return &Cat{
		CatId: u.CatId,
		Url:   u.Url,
		Tag:   u.Tag,
	}
}

func (u *Cat) ToCatsResponse() *CatResponse {
	if u != nil {
		return &CatResponse{
			CatId: u.CatId,
			Url:   u.Url,
			Tag:   u.Tag,
		}
	}
	return nil
}

type Cat struct {
	CatId string `json:"cat_id"`
	Url   string `json:"url"`
	Tag   string `json:"tag"`
}

type CatRequest struct {
	CatId string `json:"cat_id"`
	Url   string `json:"url"`
	Tag   string `json:"tag"`
}

type CatResponse struct {
	CatId string `json:"cat_id"`
	Url   string `json:"url"`
	Tag   string `json:"tag"`
}
