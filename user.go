package moltencord

import (
	"encoding/json"

	"github.com/MoltenCoreDev/moltencord/utils"
)

// Represents a discord user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot,omitempty"`
	MfaEnabled    bool   `json:"mfa_enabled,omitempty"`
	Banner        string `json:"banner,omitempty"`
	VerifiedEmail bool   `json:"verified,omitempty"`
	NitroType     int    `json:"premium_type,omitempty"`
}

// TODO: implement file support

func (u *User) SendMessage(msg Message) error {
	msg.Recipient = u.ID
	jsonObject, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := utils.MakeRequest("POST", baseUrl+"/users/@me/channels", jsonObject)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return err
}

func (u *User) Send(msg string) error {
	msgObj := Message{Content: msg, Recipient: u.ID}
	jsonObject, err := json.Marshal(msgObj)
	if err != nil {
		return err
	}
	resp, err := utils.MakeRequest("POST", baseUrl+"/users/@me/channels", jsonObject)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return err
}
