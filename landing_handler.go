package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
	"time"
)

// LandingHandler struct
type LandingHandler struct {
	conf  *Config
	db    *DBHandler
	roles *Roles
}

// Join function
func (h *LandingHandler) Join(s *discordgo.Session, m *discordgo.GuildCreate) {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{h.conf.DBConfig.MongoHost},
		Timeout:  30 * time.Second,
		Database: h.conf.DBConfig.MongoDB,
		Username: h.conf.DBConfig.MongoUser,
		Password: h.conf.DBConfig.MongoPass,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return
	}
	defer session.Close()

	CheckConfig(session, h.db, m.Guild.ID)
	err = h.roles.CreateAllRoles(s, m.Guild.ID)
	if err != nil {
		fmt.Println(m.Guild.ID + " - Error: " + err.Error())
		return
	}
	return
}

// Leave function
func (h *LandingHandler) Leave(s *discordgo.Session, m *discordgo.GuildDelete) {
	return
}
