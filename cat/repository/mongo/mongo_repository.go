package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

const (
	timeFormat     = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
	collectionName = "cat"
)

func NewMongoRepository(DB *mongo.Database) domain.CatRepository {
	return &mongoRepository{DB, DB.Collection(collectionName)}
}

func (m *mongoRepository) Store(ctx context.Context, cat *domain.Cat) (*domain.Cat, error) {
	var (
		err error
	)

	_, err = m.Collection.InsertOne(ctx, cat)
	if err != nil {
		return cat, err
	}

	return cat, nil
}

func (m *mongoRepository) GetOne(ctx context.Context, id string) (*domain.Cat, error) {
	var (
		cat *domain.Cat
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return cat, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&cat)
	if err != nil {
		return cat, err
	}

	return cat, nil
}

func (m *mongoRepository) GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]domain.Cat, int64, error) {

	var (
		permissions []domain.Cat
		skip        int64
		opts        *options.FindOptions
	)

	skip = (p * rp) - rp
	if setsort != nil {
		opts = options.MergeFindOptions(
			options.Find().SetLimit(rp),
			options.Find().SetSkip(skip),
			options.Find().SetSort(setsort),
		)
	} else {
		opts = options.MergeFindOptions(
			options.Find().SetLimit(rp),
			options.Find().SetSkip(skip),
		)
	}

	cursor, err := m.Collection.Find(
		ctx,
		filter,
		opts,
	)

	if err != nil {
		return nil, 0, err
	}
	if cursor == nil {
		return nil, 0, fmt.Errorf("nil cursor value")
	}
	err = cursor.All(ctx, &permissions)
	if err != nil {
		return nil, 0, err
	}

	count, err := m.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return permissions, 0, err
	}

	return permissions, count, err
}

func (m *mongoRepository) Update(ctx context.Context, cat *domain.Cat, id string) (*domain.Cat, error) {
	var (
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return cat, err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$set": bson.M{
		"name":       cat.Name,
		"species":    cat.Species,
		"legs":       cat.Legs,
		"updated_at": time.Now(),
	}}

	_, err = m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cat, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(cat)
	if err != nil {
		return cat, err
	}
	return cat, nil
}
