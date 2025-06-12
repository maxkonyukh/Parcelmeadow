package services

import (
	"fmt"
	"parcelmeadow/internal/api/generated/models"
	"parcelmeadow/internal/convert"
	"parcelmeadow/internal/database"
)

type Storage interface {
	SaveParcel(parcel database.Parcel) (database.Parcel, error)
	GetTodayParcels() ([]database.Parcel, error)
	GetTodayRoutes() ([]database.Route, error)
}

type ParcelmeadowService struct {
	storage Storage
}

func NewParcelmeadowService(storage Storage) (*ParcelmeadowService, error) {
	if storage == nil {
		return nil, fmt.Errorf("storage is nil")
	}

	return &ParcelmeadowService{
		storage: storage,
	}, nil
}

func (p *ParcelmeadowService) CreateParcel(parcel *models.ParcelV1) (models.ParcelV1, error) {
	if parcel == nil {
		return models.ParcelV1{}, fmt.Errorf("parcel is nil")
	}

	dbParcel := convert.ParcelApiToDb(*parcel)

	res, err := p.storage.SaveParcel(dbParcel)
	if err != nil {
		return models.ParcelV1{}, fmt.Errorf("failed to save parcel: %w", err)
	}

	newParcel := convert.ParcelDbToApi(res)

	return newParcel, nil
}

func (p *ParcelmeadowService) GetTodayParcels() ([]*models.ParcelV1, error) {
	parcels, err := p.storage.GetTodayParcels()
	if err != nil {
		return nil, fmt.Errorf("failed to get today's parcels: %w", err)
	}

	var apiParcels []*models.ParcelV1
	for _, parcel := range parcels {
		apiParcel := convert.ParcelDbToApi(parcel)
		apiParcels = append(apiParcels, &apiParcel)
	}

	return apiParcels, nil
}

func (p *ParcelmeadowService) GetTodayRoutes() ([]*models.RouteV1, error) {
	routes, err := p.storage.GetTodayRoutes()
	if err != nil {
		return nil, fmt.Errorf("failed to get today's routes: %w", err)
	}

	var apiRoutes []*models.RouteV1
	for _, route := range routes {
		apiRoute := convert.RouteDbToApi(route)
		apiRoutes = append(apiRoutes, &apiRoute)
	}

	return apiRoutes, nil
}
