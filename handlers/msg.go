package handlers

import (
	dsgo "github.com/bwmarrin/discordgo"
)

type Msg struct {
	Session  *dsgo.Session
	MsgEvent *dsgo.MessageCreate
}

type Handler func(m *Msg)
