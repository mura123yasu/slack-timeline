package command

import (
	"github.com/mura123yasu/slack-timeline/domain"
)

// aggregate today's timelines
func Aggregate(req domain.SlackRequest) domain.SlackResponse {
	res := domain.SlackResponse{
		Text:         "Aggreegated result",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
