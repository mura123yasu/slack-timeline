package command

import (
	"fmt"

	"github.com/mura123yasu/slack-timeline/domain"
)

// cache task name
func Cache(req domain.SlackRequest) {
	// do nothing
	fmt.Println(req.Text)
}
