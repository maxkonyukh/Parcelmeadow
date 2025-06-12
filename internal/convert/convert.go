package convert

import (
	"parcelmeadow/internal/api/generated/models"
	"parcelmeadow/internal/database"
	"time"
)

func ParcelApiToDb(parcel models.ParcelV1) database.Parcel {
	return database.Parcel{
		Id:        parcel.ID,
		Weight:    parcel.Weight,
		Status:    database.ParcelStatus(parcel.Status),
		PostCode:  parcel.PostCode,
		Address:   parcel.Address,
		CreatedAt: time.Now().UTC(),
	}
}

func ParcelDbToApi(parcel database.Parcel) models.ParcelV1 {
	return models.ParcelV1{
		ID:       parcel.Id,
		Weight:   parcel.Weight,
		Status:   string(parcel.Status),
		PostCode: parcel.PostCode,
		Address:  parcel.Address,
	}
}

func StopDbToApi(stop database.Stop) models.StopV1 {
	var parcels []*models.ParcelV1
	for _, parcel := range stop.Parcels {
		apiParcel := ParcelDbToApi(parcel)
		parcels = append(parcels, &apiParcel)
	}

	return models.StopV1{
		ID:      stop.Id,
		Address: stop.Address,
		Parcels: parcels,
	}
}

func RouteDbToApi(route database.Route) models.RouteV1 {
	var stops []*models.StopV1
	for _, stop := range route.Stops {
		apiStop := StopDbToApi(stop)
		stops = append(stops, &apiStop)
	}

	return models.RouteV1{
		ID:    route.Id,
		Stops: stops,
	}
}
