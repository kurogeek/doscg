package mapService

import (
	"context"
	"doscg/pkg/doscg"
	"doscg/pkg/entity"
	"net/http"

	"googlemaps.github.io/maps"
)

type googleService struct {
	GoogleClient *maps.Client
}

func NewGoogleService(key string, options ...maps.ClientOption) (doscg.MapService, error) {
	keyOption := maps.WithAPIKey(key)
	options = append(options, keyOption)
	mc, err := newGoogleClient(options...)
	if err != nil {
		return &googleService{}, err
	}
	return &googleService{
		GoogleClient: mc,
	}, nil
}

func WithHTTPClient(c *http.Client) maps.ClientOption {
	return maps.WithHTTPClient(c)
}

func newGoogleClient(options ...maps.ClientOption) (*maps.Client, error) {
	gc, err := maps.NewClient(options...)
	return gc, err
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

	decodedPolylines, err := routes[0].OverviewPolyline.Decode()
	if err != nil {
		return res, err
	}

	for _, e := range decodedPolylines {
		res.Polyline = append(res.Polyline, entity.LatLng{
			Lat: e.Lat,
			Lng: e.Lng,
		})
	}
	return res, nil
}
