package main

import (
	"flag"
	"log"
	"telegabot/clients/events/telegram"
	tgClient "telegabot/clients/telegram"
	eventconsumer "telegabot/consumer/event-consumer"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

// Получение токена и обработка ошибок
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

// получение хоста API telegram
// func mustHost() string {
// 	host := flag.String(
// 		"tg-host",
// 		"",
// 		"host to connect API telegram",
// 	)

// 	flag.Parse()

// 	if *host == "" {
// 		log.Fatal("host is not specified")
// 	}

// 	return *host
// }
