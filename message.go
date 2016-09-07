package cpibot

import (
	"encoding/base64"
	"encoding/xml"
)

// struct for containing Content node of the reply
type Content struct {
	Type    string `xml:"TYPE,attr"` // Type of the message
	Content string `xml:",chardata"` // Content of the message
}

// struct for containing Message xml of the reply
type Message struct {
	Channel string  `xml:"CHANNEL,attr"` // Channel : private or bradcast
	From    string  `xml:"FROM"`         // Nickname of the sender
	Content Content `xml:"CONTENT"`      // Content struct
}

// GetMessageStruct is used convert the xml sent by the server to native golang structs.
// It returns a Message struct object and an error.
func GetMessageStruct(message string) (Message, error) {
	v := Message{}
	err := xml.Unmarshal([]byte(message), &v)
	if err != nil {
		return Message{}, err
	}
	b, err := base64.StdEncoding.DecodeString(v.Content.Content)
	if err != nil {
		return Message{}, err
	}
	v.Content.Content = string(b)
	return v, nil
}
