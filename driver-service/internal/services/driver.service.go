package services

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DriverService struct {
	repo repositories.DriverRepository
}

func NewDriverService(repo repositories.DriverRepository) *DriverService {
	return &DriverService{
		repo: repo,
	}
}

func (s *DriverService) AddDriver(driverToAdd models.AddDriverRequest) (bson.ObjectID, error) {
	driver := driverToAdd.ToDriver()
	return s.repo.Insert(context.Background(), driver)
}

func (s *DriverService) UpdateDriver(id string, driverToUpdate models.UpdateDriverRequest) error {
	driverId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("update driver: %w", err)
	}

	updates, err := driverToUpdate.ToUpdateDoc()
	if err != nil {
		return fmt.Errorf("update driver: %w", err)
	}

	return s.repo.UpdateById(context.Background(), driverId, updates)
}

func (s *DriverService) GetDrivers(req models.ListDriversRequest) ([]models.ListDriversResponse, error) {
	opts := options.Find()
	var metadata *models.Metadata

	if req.Page != 0 && req.PageSize != 0 {
		skip := int64((req.Page - 1) * req.PageSize)
		opts = options.Find().SetSkip(skip).SetLimit(int64(req.PageSize))
		metadata = &models.Metadata{
			Page:     req.Page,
			PageSize: req.PageSize,
		}
	}

	drivers, count, err := s.repo.FindDrivers(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("get drivers: %w", err)
	}

	if metadata != nil && metadata.PageSize > 0 {
		metadata.TotalCount = int(count)
		metadata.TotalPages = int((count + int64(req.PageSize) - 1) / int64(req.PageSize))
	}

	return []models.ListDriversResponse{
		{
			Drivers: drivers,
			Meta:    metadata,
		},
	}, nil
}

func (s *DriverService) GetNearbyDrivers(req models.ListNearbyDriversRequest) ([]models.NearbyDrivers, error) {
	filter := bson.M{
		"taxiType": req.TaxiType,
	}

	drivers, _, err := s.repo.FindDrivers(context.Background(), filter, nil)
	if err != nil {
		return nil, fmt.Errorf("get drivers: %w", err)
	}

	nearbyDrivers := []models.NearbyDrivers{}

	for _, driver := range drivers {
		distance := calcDistance(req.Latitude, req.Longitude, driver.Latitude, driver.Longitude)
		if distance < 6 {
			nearbyDrivers = append(nearbyDrivers, models.NearbyDrivers{
				FirstName:  driver.FirstName,
				LastName:   driver.LastName,
				Plate:      driver.Plate,
				DistanceKm: distance,
			})
		}
	}

	sort.Slice(nearbyDrivers, func(i, j int) bool {
		return nearbyDrivers[i].DistanceKm < nearbyDrivers[j].DistanceKm
	})

	return nearbyDrivers, nil
}

func calcDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Earth's radius in kilometers => 3440,1 NM (Nautical Mile )

	lat1Rad, lon1Rad := degreeToRadian(lat1), degreeToRadian(lon1)
	lat2Rad, lon2Rad := degreeToRadian(lat2), degreeToRadian(lon2)

	diffLat := lat2Rad - lat1Rad
	diffLon := lon2Rad - lon1Rad

	a := math.Sin(diffLat/2)*math.Sin(diffLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(diffLon/2)*math.Sin(diffLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func degreeToRadian(degree float64) float64 {
	return degree * math.Pi / 180
}
