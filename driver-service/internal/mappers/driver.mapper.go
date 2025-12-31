package mappers

import (
	"encoding/json"
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
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		Plate:     driver.Plate,
	}
}

func ToUpdateDoc(req *pb.UpdateDriverRequest) (bson.M, error) {

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	update := bson.M{}
	if err := bson.Unmarshal(data, &update); err != nil {
		return nil, err
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
