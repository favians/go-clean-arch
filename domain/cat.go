package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cat struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Species   string             `bson:"species" json:"species" validate:"required"`
	Legs      int                `bson:"legs" json:"legs" validate:"required"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type CatRepository interface {
	InsertOne(ctx context.Context, u *Cat) (*Cat, error)
	FindOne(ctx context.Context, id string) (*Cat, error)
	GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]Cat, int64, error)
	UpdateOne(ctx context.Context, cat *Cat, id string) (*Cat, error)
}

type CatUsecase interface {
	InsertOne(ctx context.Context, u *Cat) (*Cat, error)
	FindOne(ctx context.Context, id string) (*Cat, error)
	GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]Cat, int64, error)
	UpdateOne(ctx context.Context, cat *Cat, id string) (*Cat, error)
}
