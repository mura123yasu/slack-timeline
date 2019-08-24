package command

import (
	"github.com/mura123yasu/slack-timeline/domain"
)

// cache task name
func Cache(req domain.SlackRequest) domain.SlackResponse {
	res := domain.SlackResponse{
		Text:         "cached result",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
