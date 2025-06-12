package database

import "time"

type Parcel struct {
	Id        string
	Status    ParcelStatus
	Weight    float32
	PostCode  string
	Address   string
	CreatedAt time.Time
}

type Route struct {
	Id    string
	Stops DeliveryStops
}

type Stop struct {
	Id      string
	Address string
	Parcels []Parcel
}

type ParcelStatus string

type RouteName string

const (
	RouteUnknown RouteName = "unknown"
	Route1       RouteName = "R001"
	Route2       RouteName = "R002"
	Route3       RouteName = "R003"
	Route4       RouteName = "R004"
	Route5       RouteName = "R005"
)

type DeliveryRoutes map[RouteName]Route

type DeliveryStops map[string]Stop
