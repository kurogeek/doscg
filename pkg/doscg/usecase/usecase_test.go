package usecase

import (
	"doscg/pkg/doscg/usecase/mocks"
	"doscg/pkg/entity"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestFindXYZ(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMap := mocks.NewMockMapService(mockCtrl)
	mockMessage := mocks.NewMockMessageService(mockCtrl)

	t.Run("success", func(t *testing.T) {
		doscgService := NewDoSCGService(mockMap, mockMessage)
		xyz := doscgService.FindXYZ()

		expected := entity.XYZ{
			X: 3,
			Y: 3,
			Z: 33,
		}
		assert.DeepEqual(t, xyz, expected)
	})
}

func TestFindBC(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMap := mocks.NewMockMapService(mockCtrl)
	mockMessage := mocks.NewMockMessageService(mockCtrl)

	t.Run("success", func(t *testing.T) {
		doscgService := NewDoSCGService(mockMap, mockMessage)
		bc := doscgService.FindBC(23, -21)

		expected := entity.BC{
			B: 2,
			C: -42,
		}
		assert.DeepEqual(t, bc, expected)
	})
}

func TestFindBestWayFromSCGToCentrallWorld(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMap := mocks.NewMockMapService(mockCtrl)
	mockMessage := mocks.NewMockMessageService(mockCtrl)

	t.Run("fail-map-service-error", func(t *testing.T) {

		mockMap.EXPECT().FindBestWayFromSCGToCentrallWorld(gomock.AssignableToTypeOf("origin"), gomock.AssignableToTypeOf("destination")).
			Return(entity.BestRoute{}, errors.New("map service error"))

		doscgService := NewDoSCGService(mockMap, mockMessage)

		route, err := doscgService.FindBestWayFromSCGToCentrallWorld()

		assert.Error(t, err, "map service error")
		assert.DeepEqual(t, route, entity.BestRoute{})

	})

	t.Run("success", func(t *testing.T) {
		expectRoute := entity.BestRoute{
			OriginLocation: entity.Place{
				Name: "point a",
				Location: entity.LatLng{
					Lat: 1.0,
					Lng: 1.0,
				},
			},
			DestinationLocation: entity.Place{
				Name: "point b",
				Location: entity.LatLng{
					Lat: 3.0,
					Lng: 3.0,
				},
			},
			Polyline: []entity.LatLng{
				{
					Lat: 1.0,
					Lng: 1.0,
				}, {
					Lat: 2.0,
					Lng: 2.0,
				}, {
					Lat: 3.0,
					Lng: 3.0,
				},
			},
		}
		mockMap.EXPECT().FindBestWayFromSCGToCentrallWorld(gomock.AssignableToTypeOf("origin"), gomock.AssignableToTypeOf("destination")).
			Return(expectRoute, nil).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		route, err := doscgService.FindBestWayFromSCGToCentrallWorld()
		assert.NilError(t, err)
		assert.DeepEqual(t, route, expectRoute)
	})
}

func TestBotHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMap := mocks.NewMockMapService(mockCtrl)
	mockMessage := mocks.NewMockMessageService(mockCtrl)

	t.Run("fail-reply-error", func(t *testing.T) {
		inMessage := entity.BotMessage{
			ReplyToken: "token",
			Text:       "hi",
		}
		mockMessage.EXPECT().SendReply(gomock.AssignableToTypeOf(inMessage)).
			Return(errors.New("send reply error")).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		err := doscgService.BotHandler(inMessage)
		assert.Error(t, err, "send reply error")
	})

	t.Run("success-reply-hi", func(t *testing.T) {
		inMessage := entity.BotMessage{
			ReplyToken: "token",
			Text:       "hi",
		}
		mockMessage.EXPECT().SendReply(gomock.AssignableToTypeOf(inMessage)).
			Return(nil).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		err := doscgService.BotHandler(inMessage)
		assert.NilError(t, err)
	})

	t.Run("success-reply-hello", func(t *testing.T) {
		inMessage := entity.BotMessage{
			ReplyToken: "token",
			Text:       "hello",
		}
		mockMessage.EXPECT().SendReply(gomock.AssignableToTypeOf(inMessage)).
			Return(nil).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		err := doscgService.BotHandler(inMessage)
		assert.NilError(t, err)
	})

	t.Run("fail-notify", func(t *testing.T) {
		inMessage := entity.BotMessage{
			ReplyToken: "token",
			Text:       "test",
		}
		mockMessage.EXPECT().NotifyBotError(gomock.AssignableToTypeOf("notify message")).
			Return(errors.New("notify error")).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		err := doscgService.BotHandler(inMessage)
		assert.Error(t, err, "notify error")

	})

	t.Run("success-notify", func(t *testing.T) {
		inMessage := entity.BotMessage{
			ReplyToken: "token",
			Text:       "test",
		}
		mockMessage.EXPECT().NotifyBotError(gomock.AssignableToTypeOf("notify message")).
			Return(nil).
			Times(1)

		doscgService := NewDoSCGService(mockMap, mockMessage)

		err := doscgService.BotHandler(inMessage)
		assert.NilError(t, err)
	})
}
