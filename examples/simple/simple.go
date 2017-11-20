package main

import (
	"fmt"
	"github.com/gmolveau/go-ircevent"
)

func main() {
	ircnick1 := "bot123"
	irccon := irc.IRC(ircnick1, "Bot123")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = false
	irccon.Proxy = "127.0.0.1:9050"
	// irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#chann") })
	irccon.AddCallback("366", func(e *irc.Event) {})
	err := irccon.Connect("testtest123432.onion:6667")
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
