package feedback

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func (r *Repository) Create(ctx context.Context, feedback *Feedback) error {
	feedback.ID = ""

	_, err := r.DB.InsertOne(ctx, feedback)
	return err
}

func (r *Repository) GetAll(ctx context.Context) ([]*Feedback, error) {
	feedbacks := make([]*Feedback, 0)

	cursor, err := r.DB.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &feedbacks)
	return feedbacks, err
}
