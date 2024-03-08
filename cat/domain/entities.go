package domain

import (
	"context"
	url "net/url"

	"github.com/hsxflowers/cat-api/exceptions"
)

const MIN_LENGTH_CAT_ID = 3

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

	if u.Url == "" {
		_, err := url.ParseRequestURI(u.Url)
		if err != nil {
			return exceptions.ErrUrlIsNotValid
		}
	}

	return nil
}

func (u *CatRequest) ToCatsDomain() *Cat {
	return &Cat{
		Url: u.Url,
		Tag: u.Tag,
	}
}

func (u *Cat) ToCatsResponse() *CatResponse {
	if u != nil {
		return &CatResponse{
			Url: u.Url,
			Tag: u.Tag,
		}
	}
	return nil
}

type Cat struct {
	Url string `json:"url"`
	Tag string `json:"tag"`
}

type CatRequest struct {
	Url string `json:"url"`
	Tag string `json:"tag"`
}

type CatResponse struct {
	Url string `json:"url"`
	Tag string `json:"tag"`
}
