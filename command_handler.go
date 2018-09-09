package main

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
	"strings"
	"time"
)

// CommandHandler struct
type CommandHandler struct {
	conf  *Config
	db    *DBHandler
	roles *Roles
}

// Read function
func (h *CommandHandler) Read(s *discordgo.Session, m *discordgo.MessageCreate) {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{h.conf.DBConfig.MongoHost},
		Timeout:  30 * time.Second,
		Database: h.conf.DBConfig.MongoDB,
		Username: h.conf.DBConfig.MongoUser,
		Password: h.conf.DBConfig.MongoPass,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Println("Error connecting to DB: " + err.Error())
		return
	}
	defer session.Close()

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	guildID := channel.GuildID

	success, command, payload := SafeInput(s, m, guildID, h.conf, *session, h.db)
	if !success {
		return
	}

	// Used to debug roles - don't uncomment this unless you want a ton of useless logs
	/*
		roles, err := s.GuildRoles(guildID)
		if err != nil {
			return
		}
		for _, role := range roles {
			fmt.Println(role.Name + " " + strconv.Itoa(role.Color))
		}
	*/

	if !h.db.GuildConfigExists(guildID, *session) {
		err = h.db.NewGuildConfig(guildID, *session)
		if err != nil {
			return
		}
	}

	// Done
	if command == "help" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/yamamushi/du-authbot")
		return
	}
	// Done
	if command == "auth" {
		if !h.VerifyUser(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}

		if len(m.Mentions) > 0 {
			if !h.VerifyModerator(guildID, *session, s, m) {
				s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
				return
			}
			fmt.Println(m.Mentions[0].ID)
			err = h.AuthUser(m.Mentions[0].ID, *session, guildID, s)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
				return
			}
			s.ChannelMessageSend(m.ChannelID, m.Mentions[0].Mention()+" has been synchronized")
			return
		}

		err = h.AuthUser(m.Author.ID, *session, guildID, s)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" has been synchronized")
		return
	}
	// Done
	if command == "authall" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
			return
		}

		err = h.AuthAllUsers(*session, guildID, s)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "All users have been successfully synchronized")
		return
	}
	// Done
	if command == "setcp" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
			return
		}
		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}
		if len(payload[0]) > 1 {
			s.ChannelMessageSend(m.ChannelID, "The configured command prefix cannot exceed 1 character.")
			return
		}
		err = h.SetCP(*session, guildID, payload[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Command prefix for this server has been set to: "+payload[0])
		return
	}
	// Done
	if command == "reset" {
		if !h.VerifyUser(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}

		if len(m.Mentions) > 0 {
			if !h.VerifyModerator(guildID, *session, s, m) {
				s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
				return
			}

			fmt.Println(m.Mentions[0].ID)
			err = h.ResetUser(m.Mentions[0].ID, *session, guildID, s)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
				return
			}
			s.ChannelMessageSend(m.ChannelID, m.Mentions[0].Mention()+" has been reset")
			return
		}
		err = h.ResetUser(m.Author.ID, *session, guildID, s)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" has been reset")
		return
	}
	// Done
	if command == "resetall" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
			return
		}
		err = h.ResetAllUsers(*session, guildID, s)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "All user backer roles have been successfully reset")
		return
	}
	// Done
	if command == "addmodrole" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
		}

		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}

		if err = h.AddModRole(concatPayload(payload), guildID, *session, s); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Moderator role added successfully")
		return
	}
	// Done
	if command == "removemodrole" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
		}

		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}

		if err = h.RemoveModRole(concatPayload(payload), guildID, *session, s); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Moderator role removed successfully")
		return
	}
	// Done
	if command == "allowusers" {
		if !h.VerifyModerator(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}

		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}
		option := strings.ToLower(payload[0])
		if option != "false" && option != "true" {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects true or false as an option")
			return
		}
		if err = h.AllowUsers(option, guildID, *session); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}

		if option == "true" {
			s.ChannelMessageSend(m.ChannelID, "User commands enabled")
			return
		}

		s.ChannelMessageSend(m.ChannelID, "User commands disabled")
		return
	}
	// Done
	if command == "adduserrole" {
		if !h.VerifyModerator(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}

		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}

		if err = h.AddUserRole(concatPayload(payload), guildID, *session, s); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "User role added successfully")
		return
	}
	// Done
	if command == "removeuserrole" {
		if !h.VerifyModerator(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}

		if len(payload) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Command "+command+" expects an argument!")
			return
		}

		if err = h.RemoveUserRole(concatPayload(payload), guildID, *session, s); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "User role removed successfully")
		return
	}
	// Done
	if command == "rebuildroles" {
		if !h.VerifyOwner(guildID, s, m) {
			s.ChannelMessageSend(m.ChannelID, "This command can only be run by the server owner.")
			return
		}

		err = h.RebuildRoles(s, guildID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Roles Rebuilt")
		return
	}
	// Done
	if command == "listmodroles" {
		if !h.VerifyModerator(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}
		h.ListModRoles(guildID, *session, s, m)
		return
	}
	// Done
	if command == "listuserroles" {
		if !h.VerifyModerator(guildID, *session, s, m) {
			s.ChannelMessageSend(m.ChannelID, "You are not allowed to use this command!")
			return
		}
		h.ListUserRoles(guildID, *session, s, m)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Unrecognized command: "+command)
	return
}

// VerifyOwner function
func (h *CommandHandler) VerifyOwner(guildID string, s *discordgo.Session, m *discordgo.MessageCreate) bool {

	guild, err := s.Guild(guildID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
		return false
	}

	if guild.OwnerID != m.Author.ID {
		if h.conf.DiscordConfig.BotDevMode {
			if m.Author.ID == h.conf.DiscordConfig.BotDevID {
				return true
			}
			return false
		}
		return false
	}
	return true
}

// GetAllUsers function
func (h *CommandHandler) GetAllUsers(guildID string, s *discordgo.Session) (userList []*discordgo.Member, err error) {

	guild, err := s.Guild(guildID)
	if err != nil {
		return nil, err
	}
	//fmt.Println(strconv.Itoa(len(guild.Members)))
	//for _, member := range guild.Members {
	//	fmt.Println(strconv.Itoa(i) + " - " + member.User.ID)
	//	fmt.Println(member.JoinedAt)
	//	parsedTime, _ := time.Parse(time.RFC3339, member.JoinedAt)
	//	fmt.Println(parsedTime.String())
	//}
	//s.GuildMembers(h.conf.DiscordConfig.GuildID, "", 250)
	return guild.Members, nil
}

// VerifyModerator function
func (h *CommandHandler) VerifyModerator(guildID string, session mgo.Session, s *discordgo.Session, m *discordgo.MessageCreate) bool {

	if h.VerifyOwner(guildID, s, m) {
		return true
	}

	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return false
	}

	member, err := s.GuildMember(guildID, m.Author.ID)
	if err != nil {
		return false
	}

	for _, userRole := range member.Roles {
		if stringInSlice(userRole, guildConfig.ModeratorRoles) {
			return true
		}
	}

	return false
}

