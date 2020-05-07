package mapService

import (
	"context"
	"doscg/pkg/doscg"
	"doscg/pkg/entity"

	"googlemaps.github.io/maps"
)

type googleService struct {
	GoogleClient *maps.Client
}

func NewGoogleClient(key string) (*maps.Client, error) {
	gc, err := maps.NewClient(maps.WithAPIKey(key))
	return gc, err
}

func NewGoogleService(mc *maps.Client) doscg.MapService {
	return &googleService{
		GoogleClient: mc,
	}
}

func (gs googleService) FindBestWayFromSCGToCentrallWorld(origin string, destination string) (entity.BestRoute, error) {
	var res entity.BestRoute
	request := maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	routes, _, err := gs.GoogleClient.Directions(context.Background(), &request)
	if err != nil {
		return res, err
	}
	if len(routes) == 0 {
		return res, entity.NoRouteError
	}
	legs := routes[0].Legs

	if len(legs) == 0 {
		return res, entity.NoLegsFound
	}
	leg := legs[0]

	res.OriginLocation.Name = leg.StartAddress
	res.OriginLocation.Location.Lat = leg.StartLocation.Lat
	res.OriginLocation.Location.Lng = leg.StartLocation.Lng
	res.DestinationLocation.Name = leg.EndAddress
	res.DestinationLocation.Location.Lat = leg.EndLocation.Lat
	res.DestinationLocation.Location.Lng = leg.EndLocation.Lng
	res.Polyline = routes[0].OverviewPolyline.Points
}
