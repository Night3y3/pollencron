package tests

import (
	"testing"

	"github.com/Night3y3/pollencron/allergy_api"
)

func TestAllergyApi(t *testing.T) {
	message, err := allergy_api.GetHourlyLoadData()
	if err != nil {
		t.Errorf("Error while getting hourly load data: %v", err)
		return
	}

	if message == nil {
		t.Errorf("Message is nil")
		return
	}

	if *message == "" {
		t.Errorf("Message is empty")
	}

	message, err = allergy_api.GetCurrentChartData()
	if err != nil {
		t.Errorf("Error while getting current chart data: %v", err)
		return
	}

	if message == nil {
		t.Errorf("Message is nil")
		return
	}

	if *message == "" {
		t.Errorf("Message is empty")
	}
}