// VerifyUser function
func (h *CommandHandler) VerifyUser(guildID string, session mgo.Session, s *discordgo.Session, m *discordgo.MessageCreate) bool {

	if h.VerifyModerator(guildID, session, s, m) {
		return true
	}

	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return false
	}

	if !guildConfig.UserAccess {
		return false
	}

	member, err := s.GuildMember(guildID, m.Author.ID)
	if err != nil {
		return false
	}

	for _, userRole := range member.Roles {
		if stringInSlice(userRole, guildConfig.UserRoles) {
			return true
		}
	}

	return false
}

// AuthUser function
func (h *CommandHandler) AuthUser(userID string, session mgo.Session, guildID string, s *discordgo.Session) error {
	c := session.DB(h.conf.DBConfig.MongoDB).C(h.conf.DBConfig.BackerRecordColumn)
	record, err := h.db.GetRecordFromDB(userID, *c)
	if err != nil {
		return err
	}
	h.roles.UpdateRoles(s, record, guildID)
	return nil
}

// AuthAllUsers function
func (h *CommandHandler) AuthAllUsers(session mgo.Session, guildID string, s *discordgo.Session) error {

	guild, err := s.Guild(guildID)
	if err != nil {
		return err
	}

	c := session.DB(h.conf.DBConfig.MongoDB).C(h.conf.DBConfig.BackerRecordColumn)
	for _, member := range guild.Members {
		record, err := h.db.GetRecordFromDB(member.User.ID, *c)
		if err == nil {
			h.roles.UpdateRoles(s, record, guildID)
		}
	}
	return nil
}

