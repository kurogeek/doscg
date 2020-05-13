package mapService

import (
	"doscg/pkg/entity"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"gotest.tools/v3/assert"
)

const filePath = "./test-data/"

func TestFindBestWayFromSCGToCentrallWorld(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("fail-not-found", func(t *testing.T) {

		gService, err := NewGoogleService("test-key")
		assert.NilError(t, err)

		data, err := readFile("fail-not-found.json")
		assert.NilError(t, err)

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"
		httpmock.RegisterResponder("GET",
			url,
			httpmock.NewStringResponder(200, data))

		route, err := gService.FindBestWayFromSCGToCentrallWorld("origin", "destination")

		info := httpmock.GetCallCountInfo()
		calls := info[fmt.Sprintf("GET %s", url)]

		assert.Assert(t, calls == 1)
		assert.ErrorContains(t, err, "NOT_FOUND")
		assert.DeepEqual(t, route, entity.BestRoute{})

	})

	t.Run("fail-no-route", func(t *testing.T) {
		gService, err := NewGoogleService("test-key")
		assert.NilError(t, err)

		data, err := readFile("fail-no-route.json")
		assert.NilError(t, err)

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"
		httpmock.RegisterResponder("GET",
			url,
			httpmock.NewStringResponder(200, data))

		route, err := gService.FindBestWayFromSCGToCentrallWorld("origin", "destination")

		info := httpmock.GetCallCountInfo()
		calls := info[fmt.Sprintf("GET %s", url)]

		assert.Assert(t, calls == 1)
		assert.DeepEqual(t, route, entity.BestRoute{})
		assert.Error(t, err, entity.NoRouteError.Error())

	})

	t.Run("fail-no-legs", func(t *testing.T) {
		gService, err := NewGoogleService("test-key")
		assert.NilError(t, err)

		data, err := readFile("fail-no-legs.json")
		assert.NilError(t, err)

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"

		httpmock.RegisterResponder("GET",
			url,
			httpmock.NewStringResponder(200, data))

		route, err := gService.FindBestWayFromSCGToCentrallWorld("origin", "destination")

		info := httpmock.GetCallCountInfo()
		calls := info[fmt.Sprintf("GET %s", url)]

		assert.Assert(t, calls == 1)
		assert.DeepEqual(t, route, entity.BestRoute{})
		assert.Error(t, err, entity.NoLegsFound.Error())
	})

	t.Run("success", func(t *testing.T) {
		gService, err := NewGoogleService("test-key")
		assert.NilError(t, err)

		data, err := readFile("success.json")
		assert.NilError(t, err)

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"

		httpmock.RegisterResponder("GET",
			url,
			httpmock.NewStringResponder(200, data))

		route, err := gService.FindBestWayFromSCGToCentrallWorld("origin", "destination")

		expectRoute := entity.BestRoute{
			OriginLocation: entity.Place{
				Name:     "Chiang Mai International Airport (CNX), 60 Mahidol Rd, Amphoe Mueang Chiang Mai, Chang Wat Chiang Mai 50200, Thailand",
				Location: entity.LatLng{Lat: 18.766545, Lng: 98.968487},
			},
			DestinationLocation: entity.Place{
				Name:     "2 Mahidol Rd, Tambon Pa Daet, Amphoe Mueang Chiang Mai, Chang Wat Chiang Mai 50100, Thailand",
				Location: entity.LatLng{Lat: 18.7702169, Lng: 98.9749753},
			},
			Polyline: []entity.LatLng{
				{Lat: 18.766540000000003, Lng: 98.96849},
				{Lat: 18.77118, Lng: 98.96847000000001},
				{Lat: 18.771980000000003, Lng: 98.96846000000001},
				{Lat: 18.77205, Lng: 98.96846000000001},
				{Lat: 18.77204, Lng: 98.96901000000001},
				{Lat: 18.77202, Lng: 98.96960000000001},
				{Lat: 18.77203, Lng: 98.97014000000001},
				{Lat: 18.77195, Lng: 98.97074},
				{Lat: 18.77166, Lng: 98.97234},
				{Lat: 18.77154, Lng: 98.97291000000001},
				{Lat: 18.77148, Lng: 98.9731},
				{Lat: 18.771230000000003, Lng: 98.97365},
				{Lat: 18.77051, Lng: 98.97484000000001},
				{Lat: 18.77027, Lng: 98.97525},
				{Lat: 18.77018, Lng: 98.97543},
				{Lat: 18.77015, Lng: 98.97560000000001},
				{Lat: 18.77014, Lng: 98.97584},
				{Lat: 18.77014, Lng: 98.97593},
				{Lat: 18.77015, Lng: 98.97600000000001},
				{Lat: 18.770110000000003, Lng: 98.97602},
				{Lat: 18.77006, Lng: 98.97603000000001},
				{Lat: 18.77005, Lng: 98.97596000000001},
				{Lat: 18.77004, Lng: 98.97587000000001},
				{Lat: 18.770010000000003, Lng: 98.97575},
				{Lat: 18.769930000000002, Lng: 98.97562},
				{Lat: 18.76987, Lng: 98.97552},
				{Lat: 18.769910000000003, Lng: 98.97548},
				{Lat: 18.769930000000002, Lng: 98.97543},
				{Lat: 18.77004, Lng: 98.97525},
				{Lat: 18.770220000000002, Lng: 98.97498},
			},
		}

		info := httpmock.GetCallCountInfo()
		calls := info[fmt.Sprintf("GET %s", url)]

		assert.Assert(t, calls == 1)
		assert.NilError(t, err)
		assert.DeepEqual(t, route, expectRoute)
	})
}

func readFile(name string) (string, error) {
	out, err := ioutil.ReadFile(filePath + name)
	return string(out), err
}
