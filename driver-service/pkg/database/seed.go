package database

import (
	"context"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var seeds = []models.Driver{
	{
		FirstName: "Ahmet",
		LastName:  "Yilmaz",
		Plate:     "34ABC123",
		TaxiType:  "standard",
		CarBrand:  "Toyota",
		CarModel:  "Corolla",
		Latitude:  41.0082,
		Longitude: 28.9784,
	},
	{
		FirstName: "Mehmet",
		LastName:  "Kaya",
		Plate:     "34DEF456",
		TaxiType:  "standard",
		CarBrand:  "Renault",
		CarModel:  "Clio",
		Latitude:  41.0369,
		Longitude: 28.9846,
	},
	{
		FirstName: "Ayse",
		LastName:  "Demir",
		Plate:     "34GHI789",
		TaxiType:  "premium",
		CarBrand:  "BMW",
		CarModel:  "3 Series",
		Latitude:  41.0430,
		Longitude: 29.0100,
	},
	{
		FirstName: "Fatma",
		LastName:  "Celik",
		Plate:     "34JKL012",
		TaxiType:  "standard",
		CarBrand:  "Fiat",
		CarModel:  "Egea",
		Latitude:  40.9900,
		Longitude: 29.0300,
	},
	{
		FirstName: "Can",
		LastName:  "Aydin",
		Plate:     "34MNO345",
		TaxiType:  "premium",
		CarBrand:  "Mercedes",
		CarModel:  "C200",
		Latitude:  41.0150,
		Longitude: 29.0550,
	},
}

func SeedDrivers(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := db.Collection(DriversCollection)
	now := time.Now().UTC()

	opts := options.UpdateOne().SetUpsert(true)
	for _, seed := range seeds {
		filter := bson.M{"plate": seed.Plate}
		update := bson.M{
			"$set": bson.M{
				"firstName": seed.FirstName,
				"lastName":  seed.LastName,
				"plate":     seed.Plate,
				"taxiType":  seed.TaxiType,
				"carBrand":  seed.CarBrand,
				"carModel":  seed.CarModel,
				"lat":       seed.Latitude,
				"lon":       seed.Longitude,
				"updatedAt": now,
			},
			"$setOnInsert": bson.M{
				"createdAt": now,
			},
		}

		if _, err := coll.UpdateOne(ctx, filter, update, opts); err != nil {
			log.Fatalf("upsert seed driver %s: %v", seed.Plate, err)
		}
	}

	return nil
}
