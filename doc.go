/*
  Package used to implement the chatPi Bots easily.
  This package is used to implement chatPi bots easily without having to interact with the lower level tcp connections.
  A simple bot can be written by implementin a function which has cpibot.Message as a parameter and returning a string to
  send back. All the broadcasted messages and private messages to the bot written can be handled by the bot. An example bot
  is written down here. This code can be found at : https://github.com/aki237/cpibot-vpnbook


     package main

     import (
  	   "fmt"

  	   "github.com/aki237/cpibot"
       	   "github.com/aki237/vpnbook"
     )

     // The function that runs when a broadcast message or a private message arrives. Key function that drives this bot :P
     // returns string and a boolean for : string to return , whether to send the message or not.
     func returnFunc(msg cpibot.Message) (string, bool) {
  	   if msg.Content.Content == "vpnbook password" {
  		   return vpnbook.GetPassword(), true
  	   }
  	   return "", false
     }

     // Regular main function setup
     func main() {
  	  bot, err := cpibot.NewBot("vpnbook", "192.168.0.100:6672", returnFunc)   //creating a new bot (See Doc for the function)
  	  if err != nil {                                                          //Boring error handling
  		  fmt.Println(err)
  		  return
  	  }
  	  fmt.Println(bot.Run())                                                   //Run the function and Print errors if any
      }

  This is a simple bot that returns vpnbook (a free openvpn service) password using vpnbook (another golang package).
  When the string `vpnbook password` is sent in the broadcast channel or privately to this bot, this bot sends back vpnbook
  password.

  Bots can be as complex as it gets. This bot can be implemented in any language of choice, this is just golang simplification library
*/
package cpibot
