package slackservice

import "fmt"

type MockSlackClient struct {
	LastMessage string
	LastChannel string
}

func (m *MockSlackClient) SendMessage(channelID string, text string) error {
	m.LastChannel = channelID
	m.LastMessage = text
	fmt.Printf("[MOCK] Message sent to %s: %s\n", channelID, text)
	return nil
}
