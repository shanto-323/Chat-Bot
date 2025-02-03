package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Deepseek(query string) (*ChatCompletionResponse, error) {
	apiKey := os.Getenv("c_api")
	baseUrl := "https://openrouter.ai/api/v1/chat/completions"

	if query == "" {
		return nil, fmt.Errorf("string is empty")
	}

	payload := Payload{
		Model: "openai/gpt-3.5-turbo",
		Messages: []Messages{
			{
				Role:    "user",
				Content: query,
			},
		}
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	chatResponse := &ChatCompletionResponse{}
	if err = json.NewDecoder(resp.Body).Decode(chatResponse); err != nil {
		return nil, err
	}
	return chatResponse, nil
}
