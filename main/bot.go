package main

import (
	"github.com/GianniGM/telebot"
	"strings"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "NameBot",
		ApiKey:  "123456789:AAAAAAAPIKEYHEREEEEEEE",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	//in Go le funzioni sono dei tipi

	//bot.Start (c Configuration, resp Responder)
	//dove type Responder func(telebot.TeleMessage) (string, erro)

	bot.Start(conf, echoFunction)
}

func echoFunction(recMess telebot.TeleMessage) (string, error) {
	var answer string

	user := recMess.From.Uname
	if user == "" {
		user = "JohnDoe"
	} else {
		user = "@" + user
	}
	//parsare la stringa e prendere il comando e il messaggio
	mess := recMess.Text

	message := strings.SplitAfterN(mess, " ", 2)
	cmd := strings.TrimSpace(message[0])

	switch cmd {
	case "/start":
		answer = "Welcome to EchoBot!"
	case "/help":
		answer = "type /echo if wanna use echoBot in chat groups\n"
	case "/echo":
		if len(message) > 1 && message[1] != "" {
			answer = "Hi " + user + " you typed: " + message[1] + "!"
		} else {
			answer = user + "you typed nothing!"
		}
	default: //if no comamnd insert in message
		if mess != "" {
			answer = user + " typed: " + mess + "!"
		} else {
			answer = user + "you typed nothing!"
		}
	}
	return answer, nil
}
