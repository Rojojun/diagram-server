package handler

import (
	"diagram-server/internal/domain"
	"diagram-server/internal/service"
	"encoding/json"
	"net/http"
	"time"
)

type DiagramHandler struct {
	svc service.DiagramService
}

func NewDiagramHandler(svc service.DiagramService) *DiagramHandler {
	return &DiagramHandler{svc: svc}
}

func (h *DiagramHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateDiagramDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := service.CreateDiagramRequest{
		Title:       dto.Title,
		Description: dto.Description,
		Tables:      toTableDomains(dto.Tables),
	}

	diagram, err := h.svc.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toResponse(diagram))
}

func (h *DiagramHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	diagram, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toResponse(diagram))
}

func (h *DiagramHandler) GetAllByType(w http.ResponseWriter, r *http.Request) {
	dtype := r.PathValue("type")

	diagrams, err := h.svc.GetAllByType(r.Context(), domain.DiagramType(dtype))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := make([]DiagramResponse, len(diagrams))
	for i, d := range diagrams {
		responses[i] = toResponse(d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func (h *DiagramHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.svc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ========== Mappers ==========

func toResponse(d domain.Diagram) DiagramResponse {
	resp := DiagramResponse{
		ID:        d.GetId(),
		Type:      string(d.GetDiagramType()),
		CreatedAt: d.GetCreatedAt().Format(time.RFC3339),
	}

	if erd, ok := d.(*domain.ERDiagramDomain); ok {
		resp.Title = erd.Title
		resp.Description = erd.Description
		resp.ModifiedAt = erd.ModifiedAt.Format(time.RFC3339)
		resp.Tables = toTableDTOs(erd.Tables)
	}

	return resp
}

func toTableDomains(dtos []TableDTO) []domain.TableDomain {
	if dtos == nil {
		return nil
	}

	result := make([]domain.TableDomain, len(dtos))
	for i, dto := range dtos {
		result[i] = domain.TableDomain{
			Name:      dto.Name,
			Columns:   toColumnDomains(dto.Columns),
			Relations: toRelationDomains(dto.Relations),
		}
	}
	return result
}

func toColumnDomains(dtos []ColumnDTO) *[]domain.ColumnDomain {
	if dtos == nil {
		return nil
	}

	result := make([]domain.ColumnDomain, len(dtos))
	for i, dto := range dtos {
		result[i] = domain.ColumnDomain{
			Name:        dto.Name,
			Type:        dto.Type,
			PK:          dto.PK,
			Nullable:    dto.Nullable,
			Description: dto.Description,
		}
	}
	return &result
}

func toRelationDomains(dtos []RelationDTO) *[]domain.RelationDomain {
	if dtos == nil {
		return nil
	}

	result := make([]domain.RelationDomain, len(dtos))
	for i, dto := range dtos {
		result[i] = domain.RelationDomain{
			From: dto.From,
			To:   dto.To,
			Type: domain.RelationType(dto.Type),
		}
	}
	return &result
}

func toTableDTOs(tables []domain.TableDomain) []TableDTO {
	if tables == nil {
		return nil
	}

	result := make([]TableDTO, len(tables))
	for i, t := range tables {
		result[i] = TableDTO{
			Name:      t.Name,
			Columns:   toColumnDTOs(t.Columns),
			Relations: toRelationDTOs(t.Relations),
		}
	}
	return result
}

func toColumnDTOs(columns *[]domain.ColumnDomain) []ColumnDTO {
	if columns == nil {
		return nil
	}

	result := make([]ColumnDTO, len(*columns))
	for i, c := range *columns {
		result[i] = ColumnDTO{
			Name:        c.Name,
			Type:        c.Type,
			PK:          c.PK,
			Nullable:    c.Nullable,
			Description: c.Description,
		}
	}
	return result
}

func toRelationDTOs(relations *[]domain.RelationDomain) []RelationDTO {
	if relations == nil {
		return nil
	}

	result := make([]RelationDTO, len(*relations))
	for i, r := range *relations {
		result[i] = RelationDTO{
			From: r.From,
			To:   r.To,
			Type: string(r.Type),
		}
	}
	return result
}
