package utils

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

func SendSlackMessage(message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	err = godotenv.Load()
	if err != nil {
		return err
	}

	slackWebhookURL := os.Getenv("SLACK_WEBHOOK")

	MakeHTTPRequest(slackWebhookURL, "POST", nil, nil, bytes.NewBuffer(body), "")
	
	return nil
}