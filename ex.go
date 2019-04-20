package main




import (
	dsgo "github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	strs "strings"
//	"fmt"
)

const tok string = "***********************************************************"
const bpr string = "!"
const bpf string = ""
var botID string
var isPersonal bool = true
var botUser *dsgo.User
func main() {
	dg,err := dsgo.New("Bot " + tok)
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(msgHandler)
	
	u,err := dg.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	botUser = u
	if botUser == nil {
		log.Panic("empty")
	}
	botID = u.ID

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer dg.Close()

	log.Printf(`Now running. Press CTRL-C to exit.`)
	log.Printf(`bot username is ` + botUser.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func msgHandler(s * dsgo.Session,m *dsgo.MessageCreate) {

	if m.Author.ID == botID {
		return 
	}
	
	if !strs.HasPrefix(m.Content,bpr) {
		return
	}
	if !strs.HasSuffix(m.Content,bpf) {
		return
	}
	
	msg := m.Content[len(bpr):len(m.Content)-len(bpf)]
	log.Println(msg)
	if msg == "" {
		_,_ = s.ChannelMessageSend(m.ChannelID, "thank you for watching")
	}
}
