package controller

import (
	"context"

	"github.com/nikoksr/notify/service/telegram"
	"github.com/nikoksr/notify"
)

func generateNotification(s bool, n, m string){
	telegramService, _ := telegram.New("1848477503:AAH5bTAVcAsKEX1bGtBbpk5aEhvHjM-ooHY")
	telegramService.AddReceivers(-593072923)

	notify.UseServices(telegramService)

	if s == true{
		_ = notify.Send(
			context.Background(),
			"Input New IP Address by " + n,
			"location input in " + m +", waiting to be approved!",
		)
	} else {
		_ = notify.Send(
			context.Background(),
			"IP Address edited by " + n,
			"location input in " + m +", waiting to be approved!",
		)
	}
}