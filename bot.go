package jubibot


import (
	dsgo "github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	strs "strings"
//	"fmt"
)

type Handler func(s *dsgo.Session,msg *dsgo.MessageCreate);

type Bot struct {
	dg *dsgo.Session
	buser *sgo.User
	botID string
	msgh Handler
	token string
	suf string
	pref string
}

func (this *Bot) init() {
	if ( msgh == nil ) {
		return
	}
	this.dg = dsgo.New("Bot " + this.token)
	dg.AddHandler(msgh)

	u,err := dg.User("@me")

	if err != nil {
		log.Fatal(err)
	}

	this.buser = u

	botID = u.ID

}
func (this *Bot) run() {
	err := dg.Open()

	if err != nil {
		log.Fatal(err)
	}

}

func (this *Bot) end() {
	dg.Close()
}

