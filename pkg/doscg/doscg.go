package doscg

import (
	"doscg/pkg/entity"
)

type DoSCG interface {
	FindXYZ() entity.XYZ
	FindBC(ans1 int, ans2 int) entity.BC
	FindBestWayToSCG(start string)
}

type MapService interface {
	FindBestWayToSCG(start string)
}
