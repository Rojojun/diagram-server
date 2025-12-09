package domain

import (
	"testing"
	"time"
)

func TestDiagramDomain_GetCreatedAt(t *testing.T) {
	fixedTime := time.Date(2025, 12, 9, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name   string
		target DiagramDomain
		want   time.Time
	}{
		{
			name:   "초기화된 시간을 반환한다",
			target: DiagramDomain{CreatedAt: fixedTime},
			want:   fixedTime,
		},
		{
			name:   "시간이 초기화 되지 않으면 오류를 반환한다",
			target: DiagramDomain{},
			want:   time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.GetCreatedAt()

			if !got.Equal(tt.want) {
				t.Errorf("GetCreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagramDomain_GetId(t *testing.T) {
	type fields struct {
		ID          string
		Title       string
		Description *string
		Parent      *DiagramDomain
		Child       *DiagramDomain
		CreatedAt   time.Time
		ModifiedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DiagramDomain{
				ID:          tt.fields.ID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Parent:      tt.fields.Parent,
				Child:       tt.fields.Child,
				CreatedAt:   tt.fields.CreatedAt,
				ModifiedAt:  tt.fields.ModifiedAt,
			}
			if got := d.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERDiagramDomain_GetDiagramType(t *testing.T) {
	type fields struct {
		DiagramDomain DiagramDomain
		Tables        []TableDomain
	}
	tests := []struct {
		name   string
		fields fields
		want   DiagramType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ERDiagramDomain{
				DiagramDomain: tt.fields.DiagramDomain,
				Tables:        tt.fields.Tables,
			}
			if got := d.GetDiagramType(); got != tt.want {
				t.Errorf("GetDiagramType() = %v, want %v", got, tt.want)
			}
		})
	}
}
