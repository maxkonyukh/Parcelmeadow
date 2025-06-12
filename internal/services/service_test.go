package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"parcelmeadow/internal/api/generated/models"
	"parcelmeadow/internal/database"
	"testing"
	"time"

	mock_services "parcelmeadow/internal/services/mock/generated"
)

func setupService(ctrl *gomock.Controller) (*ParcelmeadowService, mock_services.MockStorage) {
	storage := mock_services.NewMockStorage(ctrl)
	service, _ := NewParcelmeadowService(storage)
	return service, *storage
}

func TestParcelmeadowService_CreateParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("should create a parcel", func(t *testing.T) {
		service, storage := setupService(ctrl)

		storage.EXPECT().SaveParcel(gomock.Any()).Return(database.Parcel{}, nil).Times(1)

		parcel := &models.ParcelV1{
			ID:       uuid.NewString(),
			Weight:   1.5,
			Status:   "created",
			PostCode: "12345",
			Address:  "123 Meadow Lane",
		}

		savedParcel, err := service.CreateParcel(parcel)
		assert.NotNil(t, savedParcel)
		assert.NoError(t, err)
	})

	t.Run("should return error for nil parcel", func(t *testing.T) {
		service, _ := setupService(ctrl)
		_, err := service.CreateParcel(nil)
		assert.Error(t, err)
	})

	t.Run("should return error when storage fails", func(t *testing.T) {
		service, storage := setupService(ctrl)

		storage.EXPECT().SaveParcel(gomock.Any()).Return(database.Parcel{}, fmt.Errorf("error")).Times(1)

		parcel := &models.ParcelV1{
			ID:       uuid.NewString(),
			Weight:   1.5,
			Status:   "created",
			PostCode: "12345",
			Address:  "123 Meadow Lane",
		}

		_, err := service.CreateParcel(parcel)
		assert.Error(t, err)
	})
}

func TestParcelmeadowService_GetTodayParcels(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("should get today's parcels", func(t *testing.T) {
		service, storage := setupService(ctrl)

		expectedParcels := []database.Parcel{
			{Id: uuid.NewString(), Weight: 1.5, Status: "created", PostCode: "12345", Address: "123 Meadow Lane", CreatedAt: time.Now().UTC()},
		}

		storage.EXPECT().GetTodayParcels().Return(expectedParcels, nil).Times(1)

		parcels, err := service.GetTodayParcels()
		assert.NoError(t, err)
		assert.Len(t, parcels, 1)
		assert.Equal(t, expectedParcels[0].Id, parcels[0].ID)
	})

	t.Run("should return empty slice when no parcels found", func(t *testing.T) {
		service, storage := setupService(ctrl)

		storage.EXPECT().GetTodayParcels().Return([]database.Parcel{}, nil).Times(1)

		parcels, err := service.GetTodayParcels()
		assert.NoError(t, err)
		assert.Len(t, parcels, 0)
	})

	t.Run("should return error when storage fails", func(t *testing.T) {
		service, storage := setupService(ctrl)

		storage.EXPECT().GetTodayParcels().Return(nil, fmt.Errorf("error")).Times(1)

		parcels, err := service.GetTodayParcels()
		assert.Error(t, err)
		assert.Nil(t, parcels)
	})
}

func TestParcelmeadowService_GetTodayRoutes(t *testing.T) {
	t.Run("should get today's routes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		service, storage := setupService(ctrl)

		expectedRoutes := []database.Route{
			{Id: uuid.NewString(), Stops: database.DeliveryStops{
				"stop1": {Id: "stop1", Address: "123 Meadow Lane", Parcels: []database.Parcel{{Id: uuid.NewString(), Weight: 1.5, Status: "created", PostCode: "12345", Address: "123 Meadow Lane", CreatedAt: time.Now().UTC()}}},
			}},
		}

		storage.EXPECT().GetTodayRoutes().Return(expectedRoutes, nil).Times(1)

		routes, err := service.GetTodayRoutes()
		assert.NoError(t, err)
		assert.Len(t, routes, 1)
		assert.Equal(t, expectedRoutes[0].Id, routes[0].ID)
	})

	t.Run("should return empty slice when no routes found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		service, storage := setupService(ctrl)

		storage.EXPECT().GetTodayRoutes().Return([]database.Route{}, nil).Times(1)

		routes, err := service.GetTodayRoutes()
		assert.NoError(t, err)
		assert.Len(t, routes, 0)
	})

	t.Run("should return error when storage fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		service, storage := setupService(ctrl)

		storage.EXPECT().GetTodayRoutes().Return(nil, fmt.Errorf("error")).Times(1)

		routes, err := service.GetTodayRoutes()
		assert.Error(t, err)
		assert.Nil(t, routes)
	})
}