// ResetUser function
func (h *CommandHandler) ResetUser(userID string, session mgo.Session, guildID string, s *discordgo.Session) error {
	c := session.DB(h.conf.DBConfig.MongoDB).C(h.conf.DBConfig.BackerRecordColumn)
	record, err := h.db.GetRecordFromDB(userID, *c)
	if err != nil {
		return err
	}
	fmt.Println(record.UserID)
	return h.roles.ResetAuth(record, s, guildID)
}

// ResetAllUsers function
func (h *CommandHandler) ResetAllUsers(session mgo.Session, guildID string, s *discordgo.Session) error {

	guild, err := s.Guild(guildID)
	if err != nil {
		return err
	}

	c := session.DB(h.conf.DBConfig.MongoDB).C(h.conf.DBConfig.BackerRecordColumn)
	for _, member := range guild.Members {
		record, err := h.db.GetRecordFromDB(member.User.ID, *c)
		if err == nil {
			h.roles.ResetAuth(record, s, guildID)
		}
	}
	return nil
}

// SetCP function
func (h *CommandHandler) SetCP(session mgo.Session, guildID string, prefix string) error {

	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	guildConfig.CommandPrefix = prefix

	return h.db.SaveGuildConfigToDB(guildConfig, session)
}

// RebuildRoles function
func (h *CommandHandler) RebuildRoles(s *discordgo.Session, guildID string) error {
	return h.roles.CreateAllRoles(s, guildID)
}

// AllowUsers function
func (h *CommandHandler) AllowUsers(option string, guildID string, session mgo.Session) error {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	if option == "true" {
		guildConfig.UserAccess = true
	} else {
		guildConfig.UserAccess = false
	}

	return h.db.SaveGuildConfigToDB(guildConfig, session)
}

// AddUserRole function
func (h *CommandHandler) AddUserRole(rolename string, guildID string, session mgo.Session, s *discordgo.Session) error {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		return err
	}

	for _, guildRole := range guildRoles {
		if rolename == guildRole.Name {
			if stringInSlice(rolename, guildConfig.UserRoles) {
				return errors.New("Role already in list")
			}
			guildConfig.UserRoles = append(guildConfig.UserRoles, rolename)
			return h.db.SaveGuildConfigToDB(guildConfig, session)
		}
	}
	return errors.New("Role " + rolename + " does not exist in guild")
}

// AddModRole function
func (h *CommandHandler) AddModRole(rolename string, guildID string, session mgo.Session, s *discordgo.Session) error {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		return err
	}

	for _, guildRole := range guildRoles {
		if rolename == guildRole.Name {
			if stringInSlice(rolename, guildConfig.ModeratorRoles) {
				return errors.New("Role already in list")
			}
			guildConfig.ModeratorRoles = append(guildConfig.ModeratorRoles, rolename)
			return h.db.SaveGuildConfigToDB(guildConfig, session)
		}
	}
	return errors.New("Role " + rolename + " does not exist in guild")
}

// RemoveUserRole function
func (h *CommandHandler) RemoveUserRole(rolename string, guildID string, session mgo.Session, s *discordgo.Session) error {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	if !stringInSlice(rolename, guildConfig.UserRoles) {
		return errors.New("Role is not in list")
	}
	guildConfig.UserRoles = RemoveStringFromSlice(guildConfig.UserRoles, rolename)
	return h.db.SaveGuildConfigToDB(guildConfig, session)
}

// RemoveModRole function
func (h *CommandHandler) RemoveModRole(rolename string, guildID string, session mgo.Session, s *discordgo.Session) error {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		return err
	}

	if !stringInSlice(rolename, guildConfig.ModeratorRoles) {
		return errors.New("Role is not in list")
	}
	guildConfig.ModeratorRoles = RemoveStringFromSlice(guildConfig.ModeratorRoles, rolename)
	return h.db.SaveGuildConfigToDB(guildConfig, session)
}

// ListModRoles function
func (h *CommandHandler) ListModRoles(guildID string, session mgo.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
		return
	}

	output := "Moderator Roles \n```\n"
	for _, role := range guildConfig.ModeratorRoles {
		output += role + "\n"
	}
	output += "```"

	s.ChannelMessageSend(m.ChannelID, output)
	return
}

// ListUserRoles function
func (h *CommandHandler) ListUserRoles(guildID string, session mgo.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	guildConfig, err := h.db.GetGuildConfig(guildID, session)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
		return
	}

	output := "User Roles \n```\n"
	for _, role := range guildConfig.UserRoles {
		output += role + "\n"
	}
	output += "```"

	s.ChannelMessageSend(m.ChannelID, output)
	return
}
