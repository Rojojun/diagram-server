package handler

type CreateDiagramDTO struct {
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Tables      []TableDTO `json:"tables,omitempty"`
}

type TableDTO struct {
	Name          string        `json:"name"`
	OriginalQuery *string       `json:"original_query,omitempty"`
	Columns       []ColumnDTO   `json:"columns,omitempty"`
	Relations     []RelationDTO `json:"relations,omitempty"`
}

type ColumnDTO struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	PK          bool    `json:"pk"`
	Nullable    bool    `json:"nullable"`
	Description *string `json:"description,omitempty"`
}

type RelationDTO struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}

type DiagramResponse struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Tables      []TableDTO `json:"tables,omitempty"`
	Owner       string     `json:"owner"`
	CreatedAt   string     `json:"createdAt"`
	ModifiedAt  string     `json:"modifiedAt"`
}
