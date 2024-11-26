package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func SendSlackMessage(message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	slackWebhookURL := os.Getenv("SLACK_WEBHOOK")
	if slackWebhookURL == "" {
		return fmt.Errorf("SLACK_WEBHOOK environment variable is not set")
	}

	MakeHTTPRequest(slackWebhookURL, "POST", nil, nil, bytes.NewBuffer(body), "")
	
	return nil
}