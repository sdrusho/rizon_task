package slackservice

import "fmt"

type SlackClient interface {
	SendMessage(channelID string, text string) error
}
type SlackService struct {
	Slack SlackClient
}

func (a *SlackService) NotifyUser(userID, msg string) {
	postMessage := fmt.Sprintf("user %s comments: %s", userID, msg)
	_ = a.Slack.SendMessage(userID, postMessage)
}
