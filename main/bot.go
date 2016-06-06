package main

import (
	"github.com/GianniGM/telebot"
	"strings"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "EchoBot",
		ApiKey:  "123456789:AAAAAAAPIKEYHEREEEEEEE",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	//aggiungi un contatore che si incremente ogni volta che facciamo /next e cerchiamo il successivo
	//cerchiamo nell'array di sendgooglesearchrequest(mess, i) quando i arriva al massimo answer prende valore "finished"

	bot.Start(conf, func(recMess telebot.TeleMessage) (string, error) {
		var answer string

		user := recMess.From.Uname
		if user == "" {
			user = "everybody"
		} else {
			user = "@" + user
		}

		//parsare la stringa e prendere il comando e il messaggio
		mess := recMess.Text

		message := strings.SplitAfterN(mess, " ", 2)
		cmd := strings.TrimSpace(message[0])

		switch cmd {
		case "/start":
			answer = "Welcome!\ntype /echo if wanna use echoBot in chat groupstype\n/me reply your username\n/help for this message"
		case "/help":
			answer = "type /echo if wanna use echoBot in chat groups\ntype /hi for greet me"
		case "/hi":
			answer = "Hi " + user + "! :D\n"
		case "/echo":
			if len(message) > 1 && message[1] != "" {
				answer = "Hi " + user + " you typed: " + message[1] + "!"
			} else {
				answer = user + "you typed nothing!"
			}
		default:
			//no comamnd insert in message
			if mess != "" {
				answer = user + " typed: " + mess + "!"
			} else {
				answer = user + "you typed nothing!"
			}
		}
		return answer, nil
	})
}
