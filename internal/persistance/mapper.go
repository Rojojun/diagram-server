package persistance

import (
	"diagram-server/internal/domain"
	"fmt"
)

func ToModel(d domain.Diagram) *DiagramModel {
	switch v := d.(type) {
	case *domain.ERDiagramDomain:
		return toERDiagramModel(v)
	}
	return nil
}

func toERDiagramModel(d *domain.ERDiagramDomain) *DiagramModel {
	model := &DiagramModel{
		ID:          d.ID,
		Dtype:       string(domain.ERDiagram),
		Title:       d.Title,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		ModifiedAt:  d.ModifiedAt,
		Tables:      toTableModels(d.Tables),
	}

	if d.Parent != nil {
		model.ParentId = &d.Parent.ID
	}

	return model
}

func toTableModels(tables []domain.TableDomain) []TableModel {
	if tables == nil {
		return nil
	}

	result := make([]TableModel, len(tables))

	for i, t := range tables {
		result[i] = TableModel{
			Name:      t.Name,
			Columns:   toColumModels(t.Columns),
			Relations: toRelationModels(t.Relations),
		}
	}

	return result
}

func toColumModels(columns *[]domain.ColumnDomain) []ColumnModel {
	if columns == nil {
		return nil
	}

	result := make([]ColumnModel, len(*columns))

	for i, c := range *columns {
		result[i] = ColumnModel{
			Name:        c.Name,
			Type:        c.Type,
			PK:          c.PK,
			Nullable:    c.Nullable,
			Description: c.Description,
		}
	}
	return result
}

func toRelationModels(relations *[]domain.RelationDomain) []RelationModel {
	if relations == nil {
		return nil
	}

	result := make([]RelationModel, len(*relations))
	for i, r := range *relations {
		result[i] = RelationModel{
			From: r.From,
			To:   r.To,
			Type: string(r.Type),
		}
	}

	return result
}

func (m DiagramModel) ToEntity() (domain.Diagram, error) {
	switch domain.DiagramType(m.Dtype) {
	case domain.ERDiagram:
		return m.toERDiagramDomain(), nil
	}
	return nil, fmt.Errorf("unknown dtype: %s", m.Dtype)
}

func (m *DiagramModel) toERDiagramDomain() *domain.ERDiagramDomain {
	return &domain.ERDiagramDomain{
		DiagramDomain: domain.DiagramDomain{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			CreatedAt:   m.CreatedAt,
			ModifiedAt:  m.ModifiedAt,
		},
		Tables: toTableDomains(m.Tables),
	}
}

func toTableDomains(tables []TableModel) []domain.TableDomain {
	if tables == nil {
		return nil
	}

	result := make([]domain.TableDomain, len(tables))
	for i, t := range tables {
		result[i] = domain.TableDomain{
			Name:      t.Name,
			Columns:   toColumnDomains(t.Columns),
			Relations: toRelationDomains(t.Relations),
		}
	}

	return result
}

func toColumnDomains(columns []ColumnModel) *[]domain.ColumnDomain {
	if columns == nil {
		return nil
	}

	result := make([]domain.ColumnDomain, len(columns))
	for i, c := range columns {
		result[i] = domain.ColumnDomain{
			Name:        c.Name,
			Type:        c.Type,
			PK:          c.PK,
			Nullable:    c.Nullable,
			Description: c.Description,
		}
	}

	return &result
}

func toRelationDomains(relations []RelationModel) *[]domain.RelationDomain {
	if relations == nil {
		return nil
	}

	result := make([]domain.RelationDomain, len(relations))
	for i, r := range relations {
		result[i] = domain.RelationDomain{
			From: r.From,
			To:   r.To,
			Type: domain.RelationType(r.Type),
		}
	}

	return &result
}
