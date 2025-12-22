package domain

import (
	"testing"
	"time"
)

func TestBaseDiagram_ID(t *testing.T) {
	tests := []struct {
		name   string
		target BaseDiagram
		want   string
	}{
		{
			name:   "ID를 반환한다",
			target: BaseDiagram{id: "test-id-123"},
			want:   "test-id-123",
		},
		{
			name:   "ID가 비어있으면 빈 문자열을 반환한다",
			target: BaseDiagram{},
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.ID()
			if got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDiagram_Title(t *testing.T) {
	tests := []struct {
		name   string
		target BaseDiagram
		want   string
	}{
		{
			name:   "Title을 반환한다",
			target: BaseDiagram{title: "My Diagram"},
			want:   "My Diagram",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Title()
			if got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDiagram_Description(t *testing.T) {
	desc := "테스트 설명"

	tests := []struct {
		name   string
		target BaseDiagram
		want   *string
	}{
		{
			name:   "Description을 반환한다",
			target: BaseDiagram{description: &desc},
			want:   &desc,
		},
		{
			name:   "Description이 nil이면 nil을 반환한다",
			target: BaseDiagram{},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Description()
			if tt.want == nil {
				if got != nil {
					t.Errorf("Description() = %v, want nil", got)
				}
			} else if *got != *tt.want {
				t.Errorf("Description() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestBaseDiagram_Type(t *testing.T) {
	tests := []struct {
		name   string
		target BaseDiagram
		want   DiagramType
	}{
		{
			name:   "ERD 타입을 반환한다",
			target: BaseDiagram{diagramType: TypeERD},
			want:   TypeERD,
		},
		{
			name:   "FlowChart 타입을 반환한다",
			target: BaseDiagram{diagramType: TypeFlowChart},
			want:   TypeFlowChart,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Type()
			if got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDiagram_Owner(t *testing.T) {
	tests := []struct {
		name   string
		target BaseDiagram
		want   string
	}{
		{
			name:   "Owner를 반환한다",
			target: BaseDiagram{owner: "user-123"},
			want:   "user-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Owner()
			if got != tt.want {
				t.Errorf("Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDiagram_CreatedAt(t *testing.T) {
	fixedTime := time.Date(2025, 12, 9, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name   string
		target BaseDiagram
		want   time.Time
	}{
		{
			name:   "CreatedAt을 반환한다",
			target: BaseDiagram{createdAt: fixedTime},
			want:   fixedTime,
		},
		{
			name:   "초기화되지 않으면 zero time을 반환한다",
			target: BaseDiagram{},
			want:   time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.CreatedAt()
			if !got.Equal(tt.want) {
				t.Errorf("CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDiagram_ModifiedAt(t *testing.T) {
	fixedTime := time.Date(2025, 12, 10, 12, 30, 0, 0, time.UTC)

	tests := []struct {
		name   string
		target BaseDiagram
		want   time.Time
	}{
		{
			name:   "ModifiedAt을 반환한다",
			target: BaseDiagram{modifiedAt: fixedTime},
			want:   fixedTime,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.ModifiedAt()
			if !got.Equal(tt.want) {
				t.Errorf("ModifiedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBaseDiagram(t *testing.T) {
	desc := "테스트 설명"

	tests := []struct {
		name        string
		title       string
		description *string
		dtype       DiagramType
		owner       string
	}{
		{
			name:        "새 BaseDiagram을 생성한다",
			title:       "My Diagram",
			description: &desc,
			dtype:       TypeERD,
			owner:       "user-123",
		},
		{
			name:        "description이 nil이어도 생성된다",
			title:       "No Desc",
			description: nil,
			dtype:       TypeFlowChart,
			owner:       "user-456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := time.Now()
			got := NewBaseDiagram(tt.title, tt.description, tt.dtype, tt.owner)
			after := time.Now()

			if got.Title() != tt.title {
				t.Errorf("Title() = %v, want %v", got.Title(), tt.title)
			}
			if got.Type() != tt.dtype {
				t.Errorf("Type() = %v, want %v", got.Type(), tt.dtype)
			}
			if got.Owner() != tt.owner {
				t.Errorf("Owner() = %v, want %v", got.Owner(), tt.owner)
			}
			if got.ID() != "" {
				t.Errorf("ID() should be empty for new diagram, got %v", got.ID())
			}
			if got.CreatedAt().Before(before) || got.CreatedAt().After(after) {
				t.Errorf("CreatedAt() = %v, should be between %v and %v", got.CreatedAt(), before, after)
			}
			if !got.CreatedAt().Equal(got.ModifiedAt()) {
				t.Errorf("CreatedAt and ModifiedAt should be equal for new diagram")
			}
		})
	}
}

func TestRestoreBaseDiagram(t *testing.T) {
	desc := "복원된 설명"
	createdAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	modifiedAt := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	got := RestoreBaseDiagram("id-123", "Restored", &desc, TypeERD, "owner-1", createdAt, modifiedAt)

	if got.ID() != "id-123" {
		t.Errorf("ID() = %v, want id-123", got.ID())
	}
	if got.Title() != "Restored" {
		t.Errorf("Title() = %v, want Restored", got.Title())
	}
	if !got.CreatedAt().Equal(createdAt) {
		t.Errorf("CreatedAt() = %v, want %v", got.CreatedAt(), createdAt)
	}
	if !got.ModifiedAt().Equal(modifiedAt) {
		t.Errorf("ModifiedAt() = %v, want %v", got.ModifiedAt(), modifiedAt)
	}
}

func TestERDiagram_Type(t *testing.T) {
	tests := []struct {
		name   string
		target ERDiagram
		want   DiagramType
	}{
		{
			name: "ERD 타입을 반환한다",
			target: ERDiagram{
				BaseDiagram: BaseDiagram{diagramType: TypeERD},
				Tables:      []Table{},
			},
			want: TypeERD,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Type()
			if got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewERDiagram(t *testing.T) {
	desc := "ERD 설명"
	columns := []Column{{Name: "id", Type: "bigint", PK: true}}
	tables := []Table{{Name: "users", Columns: &columns}}

	got := NewERDiagram("User Schema", &desc, "owner-123", tables)

	if got.Type() != TypeERD {
		t.Errorf("Type() = %v, want %v", got.Type(), TypeERD)
	}
	if got.Title() != "User Schema" {
		t.Errorf("Title() = %v, want User Schema", got.Title())
	}
	if got.Owner() != "owner-123" {
		t.Errorf("Owner() = %v, want owner-123", got.Owner())
	}
	if len(got.Tables) != 1 {
		t.Errorf("Tables length = %v, want 1", len(got.Tables))
	}
	if got.Tables[0].Name != "users" {
		t.Errorf("Tables[0].Name = %v, want users", got.Tables[0].Name)
	}
}

func TestERDiagram_ImplementsDiagram(t *testing.T) {
	var _ Diagram = (*ERDiagram)(nil)
}
