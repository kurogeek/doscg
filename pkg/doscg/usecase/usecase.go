package usecase

import (
	"doscg/pkg/doscg"
	"doscg/pkg/entity"
)

type doSCGUsecase struct {
	mapService     doscg.MapService
	messageService doscg.MessageService
}

var messageList = [...]string{"hi", "hello"}

func NewDoSCGService(gs doscg.MapService, ms doscg.MessageService) doscg.DoSCG {
	return &doSCGUsecase{
		mapService:     gs,
		messageService: ms,
	}
}

func (ds doSCGUsecase) FindBC(ans1 int, ans2 int) entity.BC {
	var a int = 21
	var bc entity.BC

	bc.B = ds.findNumber(a, ans1)
	bc.C = ds.findNumber(a, ans2)

	return bc
}

func (ds doSCGUsecase) findNumber(a int, ans int) int {
	return ans - a
}

// FindXYZ - X, Y, 5, 9, 15, 23, Z  - Please create a new function for finding X, Y, Z value
// X, Y, 5, 9, 15, 23, Z
//  \/ \/ \/ \/  \/  \/
//   0  2  4  6  8   10
// So, Z = 23 + 10 = 33
// Y = 5 - 2 = 3
// X = Y - 0 = 3 - 0 = 3
func (ds doSCGUsecase) FindXYZ() entity.XYZ {
	var xyz entity.XYZ
	xyz.X = ds.findX()
	xyz.Y = ds.findY()
	xyz.Z = ds.findZ()
	return xyz
}

func (ds doSCGUsecase) findX() int {
	return 3
}

func (ds doSCGUsecase) findY() int {
	return 3
}

func (ds doSCGUsecase) findZ() int {
	return 33
}

func (ds doSCGUsecase) FindBestWayFromSCGToCentrallWorld() (entity.BestRoute, error) {
	var bestRoute entity.BestRoute
	origin := "SCG สำนักงานใหญ่ บางซื่อ 1 Siam Cement Alley, Bang Sue, Bangkok 10800"
	destination := "centralwOrld, 999/9 Rama I Rd, Pathum Wan, Pathum Wan District, Bangkok 10330"

	bestRoute, err := ds.mapService.FindBestWayFromSCGToCentrallWorld(origin, destination)
	if err != nil {
		return bestRoute, err
	}

	return bestRoute, nil
}

func (ds doSCGUsecase) BotHandler(inMessage entity.BotMessage) error {
	replyMessage := inMessage
	replyMessage.Text = ds.getReplyMessage(inMessage.Text)
	if replyMessage.Text != "" {
		return ds.messageService.SendReply(replyMessage)
	}
	return ds.messageService.NotifyBotError("Error! bot cannot handle request.")
}

func (ds doSCGUsecase) getReplyMessage(in string) string {
	for _, v := range messageList {
		if v == in {
			return "Hi, how can I help you?"
		}
	}
	return ""
}
