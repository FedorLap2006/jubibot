package jubibot

import (
	"log"
	"os"
	"os/signal"
//	"strings"
	"syscall"

	hs "github.com/FedorLap2006/jubibot/handlers"
	//dsgo "github.com/bwmarrin/discordgo"
	// dsgo "github.com/bwmarrin/discordgo"
)

const tok string = "NTYyMjgyNzc5ODM3MDcxMzYw.XKNHYQ.MAmzORBVxBQVHa4pZ98Dp-yC8_g"
const bpr string = "!"
const bpf string = ""

var bot Bot

func main() {
	bot.Create()
	bot.Token = tok
	bot.Suf = bpf
	bot.Pref = bpr
	bot.Handlers["new msg"] = msgHandler

	err := bot.Run()
	if err != nil {
		log.Fatal(err)
	}
	defer bot.Close()

	// dg.AddHandler(msgHandler)

	// u, err := dg.User("@me")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// botUser = u
	// if botUser == nil {
	// 	log.Panic("empty")
	// }
	// botID = u.ID

	// err = dg.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer dg.Close()

	log.Printf(`Now running. Press CTRL-C to exit.`)
	log.Printf(`bot username is ` + bot.BUser.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func msgHandler(msg_event *hs.Msg) {
	m := msg_event.MsgEvent
	s := msg_event.Session
	if m.Author.ID == bot.BotID {
		return
	}

	// if !strings.HasPrefix(m.Content, bpr) {
	// 	return
	// }
	// if !strings.HasSuffix(m.Content, bpf) {
	// 	return
	// }

	msg := m.Content[len(bpr) : len(m.Content)-len(bpf)]
	log.Println(msg)
	if msg == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "thank you for watching")
	}
}
