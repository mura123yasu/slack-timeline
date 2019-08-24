package command

import (
	"github.com/mura123yasu/slack-timeline/domain"
)

// return chart of aggregated result
func Chart(req domain.SlackRequest) domain.SlackResponse {
	res := domain.SlackResponse{
		Text:         "chart result",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
