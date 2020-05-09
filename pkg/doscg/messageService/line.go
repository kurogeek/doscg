package messageService

import (
	"doscg/pkg/doscg"
	"doscg/pkg/entity"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/utahta/go-linenotify"
)

type messageService struct {
	notifyToken   string
	NotifyClient  *linenotify.Client
	LineBotClient *linebot.Client
}

func NewMessageService(notiToken string, botChanSec string, botChanToken string) (doscg.MessageService, error) {
	nc := newNotifyClient()
	lbc, err := newLineBotClient(botChanSec, botChanToken)
	if err != nil {
		return &messageService{}, err
	}

	return &messageService{
		notifyToken:   notiToken,
		NotifyClient:  nc,
		LineBotClient: lbc,
	}, nil
}

func newNotifyClient() *linenotify.Client {
	return linenotify.New()
}

func newLineBotClient(botChanSec, botChanToken string) (*linebot.Client, error) {
	lbc, err := linebot.New(botChanSec, botChanToken)
	return lbc, err
}

func (ms messageService) SendReply(replyMessage entity.BotMessage) error {
	_, err := ms.LineBotClient.ReplyMessage(replyMessage.ReplyToken, linebot.NewTextMessage(replyMessage.Text)).Do()
	return err
}

func (ms messageService) NotifyBotError(notiText string) error {
	_, err := ms.NotifyClient.NotifyMessage(ms.notifyToken, notiText)
	return err
}
