package persistance

import (
	"context"
	"diagram-server/internal/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotFound = errors.New("diagram not found")

type DiagramRepository interface {
	Save(ctx context.Context, d domain.Diagram) (string, error)
	FindByID(ctx context.Context, id string) (domain.Diagram, error)
	FindByType(ctx context.Context, dtype domain.DiagramType) ([]domain.Diagram, error)
	Update(ctx context.Context, d domain.Diagram) error
	Delete(ctx context.Context, id string) error
}

type mongoDiagramRepository struct {
	coll *mongo.Collection
}

func NewDiagramRepository(db *mongo.Database) DiagramRepository {
	return &mongoDiagramRepository{
		coll: db.Collection("diagrams"),
	}
}

func (r *mongoDiagramRepository) Save(ctx context.Context, d domain.Diagram) (string, error) {
	model := ToModel(d)

	if model.ID == "" {
		model.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.coll.InsertOne(ctx, model)
	return model.ID, err
}

func (r *mongoDiagramRepository) FindByID(ctx context.Context, id string) (domain.Diagram, error) {
	var model DiagramModel
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&model)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return model.ToEntity()
}

func (r *mongoDiagramRepository) FindByType(ctx context.Context, dtype domain.DiagramType) ([]domain.Diagram, error) {
	cursor, err := r.coll.Find(ctx, bson.M{"dtype": string(dtype)})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return r.decodeCursor(ctx, cursor)
}

func (r *mongoDiagramRepository) Update(ctx context.Context, d domain.Diagram) error {
	model := ToModel(d)
	result, err := r.coll.ReplaceOne(ctx, bson.M{"_id": model.ID}, model)

	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *mongoDiagramRepository) Delete(ctx context.Context, id string) error {
	result, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *mongoDiagramRepository) decodeCursor(ctx context.Context, cursor *mongo.Cursor) ([]domain.Diagram, error) {
	var results []domain.Diagram

	for cursor.Next(ctx) {
		var model DiagramModel
		if err := cursor.Decode(&model); err != nil {
			return nil, err
		}
	}

	return results, nil
}
