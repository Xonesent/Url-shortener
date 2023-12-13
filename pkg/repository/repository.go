package postgres

import (
	"context"
	url_shortener "url-shortener"

	"github.com/jmoiron/sqlx"
)

type Url_api interface {
	Create_Short_URL(context.Context, *url_shortener.Link) (string, error)
	Get_Base_URL(context.Context, *url_shortener.Link) (string, error)
}

type Repository struct {
	Url_api
}

func New_Repository(db *sqlx.DB) *Repository {
	return &Repository{Url_api: New_url_api(db)}
}
