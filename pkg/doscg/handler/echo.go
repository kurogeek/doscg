package handler

import (
	"doscg/pkg/doscg"
	"doscg/pkg/entity"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

type DoSCGHandler struct {
	SCGService        doscg.DoSCG
	LineSecretChannel string
}

func NewDoSCGHandler(e *echo.Echo, du doscg.DoSCG, lineSec string) {
	handler := &DoSCGHandler{
		SCGService:        du,
		LineSecretChannel: lineSec,
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

func (dh DoSCGHandler) FindBestWayFromSCGToCentrallWorld(c echo.Context) error {
	route, err := dh.SCGService.FindBestWayFromSCGToCentrallWorld()

	if err != nil {
		return c.JSON(200, ResposeMessage{
			Error: &ErrorMessage{
				Code:    500,
				Message: err.Error(),
			},
		})
	}

	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Type:       "route",
			Attributes: route,
		},
	})
}

func (dh DoSCGHandler) BotHandler(c echo.Context) error {
	req := c.Request()
	events, err := linebot.ParseRequest(dh.LineSecretChannel, req)
	if err != nil {
		return c.JSON(200, ErrorMessage{
			Code:    400,
			Message: err.Error(),
		})
	}

	for _, e := range events {
		var inMessage entity.BotMessage
		inMessage.ReplyToken = e.ReplyToken
		if e.Type == linebot.EventTypeMessage {
			switch message := e.Message.(type) {
			case *linebot.TextMessage:
				inMessage.Text = message.Text
				err := dh.SCGService.BotHandler(inMessage)
				if err != nil {
					return c.JSON(200, ErrorMessage{
						Code:    500,
						Message: err.Error(),
					})
				}
			default:
				inMessage.Text = ""
				err := dh.SCGService.BotHandler(inMessage)
				if err != nil {
					return c.JSON(200, ErrorMessage{
						Code:    500,
						Message: err.Error(),
					})
				}
			}
		}
	}

	return c.JSON(200, DataMessage{
		Type:       "ok",
		Attributes: Success{Status: "ok"},
	})
}
