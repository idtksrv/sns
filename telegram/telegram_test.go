package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendAlertMessage(t *testing.T) {
	assert.NoError(t, SendAlertMessage([]string{"Alert1", "Alert2", "Alert3"}))
}

func TestSendRiskManagementMessage(t *testing.T) {
	assert.NoError(t, SendRiskManagementMessage([]string{"Management1", "Management2", "Management3"}))
}
