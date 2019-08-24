package command

import (
	"github.com/mura123yasu/slack-timeline/domain"
)

// suggest
func Suggest(req domain.SlackRequest) domain.SlackResponse {
	res := domain.SlackResponse{
		Text:         "suggest result",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
