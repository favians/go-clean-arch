package mongo_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/mongo/mocks"
	repo "github.com/bxcodec/go-clean-arch/user/repository/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "user"
)

func TestInsertOne(t *testing.T) {

	// Define variables for interfaces
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	// Set interfaces implementation to mocked structures
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	mockEmptyUser := &domain.User{}
	mockUserID := primitive.NewObjectID()

	t.Run("Sucsess InsertOne", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserID, nil).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.InsertOne(context.Background(), mockUser)

		assert.Equal(t, mockUser.Name, u.Name)
		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("Failed InsertOne", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.InsertOne(context.Background(), mockEmptyUser)

		assert.Equal(t, "", u.Name)
		assert.NotNil(t, err)

		collectionHelper.AssertExpectations(t)
	})
}

func TestFindOne(t *testing.T) {

	// Define variables for interfaces
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var SingleResult *mocks.SingleResult

	// Set interfaces implementation to mocked structures
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	SingleResult = &mocks.SingleResult{}

	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	mockUserID := primitive.NewObjectID()
	mockUserIDhex := mockUserID.Hex()
	filter := bson.M{"_id": mockUserID}

	t.Run("Sucsess FindOne", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, filter).Return(SingleResult, nil).Once()
		SingleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.Name = "vian"
		}).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.FindOne(context.Background(), mockUserIDhex)
		log.Println(u)

		assert.Equal(t, mockUser.Name, u.Name)
		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("Failed FindOne", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, filter).Return(SingleResult, errors.New("Unexpected")).Once()
		SingleResult.On("Decode", mock.Anything).Return(errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.FindOne(context.Background(), mockUserIDhex)

		assert.Equal(t, "", u.Name)
		assert.NotNil(t, err)

		collectionHelper.AssertExpectations(t)
	})
}

func TestUpdateOne(t *testing.T) {

	// Define variables for interfaces
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var SingleResult *mocks.SingleResult

	// Set interfaces implementation to mocked structures
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	SingleResult = &mocks.SingleResult{}

	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	mockEmptyUser := &domain.User{}

	mockUserID := primitive.NewObjectID()
	mockUserIDhex := mockUserID.Hex()
	filter := bson.M{"_id": mockUserID}

	t.Run("Sucsess UpdateOne", func(t *testing.T) {
		collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		collectionHelper.On("FindOne", mock.Anything, filter).Return(SingleResult, nil).Once()
		SingleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.Name = "vian"
		}).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.UpdateOne(context.Background(), mockUser, mockUserIDhex)
		log.Println(u)

		assert.Equal(t, mockUser.Name, u.Name)
		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("Failed UpdateOne", func(t *testing.T) {
		collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.UpdateOne(context.Background(), mockEmptyUser, mockUserIDhex)

		assert.Equal(t, "", u.Name)
		assert.NotNil(t, err)

		collectionHelper.AssertExpectations(t)
	})
}

func TestGetAllWithPage(t *testing.T) {

	// Define variables for interfaces
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var cursor *mocks.Cursor

	// Set interfaces implementation to mocked structures
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	cursor = &mocks.Cursor{}

	mockListUser := []domain.User{
		domain.User{
			ID:        primitive.NewObjectID(),
			Name:      "vian",
			Username:  "favian",
			Password:  "password",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		domain.User{
			ID:        primitive.NewObjectID(),
			Name:      "test",
			Username:  "test",
			Password:  "test",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	p := int64(1)
	rp := int64(25)
	filter := mock.Anything

	t.Run("Sucsess GetAllWithPage", func(t *testing.T) {

		collectionHelper.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cursor, nil)
		cursor.On("All", mock.Anything, mock.AnythingOfType("*[]domain.User")).
			Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]domain.User)
			*arg = mockListUser
		})
		collectionHelper.On("CountDocuments", mock.Anything, mock.Anything, mock.Anything).Return(int64(25), nil)
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		_, count, err := r.GetAllWithPage(context.Background(), rp, p, filter, nil)

		assert.Equal(t, count, int64(25))
		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})
}

func TestGetByCredential(t *testing.T) {

	// Define variables for interfaces
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var SingleResult *mocks.SingleResult

	// Set interfaces implementation to mocked structures
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	SingleResult = &mocks.SingleResult{}

	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	username := "vian"
	password := "password"

	t.Run("Sucsess GetByCredential", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(SingleResult, nil).Once()
		SingleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.Name = "vian"
		}).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.GetByCredential(context.Background(), username, password)
		log.Println(u)

		assert.Equal(t, mockUser.Name, u.Name)
		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("Failed FindOne", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(SingleResult, errors.New("Unexpected")).Once()
		SingleResult.On("Decode", mock.Anything).Return(errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		// Create new database with mocked Database interface
		r := repo.NewMongoRepository(databaseHelper)

		u, err := r.GetByCredential(context.Background(), "", "")

		assert.Equal(t, "", u.Name)
		assert.NotNil(t, err)

		collectionHelper.AssertExpectations(t)
	})
}
