// команды бота телеграм
package telegram

import (
	"log"
	"strconv"
	"strings"
	"telegabot/lib/e"
)

const (
	ConvertCmd = "/convert"
	HelpCmd    = "/help"
	StartCmd   = "/start"
	ContactCmd = "/contact"
	OrderCmd   = "/order"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	switch text {
	case ContactCmd:
		return p.sendContact(chatID)
	case OrderCmd:
		return p.sendOrder(chatID)
	case ConvertCmd:
		p.sendConvertInfo(chatID)
		return p.convert(chatID, text)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) convert(chatID int, text string) error {
	sum, err := strconv.Atoi(text)
	if err != nil {
		return e.Wrap("can't convert from string to int", err)
	}
	conv := strconv.Itoa(sum/12) + "рублей"

	return p.tg.SendMessage(chatID, conv)
}

func (p *Processor) sendOrder(chatID int) error {
	return p.tg.SendMessage(chatID, msgOrder)
}

func (p *Processor) sendContact(chatID int) error {
	return p.tg.SendMessage(chatID, msgContact)
}

func (p *Processor) sendConvertInfo(chatID int) error {
	return p.tg.SendMessage(chatID, msgConvertInfo)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
