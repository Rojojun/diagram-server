package service

import (
	"context"
	"diagram-server/internal/domain"
	"diagram-server/internal/persistance"
	"errors"
	"time"
)

type DiagramService interface {
	Create(ctx context.Context, req CreateDiagramRequest) (*domain.ERDiagramDomain, error)
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
	Description *string
	Tables      []domain.TableDomain
}

type UpdateDiagramRequest struct {
	Title       *string
	Description *string
	Tables      []domain.TableDomain
}

func (s *diagramService) Create(ctx context.Context, req CreateDiagramRequest) (*domain.ERDiagramDomain, error) {
	now := time.Now()
	diagram := &domain.ERDiagramDomain{
		DiagramDomain: domain.DiagramDomain{
			Title:       req.Title,
			Description: req.Description,
			CreatedAt:   now,
			ModifiedAt:  now,
		},
		Tables: req.Tables,
	}

	id, err := s.repo.Save(ctx, diagram)
	if err != nil {
		return nil, err
	}

	diagram.ID = id
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

	erd, ok := diagram.(*domain.ERDiagramDomain)
	if !ok {
		return errors.New("invalid diagram type")
	}

	if req.Title != nil {
		erd.Title = *req.Title
	}
	if req.Description != nil {
		erd.Description = req.Description
	}
	if req.Tables != nil {
		erd.Tables = req.Tables
	}
	erd.ModifiedAt = time.Now()

	return s.repo.Update(ctx, erd)
}

func (s *diagramService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
