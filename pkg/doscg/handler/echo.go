package handler

import (
	"doscg/pkg/doscg"

	"github.com/labstack/echo"
)

type DoSCGHandler struct {
	SCGService doscg.DoSCG
}

func NewDoSCGHandler(e *echo.Echo, du doscg.DoSCG) {
	handler := &DoSCGHandler{
		SCGService: du,
	}
	e.GET("/find-xyz", handler.FindXYZ)
}

func (dh DoSCGHandler) FindXYZ(c echo.Context) error {

	xyz := dh.SCGService.FindXYZ()

	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Type:       "xyz",
			Attributes: xyz,
		},
	})
}

func (dh DoSCGHandler) FindBC(c echo.Context) error {

	ans1 := 23
	ans2 := -21
	bc := dh.SCGService.FindBC(ans1, ans2)

	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Type:       "bc",
			Attributes: bc,
		},
	})
}
