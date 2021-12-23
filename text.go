package moltencord

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MoltenCoreDev/moltencord/utils"
)

// Represents a discord channel
type Channel struct {
	ID string `json:"id"`
	// Use either 0 for guild text channels, 2 for guild voice channels
	// Make sure to check https://discord.com/developers/docs/resources/channel#channel-object-channel-types for more info.
	Type    int    `json:"type"`
	GuildID string `json:"guild_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Topic   string `json:"topic,omitempty"`
	IsNsfw  bool   `json:"nsfw,omitempty"`
	// This field is not set by default. You have to use the FetchGuild() method to fetch it, it will automatically set
	// this field after calling it.
	Guild Guild `json:"-"`
}

// Represents a discord message
type Message struct {
	Content         string     `json:"content,omitempty"`
	Tts             bool       `json:"tts,omitempty"`
	AllowedMentions []struct{} `json:"allowed_mentions,omitempty"`
	// Used to create direct messages. Do not use in guilds.
	Recipient string  `json:"recipient_id,omitempty"`
	Channel   Channel `json:"-"`
}

// TODO: implement file support
func (c *Channel) SendMessage(msg Message) (Message, error) {
	msg.Channel = *c
	jsonObject, err := json.Marshal(msg)
	if err != nil {
		return msg, err
	}
	resp, err := utils.MakeRequest("POST", baseUrl+"/channels/"+c.ID+"/messages", jsonObject)
	if err != nil {
		return msg, err
	}
	resp.Body.Close()
	return msg, err
}

func (c *Channel) Send(msg string) (Message, error) {
	msgObj := Message{Content: msg, Tts: false, Channel: *c}
	jsonObject, err := json.Marshal(msgObj)
	if err != nil {
		return msgObj, err
	}
	resp, err := utils.MakeRequest("POST", baseUrl+"/channels/"+c.ID+"/messages", jsonObject)
	if err != nil {
		return msgObj, err
	}
	resp.Body.Close()
	return msgObj, err
}

func (c *Channel) FetchGuild() Guild {
	var Guild Guild
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+c.GuildID, []byte{})
	if err != nil {
		return Guild
	}
	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &Guild)

	c.Guild = Guild

	return Guild
}
