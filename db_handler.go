package main

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// DBHandler struct
type DBHandler struct {
	db   *DBHandler
	conf *Config
}

// BackerRecord struct
type BackerRecord struct {
	UserID       string `json:"userid"`
	HashedID     string `json:"hashedid"`
	BackerStatus string `json:"backerstatus"`
	ForumProfile string `json:"forumprofile"`
	ATV          string `json:"atv"`
	PreAlpha     string `json:"prealpha"`
	Alpha        string `json:"alpha"`
	Validated    int    `json:"validated"`
}

// GuildConfig struct
type GuildConfig struct {
	GuildID        string   `json:"guildid"`
	CommandPrefix  string   `json:"commandprefix"`
	BackerCount    int      `json:"backercount"`
	PreAlphaCount  int      `json:"prealphacount"`
	ATVCount       int      `json:"atvcount"`
	ModeratorRoles []string `json:"moderatorroles"`
	UserAccess     bool     `json:"useraccess"`
	UserRoles      []string `json:"userroles"`
}

// TestConnection function
func (h *DBHandler) TestConnection() error {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{h.conf.DBConfig.MongoHost},
		Timeout:  30 * time.Second,
		Database: h.conf.DBConfig.MongoDB,
		Username: h.conf.DBConfig.MongoUser,
		Password: h.conf.DBConfig.MongoPass,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err
	}
	defer session.Close()
	return nil
}

// SaveGuildConfigToDB function
func (h *DBHandler) SaveGuildConfigToDB(guildconfig GuildConfig, session mgo.Session) (err error) {
	c := session.DB(h.conf.DBConfig.MongoDB).C("GuildConfigs")
	_, err = c.UpsertId(guildconfig.GuildID, guildconfig)
	return err
}

// NewGuildConfig function
func (h *DBHandler) NewGuildConfig(guildID string, session mgo.Session) (err error) {
	config := GuildConfig{GuildID: guildID, CommandPrefix: "$", BackerCount: 0, UserAccess: false}
	err = h.SaveGuildConfigToDB(config, session)
	return err
}

// GetRecordFromDB function
func (h *DBHandler) GetRecordFromDB(userid string, c mgo.Collection) (record BackerRecord, err error) {
	userrecord := BackerRecord{}
	err = c.Find(bson.M{"userid": userid}).One(&userrecord)
	return userrecord, err
}

// GetGuildConfig function
func (h *DBHandler) GetGuildConfig(guildID string, session mgo.Session) (config GuildConfig, err error) {

	c := session.DB(h.conf.DBConfig.MongoDB).C("GuildConfigs")
	guildconfig := GuildConfig{}
	err = c.Find(bson.M{"guildid": guildID}).One(&guildconfig)
	return guildconfig, err
}

// GetGuildCP function
func (h *DBHandler) GetGuildCP(guildID string, session mgo.Session) (cp string, err error) {
	c := session.DB(h.conf.DBConfig.MongoDB).C("GuildConfigs")

	var result []struct {
		CP string `bson:"commandprefix"`
	}
	err = c.Find(bson.M{"guildid": guildID}).Select(bson.M{"commandprefix": 1}).All(&result)
	if len(result) < 1 {
		return "$", nil
	}
	return result[0].CP, err
}

// GuildConfigExists function
func (h *DBHandler) GuildConfigExists(guildID string, session mgo.Session) bool {
	record, err := h.GetGuildConfig(guildID, session)
	if err != nil {
		return false
	}

	if record.GuildID == "" {
		return false
	}
	return true
}

// UserHasRecord function
func (h *DBHandler) UserHasRecord(userid string, c mgo.Collection) bool {

	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return false
	}

	if record.UserID == "" {
		return false
	}
	return true
}

// UserValidated function
func (h *DBHandler) UserValidated(userid string, c mgo.Collection) bool {

	if !h.UserHasRecord(userid, c) {
		return false
	}

	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return false
	}
	if record.Validated == 0 {
		return false
	}
	return true
}

// GetBackerStatus function
func (h *DBHandler) GetBackerStatus(userid string, c mgo.Collection) (status string, err error) {
	if !h.UserHasRecord(userid, c) {
		return "", errors.New("Error: No User Record Exists")
	}
	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return "", err
	}
	return record.BackerStatus, nil
}

// GetATVStatus function
func (h *DBHandler) GetATVStatus(userid string, c mgo.Collection) (status string, err error) {
	if !h.UserHasRecord(userid, c) {
		return "", errors.New("Error: No User Record Exists")
	}
	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return "", err
	}
	return record.ATV, nil
}

// GetPreAlphaStatus function
func (h *DBHandler) GetPreAlphaStatus(userid string, c mgo.Collection) (status string, err error) {
	if !h.UserHasRecord(userid, c) {
		return "", errors.New("Error: No User Record Exists")
	}
	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return "", err
	}
	return record.PreAlpha, nil
}

// GetForumProfile function
func (h *DBHandler) GetForumProfile(userid string, c mgo.Collection) (profileurl string, err error) {
	if !h.UserHasRecord(userid, c) {
		return "", errors.New("Error: No User Record Exists")
	}
	record, err := h.GetRecordFromDB(userid, c)
	if err != nil {
		return "", err
	}
	return record.ForumProfile, nil
}
