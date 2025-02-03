package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	disGo "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Run() {
	godotenv.Load()
	d_token := os.Getenv("d_api") // Load discord api_key from .env
	const call = "$chi"           // Bot callout in discord

	session, err := disGo.New(d_token) // Create an new session
	if err != nil {
		fmt.Println(err)
		return
	}

	session.AddHandler(func(s *disGo.Session, m *disGo.MessageCreate) {
		if s.State.User.ID == m.Author.ID {
			return
		}
		content := strings.Split(m.Content, " ")
		if content[0] != call {
			fmt.Println("Wrong callout")
			return
		}

		query := ""
		for i := 1; i < len(content); i++ {
			query = query + content[i] + " " // Getting whole query
		}
		s.ChannelMessageSend(m.ChannelID, query)
	})

	session.Identify.Intents = disGo.IntentsAllWithoutPrivileged
	err = session.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer session.Close()

	fmt.Println("bot is online...")

	sc := make(chan os.Signal, 1) // Graceful shutdown
	signal.Notify(sc, syscall.SIGINT, os.Interrupt)
	<-sc
}
