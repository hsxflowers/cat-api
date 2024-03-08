package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" 
	"github.com/hsxflowers/cat-api/cat/domain"
	"github.com/hsxflowers/cat-api/exceptions"
	"github.com/labstack/gommon/log"
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) Get(ctx context.Context, tag string) (*domain.Cat, error) {
	var cat domain.Cat
    var catId, url, tagResult string

    query := "SELECT cat_id, url, tag FROM cats WHERE tag = $1 ORDER BY RANDOM() LIMIT 1"
    row := s.db.QueryRowContext(ctx, query, tag)

    err := row.Scan(&catId, &url, &tagResult)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, exceptions.New(exceptions.ErrCatNotFound, err)
        }
        log.Error("Error fetching cat from database: ", err)
        return nil, exceptions.New(exceptions.ErrInternalServer, err)
    }

    cat.CatId = catId
    cat.Url = url
    cat.Tag = tagResult

    return &cat, nil
}

func (s *SQLStore) Create(ctx context.Context, cat *domain.Cat) error {
    _, err := s.db.ExecContext(ctx, "INSERT INTO cats (url, tag) VALUES ($1, $2)", cat.Url, cat.Tag)
    if err != nil {
        log.Error("Error creating cat in database: ", err)
        return exceptions.New(exceptions.ErrInternalServer, err)
    }
    return nil
}
