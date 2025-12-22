package persistance

import (
	"diagram-server/internal/domain"
	"fmt"
)

func ToModel(d domain.Diagram) *DiagramModel {
	switch v := d.(type) {
	case *domain.ERDiagram:
		return toERDiagramModel(v)
	}
	return nil
}

func toERDiagramModel(d *domain.ERDiagram) *DiagramModel {
	return &DiagramModel{
		ID:          d.ID(),
		Dtype:       string(d.Type()),
		Title:       d.Title(),
		Description: d.Description(),
		Owner:       d.Owner(),
		CreatedAt:   d.CreatedAt(),
		ModifiedAt:  d.ModifiedAt(),
		Tables:      toTableModels(d.Tables),
	}
}

func toTableModels(tables []domain.Table) []TableModel {
	if tables == nil {
		return nil
	}

	result := make([]TableModel, len(tables))
	for i, t := range tables {
		result[i] = TableModel{
			Name:          t.Name,
			OriginalQuery: t.OriginalQuery,
			Columns:       toColumModels(t.Columns),
			Relations:     toRelationModels(t.Relations),
		}
	}

	return result
}

func toColumModels(columns *[]domain.Column) []ColumnModel {
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

func toRelationModels(relations *[]domain.Relation) []RelationModel {
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
	case domain.TypeERD:
		return m.toERDiagram(), nil
	case domain.TypeFlowChart:
		// TODO
		return nil, fmt.Errorf("flowchart not implemented")
	}
	return nil, fmt.Errorf("unknown dtype: %s", m.Dtype)
}

func (m *DiagramModel) toERDiagram() *domain.ERDiagram {
	base := domain.RestoreBaseDiagram(
		m.ID,
		m.Title,
		m.Description,
		domain.TypeERD,
		m.Owner,
		m.CreatedAt,
		m.ModifiedAt,
	)

	return &domain.ERDiagram{
		BaseDiagram: base,
		Tables:      toTableDomains(m.Tables),
	}
}

func toTableDomains(tables []TableModel) []domain.Table {
	if tables == nil {
		return nil
	}

	result := make([]domain.Table, len(tables))
	for i, t := range tables {
		result[i] = domain.Table{
			Name:          t.Name,
			OriginalQuery: t.OriginalQuery,
			Columns:       toColumnDomains(t.Columns),
			Relations:     toRelationDomains(t.Relations),
		}
	}
	return result
}

func toColumnDomains(columns []ColumnModel) *[]domain.Column {
	if columns == nil {
		return nil
	}

	result := make([]domain.Column, len(columns))
	for i, c := range columns {
		result[i] = domain.Column{
			Name:        c.Name,
			Type:        c.Type,
			PK:          c.PK,
			Nullable:    c.Nullable,
			Description: c.Description,
		}
	}
	return &result
}

func toRelationDomains(relations []RelationModel) *[]domain.Relation {
	if relations == nil {
		return nil
	}

	result := make([]domain.Relation, len(relations))
	for i, r := range relations {
		result[i] = domain.Relation{
			From: r.From,
			To:   r.To,
			Type: domain.RelationType(r.Type),
		}
	}
	return &result
}
