package main

import "github.com/Elot322/telegram_Bot_Birthday_v2/src/telegram"

func main() {
	startApp()
}

func startBot() {
	con, _ := telegram.Connect()
	telegram.Updates(con)
}

func startApp() {
	startBot()
}
