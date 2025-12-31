package mappers

import (
	"errors"
	"time"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoToDriver(proto *pb.AddDriverRequest) models.Driver {
	return models.Driver{
		FirstName: proto.FirstName,
		LastName:  proto.LastName,
		Plate:     proto.Plate,
		TaxiType:  proto.TaxiType,
		CarBrand:  proto.CarBrand,
		CarModel:  proto.CarModel,
		Latitude:  proto.Latitude,
		Longitude: proto.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func DriverToProto(driver *models.Driver) *pb.Driver {
	return &pb.Driver{
		Id:        driver.Id.Hex(),
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		Plate:     driver.Plate,
		TaxiType:  driver.TaxiType,
		CarBrand:  driver.CarBrand,
		CarModel:  driver.CarModel,
		Latitude:  driver.Latitude,
		Longitude: driver.Longitude,
		CreatedAt: timestamppb.New(driver.CreatedAt),
		UpdatedAt: timestamppb.New(driver.UpdatedAt),
	}
}

func NearbyDriverToProto(driver models.NearbyDriver) *pb.NearbyDriver {
	return &pb.NearbyDriver{
		FirstName:  driver.FirstName,
		LastName:   driver.LastName,
		Plate:      driver.Plate,
		DistanceKm: driver.DistanceKm,
	}
}

func ProtoToUpdateDoc(req *pb.UpdateDriverRequest) (bson.M, error) {
	update := bson.M{}
	if req.FirstName != nil {
		update["firstName"] = req.GetFirstName()
	}
	if req.LastName != nil {
		update["lastName"] = req.GetLastName()
	}
	if req.Plate != nil {
		update["plate"] = req.GetPlate()
	}
	if req.TaxiType != nil {
		update["taxiType"] = req.GetTaxiType()
	}
	if req.CarModel != nil {
		update["carModel"] = req.GetCarModel()
	}
	if req.CarBrand != nil {
		update["carBrand"] = req.GetCarBrand()
	}
	if req.Latitude != nil {
		update["lat"] = req.GetLatitude()
	}
	if req.Longitude != nil {
		update["lon"] = req.GetLongitude()
	}

	if len(update) == 0 {
		return nil, errors.New("no fields to update")
	}

	update["updatedAt"] = time.Now()

	return update, nil
}

func DriversToProto(drivers []models.Driver) []*pb.Driver {
	result := make([]*pb.Driver, len(drivers))
	for i, driver := range drivers {
		result[i] = DriverToProto(&driver)
	}
	return result
}

func NearbyDriversToProto(drivers []models.NearbyDriver) []*pb.NearbyDriver {
	result := make([]*pb.NearbyDriver, len(drivers))
	for i, driver := range drivers {
		result[i] = NearbyDriverToProto(driver)
	}
	return result
}

func PaginationMetadataToProto(metadata *models.PaginationMetadata) *pb.PaginationMetadata {
	if metadata == nil {
		return nil
	}

	return &pb.PaginationMetadata{
		TotalCount: metadata.TotalCount,
		TotalPages: metadata.TotalPages,
		Page:       metadata.Page,
		PageSize:   metadata.PageSize,
	}
}
