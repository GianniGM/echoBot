# EchoBot 

A Telegram bot sample written in go, it uses a fork of [Telebot](https://github.com/cortinico/telebot) from [cortinico](https://github.com/cortinico).

## Usage

When deployed on a server just type your terms to search on Telegram Client and bot will reply your username and message.

## Configuration

1) download my [telebot fork](https://github.com/GianniGM/telebot)

2) create new bot using [botFather](https://telegram.me/BotFather)

3) add your API key and botname

```go

	conf := telebot.Configuration{
		BotName: "YourBotName_bot",
		ApiKey:  "162227600:!!!YOURAPIKEY!!!!BBBBCCCCCCCCCDDDDD"
	}

```

4) build, deploy and have fun! 

## Licence

The following software is released under the [GPL3 License](https://github.com/GianniGM/goSearch/blob/master/LICENSE)