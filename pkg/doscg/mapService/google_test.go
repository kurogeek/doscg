package mapService

import (
	"doscg/pkg/entity"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"gotest.tools/v3/assert"
)

func TestNewGoogleService(t *testing.T) {}

func TestFindBestWayFromSCGToCentrallWorld(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("fail not found", func(t *testing.T) {

		gService, err := NewGoogleService("test-key")
		assert.NilError(t, err)

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"
		httpmock.RegisterResponder("GET",
			url,
			httpmock.NewStringResponder(200, `{
				"geocoded_waypoints": [
				   {
					  "geocoder_status": "ZERO_RESULTS"
				   },
				   {
					  "geocoder_status": "ZERO_RESULTS"
				   }
				],
				"routes": [],
				"status": "NOT_FOUND"
				}`))

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

		url := "https://maps.googleapis.com/maps/api/directions/json?destination=destination&key=test-key&origin=origin"

	})
}
