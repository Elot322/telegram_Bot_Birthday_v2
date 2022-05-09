package telegram

import (
	"log"

	"github.com/Elot322/telegram_Bot_Birthday_v2/lib/e"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	button1 = "Ближайший праздник"
	button2 = "Ближайшие дни рождения"
	button3 = "Дни рождения сотрудников"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(button1),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(button2),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(button3),
	),
)

func Connect() (bot *tgbotapi.BotAPI, err error) {

	defer func() { err = e.WrapIfErr("Can't connect to bot", err) }()

	bot, err = tgbotapi.NewBotAPI(mustToken())

	if err != nil {
		return nil, err
	}

	log.Printf("Connected on account %s", bot.Self.UserName)

	bot.Debug = true

	return bot, nil
}

func Updates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		switch update.Message.Text {
		case "/start":
			text := "Привет. Я телеграм-бот! Я помогу Вам не забыть самые важные праздники - Дни Рождения! Я сам пришлю Вам уведомление в случае ближайшего праздника, а так же напишу день в день, чтобы Вы не забыли поздравить близких Вам людей! Для взаимодействия со мной, используй кнопки на клавиатуре!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)
		case button1:
			return
		case button2:
			return
		case button3:
			return
		default:
			text := "Я не знаю этой комманды!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
