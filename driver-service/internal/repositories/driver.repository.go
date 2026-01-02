package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"github.com/AkifhanIlgaz/taxihub/driver-service/pkg/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DriverRepository interface {
	Insert(ctx context.Context, driver models.Driver) (bson.ObjectID, error)
	UpdateById(ctx context.Context, id bson.ObjectID, updates bson.M) error
	FindDrivers(ctx context.Context, filters bson.M, opts *options.FindOptionsBuilder) ([]models.Driver, int64, error)
}

type driverRepository struct {
	coll *mongo.Collection
}

func NewDriverRepository(mongoDb *mongo.Database) (DriverRepository, error) {
	coll := mongoDb.Collection(database.DriversCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"plate": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, fmt.Errorf("create drivers plate index: %w", err)
	}

	return &driverRepository{
		coll: coll,
	}, nil
}

func (r *driverRepository) Insert(ctx context.Context, driver models.Driver) (bson.ObjectID, error) {
	result, err := r.coll.InsertOne(ctx, driver)
	if err != nil {
		return bson.ObjectID{}, err
	}
	return result.InsertedID.(bson.ObjectID), nil
}

func (r *driverRepository) UpdateById(ctx context.Context, id bson.ObjectID, updates bson.M) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updates}
	res, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *driverRepository) FindDrivers(ctx context.Context, filters bson.M, opts *options.FindOptionsBuilder) ([]models.Driver, int64, error) {
	cursor, err := r.coll.Find(ctx, filters, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var drivers []models.Driver
	if err := cursor.All(ctx, &drivers); err != nil {
		return nil, 0, err
	}

	count, err := r.coll.CountDocuments(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	return drivers, count, nil
}
