package moltencord

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/MoltenCoreDev/moltencord/utils"
)

var baseUrl = "https://discordapp.com/api/v9"

type Client struct {
	token string
	User  User
}

func NewClient(token string) *Client {
	c := new(Client)
	utils.SetToken(token)
	return c
}

func (c *Client) Create(token string) error {
	c.token = token
	utils.SetToken(c.token)
	resp, err := utils.MakeRequest("GET", baseUrl+"/users/@me", []byte{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(msg, &c.User)

	return nil
}

func (c *Client) GetChannel(ID string) (Channel, error) {
	var channel Channel
	resp, err := utils.MakeRequest("GET", baseUrl+"/channels/"+ID, []byte{})
	if err != nil {
		return channel, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &channel)

	return channel, err

}

func (c *Client) GetGuild(ID string) (Guild, error) {
	var guild Guild
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+ID, []byte{})
	if err != nil {
		return guild, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &guild)

	return guild, err
}

func (c *Client) GetGuilds() ([]Guild, error) {
	var guilds []Guild
	resp, err := utils.MakeRequest("GET", baseUrl+"/users/@me/guilds", []byte{})
	if err != nil {
		return guilds, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &guilds)

	return guilds, err
}

func (c *Client) Close() {
	os.Exit(0) // Succesful exit
}
