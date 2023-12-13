package service

import (
	"context"
	"database/sql"
	url_shortener "url-shortener"
	"url-shortener/pkg/repository"
)

type Api_service struct {
	repos postgres.Url_api
}

func New_url_api(repos postgres.Url_api) *Api_service {
	return &Api_service{repos: repos}
}

func (a *Api_service) Create_Short_URL(c context.Context, link *url_shortener.Link) (string, error) {
	link.Short_URL = Generate_Short_URL()
	Short_URL, err := a.repos.Create_Short_URL(c, link)
	if err != nil {
		return "", err
	}
	return Short_URL, nil
}

func (a *Api_service) Get_Base_URL(c context.Context, link *url_shortener.Link) (string, error) {
	Base_URL, err := a.repos.Get_Base_URL(c, link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return Base_URL, nil
}
