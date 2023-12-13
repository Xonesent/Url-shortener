package service

import (
	"context"
	url_shortener "url-shortener"
	"url-shortener/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock_service.go

type Url_api interface {
	Create_Short_URL(context.Context, *url_shortener.Link) (string, error)
	Get_Base_URL(context.Context, *url_shortener.Link) (string, error)
}

type Service struct {
	Url_api
}

func New_Service(repos *postgres.Repository) *Service {
	return &Service{Url_api: New_url_api(repos.Url_api)}
}
