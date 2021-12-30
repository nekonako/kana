package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	config, err := readConfig()
	returnIfErr(err)

	bot, err := startBot(config.Token)
	returnIfErr(err)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	<-c
	fmt.Println("nezuko please dont die :(")
	bot.Close()

}

func startBot(token string) (bot *discordgo.Session, err error) {

	bot, err = discordgo.New("Bot " + token)
	returnIfErr(err)

	bot.AddHandler(messagehandlers)

	err = bot.Open()
	returnIfErr(err)

	fmt.Println("bot is running..")
	return

}

func messagehandlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "./code") {

		code := strings.TrimLeft(m.Content, "./code ")

		graphene := newGraphene()
		graphene.setCode(code)

		data, err := graphene.request()
		gomene(err, s, m)

		_, err = s.ChannelFileSend(m.ChannelID, "code.png", bytes.NewBuffer(data))
		gomene(err, s, m)

	}

}

func returnIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func gomene(err error, s *discordgo.Session, m *discordgo.MessageCreate) {
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "gomenasai :(")
	}
}
