package telegram

import (
	"flag"
	"log"
)

func mustToken() string {
	token := flag.String("bot-token", "", "Need telegram token for access to telegram bot!")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token not valid!")
	}

	return *token
}
