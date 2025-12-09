package persistance

import (
	"time"
)

type DiagramModel struct {
	ID          string    `bson:"_id,omitempty"`
	Dtype       string    `bson:"dtype"`
	Title       string    `bson:"title"`
	Description *string   `bson:"description,omitempty"`
	ParentId    *string   `bson:"parent_id,omitempty"`
	CreatedAt   time.Time `bson:"createdAt"`
	ModifiedAt  time.Time `bson:"modifiedAt"`

	// Dtype == ERDiagram
	Tables []TableModel `bson:"tables,omitempty"`
}

type TableModel struct {
	Name      string          `bson:"name"`
	Columns   []ColumnModel   `bson:"columns"`
	Relations []RelationModel `bson:"relations"`
}

type ColumnModel struct {
	Name        string  `bson:"name"`
	Type        string  `bson:"type"`
	PK          bool    `bson:"pk"`
	Nullable    bool    `bson:"nullable"`
	Description *string `bson:"description,omitempty"`
}

type RelationModel struct {
	From string `bson:"from"`
	To   string `bson:"to"`
	Type string `bson:"type"`
}
