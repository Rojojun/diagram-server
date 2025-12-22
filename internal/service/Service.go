package service

import (
	"context"
	"diagram-server/internal/domain"
	"diagram-server/internal/persistance"
	"errors"
)

type DiagramService interface {
	Create(ctx context.Context, req CreateDiagramRequest) (*domain.ERDiagram, error)
	GetByID(ctx context.Context, id string) (domain.Diagram, error)
	GetAllByType(ctx context.Context, dtype domain.DiagramType) ([]domain.Diagram, error)
	Update(ctx context.Context, id string, req UpdateDiagramRequest) error
	Delete(ctx context.Context, id string) error
}

type diagramService struct {
	repo persistance.DiagramRepository
}

func NewDiagramService(repo persistance.DiagramRepository) DiagramService {
	return &diagramService{repo: repo}
}

type CreateDiagramRequest struct {
	Title       string
	Owner       string
	Description *string
	Tables      []domain.Table
}

type UpdateDiagramRequest struct {
	Title       *string
	Description *string
	Tables      []domain.Table
}

func (s *diagramService) Create(ctx context.Context, req CreateDiagramRequest) (*domain.ERDiagram, error) {
	diagram := domain.NewERDiagram(
		req.Title,
		req.Description,
		req.Owner,
		req.Tables,
	)

	id, err := s.repo.Save(ctx, diagram)
	if err != nil {
		return nil, err
	}

	diagram.SetID(id)
	return diagram, nil
}

func (s *diagramService) GetByID(ctx context.Context, id string) (domain.Diagram, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *diagramService) GetAllByType(ctx context.Context, dtype domain.DiagramType) ([]domain.Diagram, error) {
	return s.repo.FindByType(ctx, dtype)
}

func (s *diagramService) Update(ctx context.Context, id string, req UpdateDiagramRequest) error {
	diagram, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	erd, ok := diagram.(*domain.ERDiagram)
	if !ok {
		return errors.New("invalid diagram type")
	}

	erd.Update(req.Title, req.Description, req.Tables)

	return s.repo.Update(ctx, erd)
}

func (s *diagramService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
