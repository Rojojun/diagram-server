package persistance

import "diagram-server/internal/domain"

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
