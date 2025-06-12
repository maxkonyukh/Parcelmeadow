package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"parcelmeadow/internal/api/generated/models"
	"parcelmeadow/internal/api/generated/restapi/operations"
)

type Service interface {
	CreateParcel(parcel *models.ParcelV1) (models.ParcelV1, error)
	GetTodayParcels() ([]*models.ParcelV1, error)
	GetTodayRoutes() ([]*models.RouteV1, error)
}

func RegisterInternalHandlers(api *operations.ParcelmeadowAPI, service Service) {
	api.PostV1ParcelsHandler = &PostV1ParcelsHandler{
		service: service,
	}

	api.GetV1ParcelsHandler = &GetV1ParcelsHandler{
		service: service,
	}

	api.GetV1ParcelsRoutesHandler = &GetV1RoutesHandler{
		service: service,
	}
}

type GetV1ParcelsHandler struct {
	service Service
}

func (s *GetV1ParcelsHandler) Handle(_ operations.GetV1ParcelsParams) middleware.Responder {
	parcels, err := s.service.GetTodayParcels()
	if err != nil {
		return operations.NewGetV1ParcelsInternalServerError()
	}

	return operations.NewGetV1ParcelsOK().WithPayload(&models.GetTodayParcelsV1Response{
		Parcels: parcels,
	})
}

type PostV1ParcelsHandler struct {
	service Service
}

func (s *PostV1ParcelsHandler) Handle(params operations.PostV1ParcelsParams) middleware.Responder {
	if params.CreateParcelV1Request == nil {
		return operations.NewPostV1ParcelsBadRequest()
	}

	_, err := s.service.CreateParcel(params.CreateParcelV1Request)
	if err != nil {
		return operations.NewPostV1ParcelsInternalServerError()
	}

	return operations.NewPostV1ParcelsOK()
}

type GetV1RoutesHandler struct {
	service Service
}

func (g *GetV1RoutesHandler) Handle(_ operations.GetV1ParcelsRoutesParams) middleware.Responder {
	routes, err := g.service.GetTodayRoutes()
	if err != nil {
		return operations.NewGetV1ParcelsRoutesInternalServerError()
	}

	return operations.NewGetV1ParcelsRoutesOK().WithPayload(&models.GetTodayRoutesV1Response{
		Routes: routes,
	})
}
