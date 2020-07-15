package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

const (
	chatID              = int64(1057858355)
	alertToken          = "1248026431:AAF6ZzklChI5zeP2_-9CzGjWOt_QVfXGSzw"
	riskManagementToken = "1360876104:AAGpFNzUWNWL0eGxLO4LQYqwhU_jy3WfXq4"
)

func SendAlertMessage(messages []string) error {
	alertBot, err := tgbotapi.NewBotAPI(alertToken)
	if err != nil {
		log.Println("SendAlertMessage failed at NewBotAPI ", err)
		return err
	}

	resMessage := ""
	for _, message := range messages {
		resMessage = resMessage + message + "\n"
	}

	_, err = alertBot.Send(tgbotapi.NewMessage(chatID, resMessage))
	if err != nil {
		log.Println("SendAlertMessage failed at Send ", err)
		return err
	}
	return nil
}

func SendRiskManagementMessage(messages []string) error {
	alertBot, err := tgbotapi.NewBotAPI(riskManagementToken)
	if err != nil {
		log.Println("SendRiskManagementMessage failed at NewBotAPI ", err)
		return err
	}

	resMessage := ""
	for _, message := range messages {
		resMessage = resMessage + message + "\n"
	}

	_, err = alertBot.Send(tgbotapi.NewMessage(chatID, resMessage))
	if err != nil {
		log.Println("SendRiskManagementMessage failed at Send ", err)
		return err
	}
	return nil
}
