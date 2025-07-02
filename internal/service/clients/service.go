package clients

import (
	"context"

	"github.com/BigDwarf/sahtian/internal/model"
	"github.com/BigDwarf/sahtian/internal/repository"
)

type Service struct {
	repo *repository.ClientRepository
}

func NewService(repo *repository.ClientRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateClient(ctx context.Context, client *model.Client) error {
	return s.repo.Create(ctx, client)
}

func (s *Service) GetAllClients(ctx context.Context) ([]*model.Client, error) {
	return s.repo.GetAll(ctx)
}
