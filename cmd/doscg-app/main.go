package main

import (
	"doscg/pkg/doscg/handler"
	"doscg/pkg/doscg/mapService"
	"doscg/pkg/doscg/messageService"
	"doscg/pkg/doscg/usecase"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	gkey := os.Getenv("GOOGLE_KEY")
	gService, err := mapService.NewGoogleService(gkey)
	if err != nil {
		panic(err)
	}

	notiToken := os.Getenv("LINE_NOTI_TOKEN")
	botChanSecret := os.Getenv("LINE_CHAN_SECRET")
	botChanToken := os.Getenv("LINE_CHAN_TOKEN")
	lService, err := messageService.NewMessageService(notiToken, botChanSecret, botChanToken)
	if err != nil {
		panic(err)
	}

	scgUsecase := usecase.NewDoSCGService(gService, lService)

	e := echo.New()
	e.Use(middleware.Logger())

	handler.NewDoSCGHandler(e, scgUsecase, botChanSecret)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))

}
