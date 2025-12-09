package domain

import "time"

type Diagram interface {
	GetDiagramType() DiagramType
	GetId() string
	GetCreatedAt() time.Time
}

type DiagramDomain struct {
	ID          string
	Title       string
	Description *string
	Parent      *DiagramDomain
	Child       *DiagramDomain
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type DiagramType string

const (
	ERDiagram DiagramType = "erdiagram"
	FlowChart DiagramType = "flowchart"
)

func (d DiagramDomain) GetId() string {
	return d.ID
}

func (d DiagramDomain) GetCreatedAt() time.Time {
	return d.CreatedAt
}

type ERDiagramDomain struct {
	DiagramDomain
	Tables []TableDomain
}

type TableDomain struct {
	Name      string
	Columns   *[]ColumnDomain
	Relations *[]RelationDomain
}

type ColumnDomain struct {
	Name        string
	Type        string
	PK          bool
	Nullable    bool
	Description *string
}

type RelationDomain struct {
	From string
	To   string
	Type RelationType
}

type RelationType string

const (
	OneToOne   RelationType = "one_to_one"
	OneToMany  RelationType = "one_to_many"
	ManyToMany RelationType = "many_to_many"
	ManyToOne  RelationType = "many_to_one"
)

func (d ERDiagramDomain) GetDiagramType() DiagramType {
	return ERDiagram
}
