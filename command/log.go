package command

import (
	"github.com/mura123yasu/slack-timeline/domain"
)

// log task
func Log(req domain.SlackRequest) domain.SlackResponse {
	res := domain.SlackResponse{
		Text:         "[slack-timeline]" + req.Text,
		Channel:      req.ChannelName,
		ResponseType: "ephemeral",
	}
	return res
}
