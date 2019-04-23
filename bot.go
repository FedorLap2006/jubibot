package jubibot

import (
	"log"
	//"fmt"
	hs "github.com/FedorLap2006/jubibot/handlers"

	dsgo "github.com/bwmarrin/discordgo"
	strs "strings"
	"errors"
)

const dSuf string = ""
const dPref string = "!"
const maxHandlers int = 10

type Handler func(s *dsgo.Session, msg *dsgo.MessageCreate)

type Bot struct {
	Dg       *dsgo.Session
	BUser    *dsgo.User
	BotID    string
	Handlers map[string]hs.Handler
	Token    string
	Suf      string
	Pref     string
}

func (this *Bot) uhandler(s *dsgo.Session, msg *dsgo.MessageCreate) {
	if len(this.Pref) != 0 && !strs.HasPrefix(msg.Content,this.Pref){
		return
	}
	if len(this.Suf) != 0 && !strs.HasSuffix(msg.Content,this.Suf){
		return
	}
	if this.Handlers == nil {
		return
	}
	mh := &hs.Msg{s,msg}
	this.Handlers["new msg"](mh)

}

func (this *Bot) Create() {
	this.Handlers = map[string]hs.Handler{} // make(map[string]hs.Handler, maxHandlers)
	this.Dg = nil
	this.BUser = nil
	this.Token = ""
	this.Suf = dSuf
	this.Pref = dPref
}

func (this *Bot) Init() {
	if len(this.Handlers) == 0 {
		return
	}
	if len(this.Token) == 0 {
		return
	}
	this.Dg, _ = dsgo.New("Bot " + this.Token)

	u, err := this.Dg.User("@me")

	if err != nil {
		log.Fatal(err)
	}
	this.BUser = u
	this.BotID = u.ID


	this.Dg.AddHandler(this.uhandler)


}
func (this *Bot) Run() error {
	this.Init()
	if this.Dg == nil {
		return errors.New("dg is nil")
	}
	err := this.Dg.Open()

	return err

}

func (this *Bot) Close() {
	this.Dg.Close()
}
