package handler

import (
	"doscg/pkg/doscg"

	"github.com/labstack/echo"
)

type DoSCGHandler struct {
	SCGService doscg.DoSCG
}

func NewDoSCGHandler(e *echo.Echo, du doscg.DoSCG) {

}
