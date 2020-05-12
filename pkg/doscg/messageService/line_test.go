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

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(500, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.ErrorContains(t, err, "500")
	})
	t.Run("fail-too-many-request", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(429, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.ErrorContains(t, err, "429")
	})
	t.Run("fail-forbidden", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(403, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.ErrorContains(t, err, "403")
	})
	t.Run("fail-unauthorized", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(401, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.ErrorContains(t, err, "401")
	})
	t.Run("fail-bad-request", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(400, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.ErrorContains(t, err, "400")
	})

	t.Run("success", func(t *testing.T) {
		lService, err := NewMessageService("test-notitoken", "test-secret", "test-chan-token")
		assert.NilError(t, err)

		httpmock.RegisterResponder("POST", url,
			httpmock.NewStringResponder(200, "{}"))

		replyMessage := entity.BotMessage{
			ReplyToken: "test-reply-token",
			Text:       "test reply",
		}
		err = lService.SendReply(replyMessage)

		assert.NilError(t, err)
	})

}

func TestNotifyBotError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

}
