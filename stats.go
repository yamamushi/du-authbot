package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
	"time"
)

// StatsCollector struct
type StatsCollector struct {
	db   *DBHandler
	conf *Config
}

// Stats struct
type Stats struct {
	BackerCount   int
	PreAlphaCount int
	ATVCount      int
}

// GatherStats function
func (h *StatsCollector) GatherStats(s *discordgo.Session) {
	for {
		time.Sleep(h.conf.StatsConfig.StatGatheringTimer * time.Minute)
		//fmt.Println("Collecting Stats")
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{h.conf.DBConfig.MongoHost},
			Timeout:  30 * time.Second,
			Database: h.conf.DBConfig.MongoDB,
			Username: h.conf.DBConfig.MongoUser,
			Password: h.conf.DBConfig.MongoPass,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err == nil {
			defer session.Close()

			for _, guild := range s.State.Guilds {
				guildConfig, err := h.db.GetGuildConfig(guild.ID, *session)
				if err == nil {
					go h.ParseUsers(guild, guildConfig, *session, s)

				}
			}
			//fmt.Println("Collection Complete")
		} else {
			fmt.Println("Error connecting to DB: " + err.Error())
		}
	}
}

// ParseUsers function
// To be completed in a follow up commit
func (h *StatsCollector) ParseUsers(guild *discordgo.Guild, guildConfig GuildConfig, session mgo.Session, s *discordgo.Session) {

	ATVRoleID, err := getRoleIDByName(s, guild.ID, "Alpha Team Vanguard")
	if err != nil {
		return
	}
	KyriumRoleID, err := getRoleIDByName(s, guild.ID, "Kyrium Founder")
	if err != nil {
		return
	}
	DiamondRoleID, err := getRoleIDByName(s, guild.ID, "Diamond Founder")
	if err != nil {
		return
	}
	EmeraldRoleID, err := getRoleIDByName(s, guild.ID, "Emerald Founder")
	if err != nil {
		return
	}
	RubyRoleID, err := getRoleIDByName(s, guild.ID, "Ruby Founder")
	if err != nil {
		return
	}
	SapphireRoleID, err := getRoleIDByName(s, guild.ID, "Sapphire Founder")
	if err != nil {
		return
	}
	GoldRoleID, err := getRoleIDByName(s, guild.ID, "Gold Founder")
	if err != nil {
		return
	}
	SilverRoleID, err := getRoleIDByName(s, guild.ID, "Silver Founder")
	if err != nil {
		return
	}
	BronzeRoleID, err := getRoleIDByName(s, guild.ID, "Bronze Founder")
	if err != nil {
		return
	}
	IronRoleID, err := getRoleIDByName(s, guild.ID, "Iron Founder")
	if err != nil {
		return
	}
	ContributorRoleID, err := getRoleIDByName(s, guild.ID, "Contributor Supporter")
	if err != nil {
		return
	}
	SponsorRoleID, err := getRoleIDByName(s, guild.ID, "Sponsor Supporter")
	if err != nil {
		return
	}
	PatronRoleID, err := getRoleIDByName(s, guild.ID, "Patron Supporter")
	if err != nil {
		return
	}

	stats := &Stats{}
	for _, member := range guild.Members {
		for _, memberRole := range member.Roles {
			if memberRole == ATVRoleID {
				stats.ATVCount++
			}
			if memberRole == KyriumRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == DiamondRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == EmeraldRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == RubyRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == SapphireRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == GoldRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
			if memberRole == SilverRoleID {
				stats.BackerCount++
			}
			if memberRole == BronzeRoleID {
				stats.BackerCount++
			}
			if memberRole == IronRoleID {
				stats.BackerCount++
			}
			if memberRole == ContributorRoleID {
				stats.BackerCount++
			}
			if memberRole == SponsorRoleID {
				stats.BackerCount++
			}
			if memberRole == PatronRoleID {
				stats.BackerCount++
				stats.PreAlphaCount++
			}
		}
	}

	guildConfig.BackerCount = stats.BackerCount
	guildConfig.PreAlphaCount = stats.PreAlphaCount
	guildConfig.ATVCount = stats.ATVCount

	h.db.SaveGuildConfigToDB(guildConfig, session)
	//fmt.Println("Guild saved " + guildConfig.GuildID + " - " + strconv.Itoa(guildConfig.BackerCount) + " - " + strconv.Itoa(guildConfig.PreAlphaCount) + " - " + strconv.Itoa(guildConfig.ATVCount))
	return
}
