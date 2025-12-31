package repositories

import (
	"context"

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

func NewDriverRepository(mongoDb *mongo.Database) DriverRepository {

	// Eger index olusturmak gerekirse burada olusturabilir, Ornegin plaka

	return &driverRepository{
		coll: mongoDb.Collection(database.DriversCollection),
	}
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
	_, err := r.coll.UpdateOne(ctx, filter, update)
	return err
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
