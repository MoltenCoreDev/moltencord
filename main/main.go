package main

import (
	"fmt"

	"github.com/MoltenCoreDev/moltencord"
)

func main() {
	c := moltencord.NewClient("Bot OTIyODc0MDQ1NDUwNTUxMzE2.YcHzbg.Yx72x5BfzipGAfG0V2nulcFFCIE")
	defer c.Close()
	guilds, err := c.GetGuilds()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	channel, err := guilds[0].CreateChannel(moltencord.Channel{Name: "test", Type: 0, IsNsfw: true, Topic: "sus"})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("channel: %v\n", channel)
}
