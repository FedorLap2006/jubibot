package jubibot

import (
	dsgo "github.com/bwmarrin/discordgo"
	"log"
)

type Handler func(s *dsgo.Session, msg *dsgo.MessageCreate)

type Bot struct {
	dg    *dsgo.Session
	buser *dsgo.User
	botID string
	msgh  Handler
	token string
	suf   string
	pref  string
}

func (this *Bot) init() {
	if this.msgh == nil {
		return
	}
	this.dg, _ = dsgo.New("Bot " + this.token)
	this.dg.AddHandler(this.msgh)

	u, err := this.dg.User("@me")

	if err != nil {
		log.Fatal(err)
	}

	this.buser = u

	this.botID = u.ID

}
func (this *Bot) run() {
	err := this.dg.Open()

	if err != nil {
		log.Fatal(err)
	}

}

func (this *Bot) end() {
	this.dg.Close()
}
