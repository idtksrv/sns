package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var (
	ErrChannelNotFound = errors.New("channel not found")
)

type Channel struct {
	Name             string
	ChatID           int64
	Token            string
	ReplyToMessageID int // 指定子頻道的 message_thread_id
}

type Telegram struct {
	bots map[string]Channel
}

// NewTelegram channel name 隨便填
func NewTelegram(channels []Channel) *Telegram {
	t := &Telegram{
		bots: map[string]Channel{},
	}
	for i := 0; i < len(channels); i++ {
		t.bots[channels[i].Name] = channels[i]
	}
	return t
}

func (t *Telegram) SendMessage(channelName string, messages []string) error {
	if _, ok := t.bots[channelName]; !ok {
		return ErrChannelNotFound
	}

	alertBot, err := tgbotapi.NewBotAPI(t.bots[channelName].Token)
	if err != nil {
		log.Println("SendAlertMessage failed at NewBotAPI ", err)
		return err
	}

	resMessage := ""
	for _, message := range messages {
		resMessage = resMessage + message + "\n"
	}

	msg := tgbotapi.NewMessage(t.bots[channelName].ChatID, resMessage)
	msg.ReplyToMessageID = t.bots[channelName].ReplyToMessageID // 指定子頻道的 message_thread_id
	_, err = alertBot.Send(msg)
	if err != nil {
		log.Println("SendAlertMessage failed at Send ", err)
		return err
	}
	return nil
}
