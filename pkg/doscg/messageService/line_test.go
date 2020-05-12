package messageService

import (
	"doscg/pkg/entity"
	"testing"

	"github.com/jarcoal/httpmock"
	"gotest.tools/v3/assert"
)

func TestSendReply(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.line.me/v2/bot/message/reply"

	t.Run("fail-internal-error", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url)

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		t.Log(err)
	})

}

func TestNotifyBotError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

}
