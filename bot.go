package logs

import (
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	sender      *tb.User
	data        string
	message     *tb.Message
	callback    *tb.Callback
	replyOnMess *tb.Message
)

func BotLogger(upd *tb.Update) bool {
	if upd.Message != nil {
		sender = upd.Message.Sender
		replyOnMess = upd.Message.ReplyTo
		message = upd.Message
	}
	if upd.Callback != nil {
		callback = upd.Callback
		sender = upd.Callback.Sender
		data = upd.Callback.Data
		message = upd.Callback.Message
		replyOnMess = message.ReplyTo
	}

	BotLogs.
		WithFields(logrus.Fields{
			"stage":          "incoming message",
			"sender":         sender,
			"data":           data,
			"message":        message,
			"callback":       callback,
			"replyOnMessage": replyOnMess,
		}).
		Infof("incoming message")

	return true
}
