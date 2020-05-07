package doscg

import (
	"doscg/pkg/entity"
)

type DoSCG interface {
	FindXYZ() entity.XYZ
	FindBC(ans1 int, ans2 int) entity.BC
	FindBestWayFromSCGToCentrallWorld() (entity.BestRoute, error)
}

type MapService interface {
	FindBestWayFromSCGToCentrallWorld(origin string, destination string) (entity.BestRoute, error)
}
