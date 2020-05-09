package doscg

import (
	"doscg/pkg/entity"
)

type DoSCG interface {
	FindXYZ() entity.XYZ
	FindBC(ans1 int, ans2 int) entity.BC
	FindBestWayFromSCGToCentrallWorld() (entity.BestRoute, error)
	BotHandler(inMessage entity.BotMessage) error
}

//go:generate mockgen -destination=./usecase/mocks/mock_map.go -package=mocks doscg/pkg/doscg MapService
type MapService interface {
	FindBestWayFromSCGToCentrallWorld(origin string, destination string) (entity.BestRoute, error)
}

//go:generate mockgen -destination=./usecase/mocks/mock_message.go -package=mocks doscg/pkg/doscg MessageService
type MessageService interface {
	SendReply(replyMessage entity.BotMessage) error
	NotifyBotError(notiText string) error
}
