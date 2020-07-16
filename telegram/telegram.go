package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type Channel struct {
	Name   string
	ChatID int64
	Token  string
}

type Telegram struct {
	bots map[string]Channel
}

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
	alertBot, err := tgbotapi.NewBotAPI(t.bots[channelName].Token)
	if err != nil {
		log.Println("SendAlertMessage failed at NewBotAPI ", err)
		return err
	}

	resMessage := ""
	for _, message := range messages {
		resMessage = resMessage + message + "\n"
	}

	_, err = alertBot.Send(tgbotapi.NewMessage(t.bots[channelName].ChatID, resMessage))
	if err != nil {
		log.Println("SendAlertMessage failed at Send ", err)
		return err
	}
	return nil
}
