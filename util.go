package main

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
	"strings"
)

// CheckConfig function
func CheckConfig(session *mgo.Session, db *DBHandler, guildID string) {
	if !db.GuildConfigExists(guildID, *session) {
		db.NewGuildConfig(guildID, *session)
	}
}

// RemoveStringFromSlice function
func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// RemoveIntFromSlice function
func RemoveIntFromSlice(s []int, r int) []int {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// SafeInput function
func SafeInput(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, conf *Config, session mgo.Session, db *DBHandler) (success bool, command string, payload []string) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return false, "", []string{""}
	}

	// Ignore bots
	if m.Author.Bot {
		return false, "", []string{""}
	}

	cp, err := db.GetGuildCP(guildID, session)
	if err != nil {
		return false, "", []string{""}
	}

	if !strings.HasPrefix(m.Content, cp) {
		return false, "", []string{""}
	}

	message := strings.Fields(m.Content)
	if len(message) < 1 {
		return false, "", []string{""}
	}

	command, payload = CleanCommand(m.Content, cp)

	return true, command, payload
}

// CleanCommand function
func CleanCommand(input string, cp string) (command string, message []string) {

	// Set our command prefix to the default one within our config file
	message = strings.Fields(input)

	// Remove the prefix from our command
	message[0] = strings.Trim(message[0], cp)
	command = message[0]
	message = RemoveStringFromSlice(message, command)

	return command, message
}

// SplitPayload function
func SplitPayload(input []string) (command string, message []string) {

	// Remove the prefix from our command
	command = input[0]
	message = RemoveStringFromSlice(input, command)

	return command, message

}

// SplitCommandFromArgs function
func SplitCommandFromArgs(input []string) (command string, message string) {

	// Remove the prefix from our command
	command = input[0]
	payload := RemoveStringFromSlice(input, command)

	for _, value := range payload {
		message = message + value + " "
	}
	return command, message
}

// RemoveFromString function
func RemoveFromString(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// CleanChannel function
func CleanChannel(mention string) string {

	mention = strings.TrimPrefix(mention, "<#")
	mention = strings.TrimSuffix(mention, ">")
	return mention

}

// MentionChannel function
func MentionChannel(channelid string, s *discordgo.Session) (mention string, err error) {
	dgchannel, err := s.Channel(channelid)
	if err != nil {
		return "", err
	}

	return "<#" + dgchannel.ID + ">", nil
}

// getRoleIDByName function
func getRoleIDByName(s *discordgo.Session, guildID string, name string) (roleid string, err error) {
	name = strings.Title(name)
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return "", err
	}
	for _, role := range roles {
		if role.Name == name {
			return role.ID, nil
		}
	}
	return "", errors.New("Role ID Not Found: " + name)
}

// stringInSlice functio
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// concatPayload function
func concatPayload(payload []string) string {
	var option string
	for _, word := range payload {
		option += word + " "
	}
	return strings.TrimSpace(option)
}
