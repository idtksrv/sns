package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	chatID              = int64(1057858355)
	alertToken          = "1248026431:AAF6ZzklChI5zeP2_-9CzGjWOt_QVfXGSzw"
	riskManagementToken = "1360876104:AAGpFNzUWNWL0eGxLO4LQYqwhU_jy3WfXq4"

	//notify special
	notifySpecialToken="5054656265:AAHFHe9Pu2BE6Pavi7tcTgIjaeA1TiBNops"
	notifySpecialChatID=int64(-746492368)
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
