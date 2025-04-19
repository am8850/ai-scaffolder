package openai

import (
	"aicoder/pkg/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	client = &http.Client{}
)

func ChatCompletion(messages []config.Message, model string, temperature float64) (string, error) {
	appConfig := config.GetConfig()

	if model == "" {
		model = appConfig.Model
	}

	response_format := "json_object"

	// Create a new payload
	payload := config.ChatRequest{
		Messages:       messages,
		Model:          model,
		Temperature:    temperature,
		ResponseFormat: &config.ChatResponsFormatType{Type: response_format},
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", appConfig.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	if appConfig.Type == "azure" {
		req.Header.Set("api-key", appConfig.Key)
	} else {
		req.Header.Set("Authorization", "Bearer "+appConfig.Key)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Unmarshal the response into a struct
	var response config.ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Print the response message
	return response.Choices[0].Message.Content, nil
}

func DisposeClient() {
	if client != nil {
		client.CloseIdleConnections()
	}
}
