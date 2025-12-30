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
		TaxiType:  "sari",
		CarBrand:  "Toyota",
		CarModel:  "Corolla",
		Latitude:  41.0102,
		Longitude: 28.9788,
	},
	{
		FirstName: "Mehmet",
		LastName:  "Kaya",
		Plate:     "34DEF456",
		TaxiType:  "korsan",
		CarBrand:  "Renault",
		CarModel:  "Clio",
		Latitude:  41.0125,
		Longitude: 28.9812,
	},
	{
		FirstName: "Ayse",
		LastName:  "Demir",
		Plate:     "34GHI789",
		TaxiType:  "uber",
		CarBrand:  "BMW",
		CarModel:  "3 Series",
		Latitude:  41.0158,
		Longitude: 28.9899,
	},
	{
		FirstName: "Fatma",
		LastName:  "Celik",
		Plate:     "34JKL012",
		TaxiType:  "tag",
		CarBrand:  "Fiat",
		CarModel:  "Egea",
		Latitude:  41.0067,
		Longitude: 28.9725,
	},
	{
		FirstName: "Can",
		LastName:  "Aydin",
		Plate:     "34MNO345",
		TaxiType:  "uber",
		CarBrand:  "Mercedes",
		CarModel:  "Vito",
		Latitude:  41.0214,
		Longitude: 28.9951,
	},
	{
		FirstName: "Elif",
		LastName:  "Arslan",
		Plate:     "34PQR678",
		TaxiType:  "sari",
		CarBrand:  "Hyundai",
		CarModel:  "i20",
		Latitude:  41.0183,
		Longitude: 28.9864,
	},
	{
		FirstName: "Hakan",
		LastName:  "Kurt",
		Plate:     "34STU901",
		TaxiType:  "korsan",
		CarBrand:  "Ford",
		CarModel:  "Focus",
		Latitude:  41.0039,
		Longitude: 28.9822,
	},
	{
		FirstName: "Zeynep",
		LastName:  "Sahin",
		Plate:     "34VWX234",
		TaxiType:  "uber",
		CarBrand:  "Audi",
		CarModel:  "A3",
		Latitude:  41.0785,
		Longitude: 29.0062,
	},
	{
		FirstName: "Mert",
		LastName:  "Yildiz",
		Plate:     "34YZA567",
		TaxiType:  "tag",
		CarBrand:  "Volkswagen",
		CarModel:  "Golf",
		Latitude:  41.0096,
		Longitude: 28.9654,
	},
	{
		FirstName: "Selin",
		LastName:  "Ozturk",
		Plate:     "34BCD890",
		TaxiType:  "sari",
		CarBrand:  "Opel",
		CarModel:  "Astra",
		Latitude:  41.0270,
		Longitude: 28.9890,
	},
	{
		FirstName: "Burak",
		LastName:  "Kilic",
		Plate:     "34EFG112",
		TaxiType:  "korsan",
		CarBrand:  "Peugeot",
		CarModel:  "301",
		Latitude:  40.9482,
		Longitude: 29.0725,
	},
	{
		FirstName: "Deniz",
		LastName:  "Aksoy",
		Plate:     "34HIJ334",
		TaxiType:  "uber",
		CarBrand:  "Honda",
		CarModel:  "Civic",
		Latitude:  41.0169,
		Longitude: 28.9995,
	},
	{
		FirstName: "Gizem",
		LastName:  "Tas",
		Plate:     "34KLM556",
		TaxiType:  "tag",
		CarBrand:  "Citroen",
		CarModel:  "C4",
		Latitude:  41.0111,
		Longitude: 28.9922,
	},
	{
		FirstName: "Emre",
		LastName:  "Polat",
		Plate:     "34NOP778",
		TaxiType:  "sari",
		CarBrand:  "Skoda",
		CarModel:  "Octavia",
		Latitude:  41.0839,
		Longitude: 28.9054,
	},
	{
		FirstName: "Derya",
		LastName:  "Kara",
		Plate:     "34QRS990",
		TaxiType:  "uber",
		CarBrand:  "Kia",
		CarModel:  "Ceed",
		Latitude:  41.0054,
		Longitude: 28.9990,
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
