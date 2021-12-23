package moltencord

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MoltenCoreDev/moltencord/utils"
)

type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	OwnerID     string `json:"owner_id"`
	// TODO: implement roles

	// Channels are only set after creating a guild. This object will be empty after running Client's GetGuild() method.
	// To fetch channels, use Guild's FetchChannels() method.
	Channels []Channel `json:"channels"`
	// Members are only set after creating a guild. This object will be empty after running Client's GetGuild() method.
	// To fetch Members, use Guild's FetchMembers() method.
	Members []Member `json:"members"`
}

type Member struct {
	User     User   `json:"user"`
	Nickname string `json:"nick"`
	Avatar   string `json:"avatar"`

	// TODO: implement roles

	Deaf bool `json:"deaf"`
	Mute bool `json:"mute"`

	JoinedTimestamp int64 `json:"joined_at"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions int    `json:"permissions"`
	Mentionalbe bool   `json:"mentionable"`
	Icon        string `json:"icon"`
}

// Fetches channels from the guild, and returns them in a slice. It also sets the `Channels` field of the guild
func (g *Guild) FetchChannels() ([]Channel, error) {
	var channels []Channel
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+g.ID+"/channels", []byte{})
	if err != nil {
		return channels, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &channels)

	g.Channels = channels

	return channels, err
}

// Fetches members from the guild, and returns them in a slice. It also sets the `Members` field of the guild
func (g *Guild) FetchMembers() ([]Member, error) {
	var members []Member
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+g.ID+"/members", []byte{})
	if err != nil {
		return members, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &members)

	g.Members = members

	return members, err
}

func (g *Guild) FetchRoles() []Role {
	var roles []Role
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+g.ID+"/roles", []byte{})
	if err != nil {
		return roles
	}
	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &roles)

	return roles
}

func (g *Guild) GetMember(UID string) Member {
	var member Member
	resp, _ := utils.MakeRequest("GET", baseUrl+"/guilds/"+g.ID+"/members/"+UID, []byte{})
	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &member)

	return member
}

func (g *Guild) GetChannel(ID string) (Channel, error) {
	var channel Channel
	resp, err := utils.MakeRequest("GET", baseUrl+"/guilds/"+g.ID+"/channels/"+ID, []byte{})
	if err != nil {
		return channel, err
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(msg, &channel)

	return channel, err
}

// TODO: implement channel creation edits & deletion
func (g *Guild) CreateChannel(c Channel) (Channel, error) {
	msg, err := json.Marshal(c)
	if err != nil {
		return c, err
	}

	resp, err := utils.MakeRequest("POST", baseUrl+"/guilds/"+g.ID+"/channels", msg)
	if err != nil {
		return Channel{}, err
	}
	defer resp.Body.Close()
	return c, nil
}
