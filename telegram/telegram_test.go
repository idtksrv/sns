package telegram

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	chatID              = int64(1057858355)
	alertToken          = "1248026431:AAF6ZzklChI5zeP2_-9CzGjWOt_QVfXGSzw"
	riskManagementToken = "1360876104:AAGpFNzUWNWL0eGxLO4LQYqwhU_jy3WfXq4"

	//notify special
	notifySpecialToken  = "5054656265:AAHFHe9Pu2BE6Pavi7tcTgIjaeA1TiBNops"
	notifySpecialChatID = int64(-746492368)
)

func TestSendNotifySpecialAlertMessage(t *testing.T) {
	s := NewTelegram([]Channel{
		Channel{
			Name:   "alert",
			ChatID: notifySpecialChatID,
			Token:  notifySpecialToken,
		},
	})
	assert.NoError(t, s.SendMessage("alert", []string{"Alert1", "Alert2", "Alert3"}))
}

func TestSendAlertMessage(t *testing.T) {
	s := NewTelegram([]Channel{
		Channel{
			Name:   "alert",
			ChatID: chatID,
			Token:  alertToken,
		},
	})
	assert.NoError(t, s.SendMessage("alert", []string{"Alert1", "Alert2", "Alert3"}))
}

func TestSendRiskManagementMessage(t *testing.T) {
	s := NewTelegram([]Channel{
		Channel{
			Name:   "riskManagement",
			ChatID: chatID,
			Token:  riskManagementToken,
		},
	})
	assert.NoError(t, s.SendMessage("riskManagement", []string{"Management1", "Management2", "Management3"}))
}

func TestSendCDNMessage(t *testing.T) {

	channelName := "CdnAlert"
	s := NewTelegram([]Channel{
		{
			Name: channelName,
			// ChatID: -4231511444,
			// Token:  "7050702530:AAGSCwBDGkVb9VgczqBRz230JY7z03tAmHY",
			ChatID:           -1001912947918, // CDN切換 頻道
			Token:            "6124231819:AAEhPPaGjXZZ7gQa5S4RiZitwj2Uj3VjCKA",
			ReplyToMessageID: 1611,
		},
	})

	env := "dev_unit_test"
	account := "test"
	selectNow := 1
	value := "1.online"
	var messages []string
	messages = append(messages, "[通知]")
	messages = append(messages, "事由: CDN切換")
	messages = append(messages, fmt.Sprintf("環境: %s", env))
	messages = append(messages, fmt.Sprintf("帳戶: %s", account))
	messages = append(messages, fmt.Sprintf("新 CDN key: %d", selectNow))
	messages = append(messages, fmt.Sprintf("新 CDN value: %s", value))
	fmt.Printf("notifyToTelegram success, time:%v, messages:%+v\n", time.Now().Format("2006-01-02 15:04:05"), messages)

	err := s.SendMessage(channelName, messages)
	if err != nil {
		assert.NoError(t, err)
		return
	}

}
