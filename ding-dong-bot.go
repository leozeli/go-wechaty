package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"slices"
	"time"

	"github.com/leozeli/go-wechaty/wechaty-puppet/schemas"
	"github.com/leozeli/go-wechaty/wechaty/user"

	"github.com/leozeli/go-wechaty/sockets"

	"github.com/leozeli/go-wechaty/wechaty"
	"github.com/mdp/qrterminal/v3"
)

func main() {


	var bot = wechaty.NewWechaty(
		
	)
	bot.OnScan(onScan).OnLogin(func(ctx *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnMessage(onMessage).OnLogout(func(ctx *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	bot.DaemonStart()
}

func onMessage(ctx *wechaty.Context, message *user.Message) {
	log.Println(message)
	if message.Type() == schemas.MessageTypeMiniProgram {
		miniApp, err := message.ToMiniProgram()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%v",miniApp.Payload())
	}

	if message.Self() {
		log.Println("Message discarded because its outgoing")
		return
	}

	if message.Age() > 2*60*time.Second {
		log.Println("Message discarded because its TOO OLD(than 2 minutes)")
		return
	}

	if message.Type() != schemas.MessageTypeText ||  !slices.Contains(sockets.GloabConfig.GetKeys(),message.Text()) {
		log.Println("Message discarded because it does not match #ding")
		return
	}
	miniapp := &schemas.MiniProgramPayload{
		AppId:        "wx4f1c1e5f2f674834",
	}
	newMessage := user.NewMiniProgram(miniapp)
	// 1. reply text 'dong'
	// _, err := message.Say(sockets.GloabConfig[message.Text()])
	_, err := message.Say(newMessage)
	if err != nil {
		log.Println(err)
		return
	}
	// miniApp := user.NewMiniProgram(&schemas.MiniProgramPayload{
	// 	AppId:        "wx4f1c1e5f2f674834",
	// })
	// 2. reply image(qrcode image)
	// fileBox := filebox.FromUrl("https://wechaty.github.io/wechaty/images/bot-qr-code.png")
	// _, err = message.Say(fileBox)
	// if err != nil {
		// log.Println(err)
		// return
	// }

	// log.Printf("REPLY with image: %s\n", fileBox)

	// 3. reply url link
	// urlLink := user.NewUrlLink(&schemas.UrlLinkPayload{
	// 	Description:  "Go Wechaty is a Conversational SDK for Chatbot Makers Written in Go",
	// 	ThumbnailUrl: "https://wechaty.js.org/img/icon.png",
	// 	Title:        "wechaty/go-wechaty",
	// 	Url:          "https://github.com/leozeli/go-wechaty",
	// })
	// _, err = message.Say(urlLink)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Printf("REPLY with urlLink: %s\n", urlLink)
}

func onScan(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
	if status == schemas.ScanStatusWaiting || status == schemas.ScanStatusTimeout {
		qrterminal.GenerateHalfBlock(qrCode, qrterminal.L, os.Stdout)

		qrcodeImageUrl := fmt.Sprintf("https://wechaty.js.org/qrcode/%s", url.QueryEscape(qrCode))
		fmt.Printf("onScan: %s - %s\n", status, qrcodeImageUrl)
		return
	}
	fmt.Printf("onScan: %s\n", status)
}
