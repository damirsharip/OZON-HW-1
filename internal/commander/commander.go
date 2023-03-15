package commander

import (
	"fmt"
	"log"

	"HW-1/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type CmdHandler func(string) string

var route map[string]CmdHandler

var UnknownCommand = errors.New("unknown command")

type Commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]CmdHandler
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Commander{
		bot:   bot,
		route: make(map[string]CmdHandler),
	}, nil
}

func (c *Commander) RegisterHandler(cmd string, f CmdHandler) {
	c.route[cmd] = f
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			//switch cmd {
			//case listCmd:
			//	msg.Text = listFunc()
			//case addCmd:
			//	msg.Text = listFunc()
			//case updateCmd:
			//	msg.Text = listFunc()
			//case deleteCmd:
			//	msg.Text = listFunc()
			//case helpCmd:
			//	msg.Text = helpFunc()
			//default:
			//
			//}
			if f, ok := c.route[cmd]; ok {
				msg.Text = f(update.Message.CommandArguments())
			} else {
				msg.Text = "unknown command"
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)

		}
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
		//if err ==
	}
	return nil
}
