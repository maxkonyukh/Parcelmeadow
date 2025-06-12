package handlers

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"parcelmeadow/internal/api/generated/models"
	"parcelmeadow/internal/api/generated/restapi/operations"
	"testing"

	mock_handlers "parcelmeadow/internal/api/handlers/mock/generated"
)

func setupHandlers(ctrl *gomock.Controller) (*PostV1ParcelsHandler, *GetV1ParcelsHandler, *GetV1RoutesHandler, mock_handlers.MockService) {
	mockStorage := mock_handlers.NewMockService(ctrl)

	postV1ParcelsHandler := &PostV1ParcelsHandler{
		service: mockStorage,
	}

	getV1ParcelsHandler := &GetV1ParcelsHandler{
		service: mockStorage,
	}

	getV1RoutesHandler := &GetV1RoutesHandler{
		service: mockStorage,
	}

	return postV1ParcelsHandler, getV1ParcelsHandler, getV1RoutesHandler, *mockStorage
}

func TestGetV1ParcelsHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful retrieval of parcels", func(t *testing.T) {
		_, getHandler, _, mockService := setupHandlers(ctrl)

		mockService.EXPECT().GetTodayParcels().Return([]*models.ParcelV1{{ID: "1"}}, nil)

		result := getHandler.Handle(operations.GetV1ParcelsParams{})
		assert.IsType(t, &operations.GetV1ParcelsOK{}, result)

		okResult := result.(*operations.GetV1ParcelsOK)
		assert.Len(t, okResult.Payload.Parcels, 1)
		assert.Equal(t, "1", okResult.Payload.Parcels[0].ID)
	})

	t.Run("Internal server error when service fails", func(t *testing.T) {
		_, getHandler, _, mockService := setupHandlers(ctrl)

		mockService.EXPECT().GetTodayParcels().Return(nil, assert.AnError)

		result := getHandler.Handle(operations.GetV1ParcelsParams{})
		assert.IsType(t, &operations.GetV1ParcelsInternalServerError{}, result)
	})
}

func TestPostV1ParcelsHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful creation of parcel", func(t *testing.T) {
		postHandler, _, _, mockService := setupHandlers(ctrl)

		mockService.EXPECT().CreateParcel(gomock.Any()).Return(models.ParcelV1{}, nil)

		params := operations.PostV1ParcelsParams{
			CreateParcelV1Request: &models.ParcelV1{},
		}

		result := postHandler.Handle(params)
		assert.IsType(t, &operations.PostV1ParcelsOK{}, result)
	})

	t.Run("Bad request when parcel is nil", func(t *testing.T) {
		postHandler, _, _, _ := setupHandlers(ctrl)

		params := operations.PostV1ParcelsParams{
			CreateParcelV1Request: nil,
		}

		result := postHandler.Handle(params)
		assert.IsType(t, &operations.PostV1ParcelsBadRequest{}, result)
	})

	t.Run("Internal server error when service fails", func(t *testing.T) {
		postHandler, _, _, mockService := setupHandlers(ctrl)

		mockService.EXPECT().CreateParcel(gomock.Any()).Return(models.ParcelV1{}, assert.AnError)

		params := operations.PostV1ParcelsParams{
			CreateParcelV1Request: &models.ParcelV1{},
		}

		result := postHandler.Handle(params)
		assert.IsType(t, &operations.PostV1ParcelsInternalServerError{}, result)
	})
}

func TestGetV1RoutesHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful retrieval of routes", func(t *testing.T) {
		_, _, getRoutesHandler, mockService := setupHandlers(ctrl)

		mockService.EXPECT().GetTodayRoutes().Return([]*models.RouteV1{{ID: "route1"}}, nil)

		result := getRoutesHandler.Handle(operations.GetV1ParcelsRoutesParams{})
		assert.IsType(t, &operations.GetV1ParcelsRoutesOK{}, result)

		okResult := result.(*operations.GetV1ParcelsRoutesOK)
		assert.Len(t, okResult.Payload.Routes, 1)
		assert.Equal(t, "route1", okResult.Payload.Routes[0].ID)
	})

	t.Run("Internal server error when service fails", func(t *testing.T) {
		_, _, getRoutesHandler, mockService := setupHandlers(ctrl)

		mockService.EXPECT().GetTodayRoutes().Return(nil, assert.AnError)

		result := getRoutesHandler.Handle(operations.GetV1ParcelsRoutesParams{})
		assert.IsType(t, &operations.GetV1ParcelsRoutesInternalServerError{}, result)
	})
}
