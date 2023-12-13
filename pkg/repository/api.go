package postgres

import (
	"context"
	"fmt"
	url_shortener "url-shortener"

	"github.com/jmoiron/sqlx"
)

type Api_postgres struct {
	db *sqlx.DB
}

func New_url_api(db *sqlx.DB) *Api_postgres {
	return &Api_postgres{db: db}
}

func (a *Api_postgres) Create_Short_URL(c context.Context, link *url_shortener.Link) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (base_url, short_url) values ($1, $2)", linksTable)
	a.db.QueryRow(query, link.Base_URL, link.Short_URL)

	return link.Short_URL, nil
}

func (a *Api_postgres) Get_Base_URL(c context.Context, link *url_shortener.Link) (string, error) {
	var base_url string
	query := fmt.Sprintf("SELECT base_url FROM %s WHERE short_url=$1", linksTable)
	err := a.db.Get(&base_url, query, link.Short_URL)

	return base_url, err
}
