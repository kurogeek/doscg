package mapService

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewGoogleService(t *testing.T) {}

func TestFindBestWayFromSCGToCentrallWorld(t *testing.T) {
	key := os.Getenv("GOOGLE_TEST_KEY")

	googleService, err := NewGoogleService(key)
	assert.NilError(t, err)

	googleService.
}
