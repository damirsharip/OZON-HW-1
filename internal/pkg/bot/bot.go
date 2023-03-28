package bot

import (
	"fmt"
	"log"

	"HW-1/config"
	commandPkg "HW-1/internal/pkg/bot/command"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type Interface interface {
	Run() error
	RegisterHandler(cmd commandPkg.Interface)
}

var route map[string]CmdHandler

func MustNew() Interface {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		log.Panic(errors.Wrap(err, "init tgbot"))
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &сommander{
		bot:   bot,
		route: make(map[string]commandPkg.Interface),
	}
}

type сommander struct {
	bot   *tgbotapi.BotAPI
	route map[string]commandPkg.Interface
}

// RegisterHandler - not thread-safe
func (c *сommander) RegisterHandler(cmd commandPkg.Interface) {
	c.route[cmd.Name()] = cmd
}

func (c *сommander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmdName := update.Message.Command(); cmdName != "" {
			//switch cmdName {
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
			if cmd, ok := c.route[cmdName]; ok {
				msg.Text = cmd.Process(update.Message.CommandArguments())
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
