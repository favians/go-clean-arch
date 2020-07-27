package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

const (
	timeFormat     = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
	collectionName = "user"
)

func NewMongoRepository(DB mongo.Database) domain.UserRepository {
	return &mongoRepository{DB, DB.Collection(collectionName)}
}

func (m *mongoRepository) InsertOne(ctx context.Context, user *domain.User) (*domain.User, error) {
	var (
		err error
	)

	_, err = m.Collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *mongoRepository) FindOne(ctx context.Context, id string) (*domain.User, error) {
	var (
		user domain.User
		err  error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &user, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&user)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (m *mongoRepository) GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]domain.User, int64, error) {

	var (
		user []domain.User
		skip int64
		opts *options.FindOptions
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
	err = cursor.All(ctx, &user)
	if err != nil {
		return nil, 0, err
	}

	count, err := m.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return user, 0, err
	}

	return user, count, err
}

func (m *mongoRepository) UpdateOne(ctx context.Context, user *domain.User, id string) (*domain.User, error) {
	var (
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$set": bson.M{
		"name":       user.Name,
		"username":   user.Username,
		"password":   user.Password,
		"updated_at": time.Now(),
	}}

	_, err = m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return user, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *mongoRepository) GetByCredential(ctx context.Context, username string, password string) (*domain.User, error) {
	var (
		user domain.User
		err  error
	)

	credential := bson.M{
		"username": username,
		"password": password,
	}

	err = m.Collection.FindOne(ctx, credential).Decode(&user)
	if err != nil {
		return &user, err
	}

	return &user, nil
}
