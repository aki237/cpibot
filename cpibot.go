package cpibot

import (
	"bufio"
	"net"
)

// The type template for the Processing function
type ProcessFunc func(Message) (string, bool)

// Struct type Bot
type Bot struct {
	c       net.Conn    // Private  : golang's internal net.Conn interface to read and write to the chat server
	botName string      // Private  : Nickname of the bot
	cookie  string      // Private  : string field to hold the cookie.
	BotFunc ProcessFunc // Exported : The actual ProcessFunc variable which gets executed.
}

// NewBot function returns a Bot struct object. Name of the bot, address of the
// chatPi server (with port) and the text processing function are passed as parameters.
func NewBot(botName string, address string, function ProcessFunc) (*Bot, error) {
	var err error
	c := &Bot{}
	c.c, err = net.Dial("tcp", address)
	if err != nil {
		return &Bot{}, err
	}
	c.botName = botName
	c.BotFunc = function
	return c, err
}

// Run function runs the main program loop. Returns an error if any.
func (b *Bot) Run() error {
	b.c.Write([]byte("JOIN " + b.botName + " samplepassword\n"))
	cookie, err := bufio.NewReader(b.c).ReadString('\n')
	if err != nil {
		return err
	}
	m, err := GetMessageStruct(cookie)
	if err != nil {
		return err
	}
	b.cookie = m.Content.Content
	for {
		process, err := bufio.NewReader(b.c).ReadString('\n')
		if err != nil {
			return err
		}
		v, err := GetMessageStruct(process)
		if err != nil {
			return err
		}
		if v.From != "*ChatPi*" && v.From != b.botName {
			go b.processAndSend(v)
		}
	}
}

// processAndSend : Unexported function which runs and calls the ProcessFunc for each Broadcast message and private message.
func (b *Bot) processAndSend(msg Message) {
	returnstring, shouldreturn := b.BotFunc(msg)
	if shouldreturn {
		if msg.Channel == "private" {
			b.c.Write([]byte("MSG WITH " + b.cookie + " TO " + msg.From + " " + returnstring + "\n"))
		} else {
			b.c.Write([]byte("BROADCAST WITH " + b.cookie + " " + returnstring + "\n"))
		}
	}
}

// Send : send a message to the specified parameter.
func (b *Bot) Send(to, message string) error {
	_, err := b.c.Write([]byte("MSG WITH" + b.cookie + " TO " + to + " " + message + "\n"))
	return err
}
