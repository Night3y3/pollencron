package tests

import (
	"testing"

	"github.com/Night3y3/pollencron/utils"
)

func TestSendSlackMessage(t *testing.T) {
	err := utils.SendSlackMessage("Test message")
	if err != nil {
		t.Errorf("Error while sending slack message: %v", err)
		return
	}
}