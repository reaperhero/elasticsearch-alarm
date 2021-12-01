package notification

import (
	"github.com/CatchZeng/dingtalk"
	"github.com/sirupsen/logrus"
)

func SendToDinding(title, msg, accessToken string) {
	client := dingtalk.NewClient(accessToken, "")
	sendmsg := dingtalk.NewMarkdownMessage().SetMarkdown(title, msg).SetAt(nil, true)
	if response, err := client.Send(sendmsg); err != nil {
		logrus.Println(err)
		logrus.Println(response.ErrMsg)
	}
	return
}
