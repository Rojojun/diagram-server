package domain

import "time"

type Diagram interface {
	Type() DiagramType
	ID() string
	CreatedAt() time.Time
	Owner() string
}

type DiagramType string

const (
	TypeERD       DiagramType = "erdiagram"
	TypeFlowChart DiagramType = "flowchart"
)

type BaseDiagram struct {
	id          string
	title       string
	description *string
	diagramType DiagramType
	owner       string
	createdAt   time.Time
	modifiedAt  time.Time
}

func (b BaseDiagram) ID() string            { return b.id }
func (b BaseDiagram) Title() string         { return b.title }
func (b BaseDiagram) Description() *string  { return b.description }
func (b BaseDiagram) Type() DiagramType     { return b.diagramType }
func (b BaseDiagram) Owner() string         { return b.owner }
func (b BaseDiagram) CreatedAt() time.Time  { return b.createdAt }
func (b BaseDiagram) ModifiedAt() time.Time { return b.modifiedAt }

func NewBaseDiagram(title string, description *string, dtype DiagramType, owner string) BaseDiagram {
	now := time.Now()
	return BaseDiagram{
		title:       title,
		description: description,
		diagramType: dtype,
		owner:       owner,
		createdAt:   now,
		modifiedAt:  now,
	}
}

func RestoreBaseDiagram(id, title string, description *string, dtype DiagramType, owner string, createdAt, modifiedAt time.Time) BaseDiagram {
	return BaseDiagram{
		id:          id,
		title:       title,
		description: description,
		diagramType: dtype,
		owner:       owner,
		createdAt:   createdAt,
		modifiedAt:  modifiedAt,
	}
}

func (b *BaseDiagram) SetID(id string) {
	b.id = id
}

func (b *BaseDiagram) UpdateTitle(title string) {
	b.title = title
	b.modifiedAt = time.Now()
}

func (b *BaseDiagram) UpdateDescription(description *string) {
	b.description = description
	b.modifiedAt = time.Now()
}

func (e *ERDiagram) UpdateTables(tables []Table) {
	e.Tables = tables
	e.modifiedAt = time.Now()
}

func (e *ERDiagram) Update(title *string, description *string, tables []Table) {
	if title != nil {
		e.title = *title
	}
	if description != nil {
		e.description = description
	}
	if tables != nil {
		e.Tables = tables
	}
	e.modifiedAt = time.Now()
}

type ERDiagram struct {
	BaseDiagram
	Tables []Table
}

func NewERDiagram(title string, description *string, owner string, tables []Table) *ERDiagram {
	return &ERDiagram{
		BaseDiagram: NewBaseDiagram(title, description, TypeERD, owner),
		Tables:      tables,
	}
}

type Table struct {
	Name          string
	OriginalQuery *string
	Columns       *[]Column
	Relations     *[]Relation
}

type Column struct {
	Name        string
	Type        string
	PK          bool
	Nullable    bool
	Description *string
}

type Relation struct {
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
