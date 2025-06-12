package database

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type InMemoryStorage struct {
	parcels     []Parcel
	routeNames  []RouteName
	routes      DeliveryRoutes
	latestRoute RouteName
}

// NewInMemoryStorage creates a new instance of InMemoryStorage.
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		parcels: []Parcel{},
		routeNames: []RouteName{
			Route1,
			Route2,
			Route3,
			Route4,
			Route5,
		},
		latestRoute: RouteUnknown,
		routes:      DeliveryRoutes{},
	}
}

func (s *InMemoryStorage) SaveParcel(parcel Parcel) (Parcel, error) {
	for _, p := range s.parcels {
		if parcel.Id == p.Id {
			return Parcel{}, fmt.Errorf("parcel with id %s already exists", parcel.Id)
		}
	}

	s.parcels = append(s.parcels, parcel)
	s.setRouteForParcel(parcel)
	return parcel, nil
}

func (s *InMemoryStorage) GetTodayParcels() ([]Parcel, error) {
	var todayParcels []Parcel
	for _, parcel := range s.parcels {
		if time.Now().UTC().Day() == parcel.CreatedAt.Day() {
			todayParcels = append(todayParcels, parcel)
		}
	}

	return todayParcels, nil
}

func (s *InMemoryStorage) setRouteForParcel(parcel Parcel) {
	nextRouteName := s.getNextRouteName()
	if route, routeExists := s.routes[nextRouteName]; routeExists {
		if stop, stopExists := route.Stops[parcel.Address]; stopExists {
			stop.Parcels = append(stop.Parcels, parcel)
			route.Stops[parcel.Address] = stop
		} else {
			route.Stops[parcel.Address] = Stop{
				Id:      uuid.NewString(),
				Address: parcel.Address,
				Parcels: []Parcel{parcel},
			}
		}
	} else {
		s.routes[nextRouteName] = Route{
			Id: string(nextRouteName),
			Stops: DeliveryStops{
				parcel.Address: Stop{
					Id:      uuid.NewString(),
					Address: parcel.Address,
					Parcels: []Parcel{parcel},
				},
			},
		}
	}
}

func (s *InMemoryStorage) GetTodayRoutes() ([]Route, error) {
	var todayRoutes []Route
	for _, route := range s.routes {
		if isValidRoute(route) {
			todayRoutes = append(todayRoutes, route)
		}
	}

	return todayRoutes, nil
}

func (s *InMemoryStorage) getNextRouteName() RouteName {
	if s.latestRoute == RouteUnknown || s.latestRoute == Route5 {
		s.latestRoute = Route1
	} else if s.latestRoute == Route1 {
		s.latestRoute = Route2
	} else if s.latestRoute == Route2 {
		s.latestRoute = Route3
	} else if s.latestRoute == Route3 {
		s.latestRoute = Route4
	} else {
		s.latestRoute = Route5
	}

	return s.latestRoute
}

func isValidRoute(route Route) bool {
	if len(route.Stops) == 0 {
		return false
	}

	for _, stop := range route.Stops {
		for _, parcel := range stop.Parcels {
			if time.Now().UTC().Day() == parcel.CreatedAt.Day() {
				return true
			}
		}
	}

	return false
}
