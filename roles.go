package main

import (
	"github.com/bwmarrin/discordgo"
)

// Roles struct
type Roles struct {
	conf *Config
}

// CreateAllRoles function
func (h *Roles) CreateAllRoles(s *discordgo.Session, guildID string) (err error) {
	if err = h.CreateATVRole(s, guildID); err != nil {
		////time.Sleep(1*time.Second)(1*time.Second)
		return err
	}
	if err = h.CreateKyriumFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateDiamondFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateEmeraldFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateRubyFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateSapphireFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateGoldFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateSilverFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateBronzeFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateIronFounderRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateContributorSupporterRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateSponsorSupporterRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreatePatronSupporterRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreateATVAuthorizedRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	if err = h.CreatePreAlphaAuthorizedRole(s, guildID); err != nil {
		//time.Sleep(1*time.Second)
		return err
	}
	//time.Sleep(1*time.Second)
	err = h.CreateForumAuthorizedRole(s, guildID)
	return err
}

// ResetAuth function
func (h *Roles) ResetAuth(record BackerRecord, s *discordgo.Session, guildID string) (err error) {

	atvStatus := record.ATV
	prealphaStatus := record.PreAlpha
	backerStatus := record.BackerStatus
	userID := record.UserID

	if backerStatus == "Iron Founder" {
		//time.Sleep(1*time.Second)
		ironRoleID, err := getRoleIDByName(s, guildID, "Iron Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, ironRoleID)

	} else if backerStatus == "Contributor" {
		//time.Sleep(1*time.Second)
		contributorRoleID, err := getRoleIDByName(s, guildID, "Contributor Supporter")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, contributorRoleID)

	} else if backerStatus == "Bronze Founder" {
		//time.Sleep(1*time.Second)
		bronzeRoleID, err := getRoleIDByName(s, guildID, "Bronze Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, bronzeRoleID)

	} else if backerStatus == "Sponsor" {
		//time.Sleep(1*time.Second)
		sponsorRoleID, err := getRoleIDByName(s, guildID, "Sponsor Supporter")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, sponsorRoleID)

	} else if backerStatus == "Silver Founder" {
		//time.Sleep(1*time.Second)
		silverRoleID, err := getRoleIDByName(s, guildID, "Silver Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, silverRoleID)

	} else if backerStatus == "Patron" {
		//time.Sleep(1*time.Second)
		patronRoleID, err := getRoleIDByName(s, guildID, "Patron Supporter")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, patronRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Gold Founder" {
		//time.Sleep(1*time.Second)
		goldRoleID, err := getRoleIDByName(s, guildID, "Gold Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, goldRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Sapphire Founder" {
		//time.Sleep(1*time.Second)
		sapphireRoleID, err := getRoleIDByName(s, guildID, "Sapphire Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, sapphireRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Ruby Founder" {
		//time.Sleep(1*time.Second)
		rubyRoleID, err := getRoleIDByName(s, guildID, "Ruby Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, rubyRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Emerald Founder" {
		//time.Sleep(1*time.Second)
		emeraldRoleID, err := getRoleIDByName(s, guildID, "Emerald Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, emeraldRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Diamond Founder" {
		//time.Sleep(1*time.Second)
		diamondRoleID, err := getRoleIDByName(s, guildID, "Diamond Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, diamondRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Kyrium Founder" {
		//time.Sleep(1*time.Second)
		kyriumRoleID, err := getRoleIDByName(s, guildID, "Kyrium Founder")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, kyriumRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)
	}

	if atvStatus == "true" {
		//time.Sleep(1*time.Second)
		ATVRoleID, err := getRoleIDByName(s, guildID, "Alpha Team Vanguard")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		ATVForumRoleID, err := getRoleIDByName(s, guildID, "ATV Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, ATVRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, ATVForumRoleID)
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)
	}

	if prealphaStatus == "true" {
		//time.Sleep(1*time.Second)
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		//time.Sleep(1*time.Second)
		s.GuildMemberRoleRemove(guildID, userID, preAlphaRoleID)
	}

	//time.Sleep(1*time.Second)
	forumLinkedRoleID, err := getRoleIDByName(s, guildID, "Forum Authorized")
	if err != nil {
		return err
	}
	//time.Sleep(1*time.Second)
	s.GuildMemberRoleRemove(guildID, userID, forumLinkedRoleID)
	return nil
}

// UpdateRoles function
func (h *Roles) UpdateRoles(s *discordgo.Session, record BackerRecord, guildID string) (err error) {

	atvStatus := record.ATV
	prealphaStatus := record.PreAlpha
	backerStatus := record.BackerStatus
	userID := record.UserID

	if backerStatus == "Iron Founder" {
		ironRoleID, err := getRoleIDByName(s, guildID, "Iron Founder")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, ironRoleID)

	} else if backerStatus == "Contributor" {
		contributorRoleID, err := getRoleIDByName(s, guildID, "Contributor Supporter")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, contributorRoleID)

	} else if backerStatus == "Bronze Founder" {
		bronzeRoleID, err := getRoleIDByName(s, guildID, "Bronze Founder")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, bronzeRoleID)

	} else if backerStatus == "Sponsor" {
		sponsorRoleID, err := getRoleIDByName(s, guildID, "Sponsor Supporter")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, sponsorRoleID)

	} else if backerStatus == "Silver Founder" {
		silverRoleID, err := getRoleIDByName(s, guildID, "Silver Founder")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, silverRoleID)

	} else if backerStatus == "Patron" {
		patronRoleID, err := getRoleIDByName(s, guildID, "Patron Supporter")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, patronRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Gold Founder" {
		goldRoleID, err := getRoleIDByName(s, guildID, "Gold Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, goldRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Sapphire Founder" {
		sapphireRoleID, err := getRoleIDByName(s, guildID, "Sapphire Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, sapphireRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Ruby Founder" {
		rubyRoleID, err := getRoleIDByName(s, guildID, "Ruby Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, rubyRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Emerald Founder" {
		emeraldRoleID, err := getRoleIDByName(s, guildID, "Emerald Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, emeraldRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Diamond Founder" {
		diamondRoleID, err := getRoleIDByName(s, guildID, "Diamond Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, diamondRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)

	} else if backerStatus == "Kyrium Founder" {
		kyriumRoleID, err := getRoleIDByName(s, guildID, "Kyrium Founder")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, kyriumRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)
	}

	if atvStatus == "true" {
		ATVRoleID, err := getRoleIDByName(s, guildID, "Alpha Team Vanguard")
		if err != nil {
			return err
		}
		ATVForumRoleID, err := getRoleIDByName(s, guildID, "ATV Authorized")
		if err != nil {
			return err
		}
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, ATVRoleID)
		s.GuildMemberRoleAdd(guildID, userID, ATVForumRoleID)
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)
	}

	if prealphaStatus == "true" {
		preAlphaRoleID, err := getRoleIDByName(s, guildID, "Pre Alpha Authorized")
		if err != nil {
			return err
		}
		s.GuildMemberRoleAdd(guildID, userID, preAlphaRoleID)
	}

	forumLinkedRoleID, err := getRoleIDByName(s, guildID, "Forum Authorized")
	if err != nil {
		return err
	}
	s.GuildMemberRoleAdd(guildID, userID, forumLinkedRoleID)
	return nil
}

/*
Alpha Team Vanguard 1146986
Kyrium Founder 10139859
Diamond Founder 11909306
Emerald Founder 5294200
Ruby Founder 14684511
Sapphire Founder 1004218
Gold Founder 14988288
Silver Founder 12632256
Bronze Founder 13467442
Iron Founder 4410189
Contributor Supporter 15105570
Sponsor Supporter 9807270
Patron Supporter 12098102
ATV Authorized 3066993
Pre Alpha Authorized 9754313
Forum Authorized 0
*/

// CreateATVRole function
func (h *Roles) CreateATVRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Alpha Team Vanguard")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		_, err = s.GuildRoleEdit(guildID, createdrole.ID, "Alpha Team Vanguard", 1146986, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Alpha Team Vanguard", 1146986, false, 0, false)
	return err
}

// CreateKyriumFounderRole function
func (h *Roles) CreateKyriumFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Kyrium Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Kyrium Founder", 10139859, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Kyrium Founder", 10139859, false, 0, false)
	return err
}

// CreateDiamondFounderRole function
func (h *Roles) CreateDiamondFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Diamond Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Diamond Founder", 11909306, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Diamond Founder", 11909306, false, 0, false)
	return err
}

// CreateEmeraldFounderRole function
func (h *Roles) CreateEmeraldFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Diamond Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Emerald Founder", 5294200, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Emerald Founder", 5294200, false, 0, false)
	return err
}

// CreateRubyFounderRole function
func (h *Roles) CreateRubyFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Diamond Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Ruby Founder", 14684511, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Ruby Founder", 14684511, false, 0, false)
	return err
}

// CreateSapphireFounderRole function
func (h *Roles) CreateSapphireFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Sapphire Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Sapphire Founder", 1004218, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Sapphire Founder", 1004218, false, 0, false)
	return err
}

// CreateGoldFounderRole function
func (h *Roles) CreateGoldFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Gold Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Gold Founder", 14988288, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Gold Founder", 14988288, false, 0, false)
	return err
}

// CreateSilverFounderRole function
func (h *Roles) CreateSilverFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Silver Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Silver Founder", 12632256, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Silver Founder", 12632256, false, 0, false)
	return err
}

// CreateBronzeFounderRole function
func (h *Roles) CreateBronzeFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Bronze Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Bronze Founder", 13467442, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Bronze Founder", 13467442, false, 0, false)
	return err
}

// CreateIronFounderRole function
func (h *Roles) CreateIronFounderRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Iron Founder")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Iron Founder", 4410189, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Iron Founder", 4410189, false, 0, false)
	return err
}

// CreateContributorSupporterRole function
func (h *Roles) CreateContributorSupporterRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Contributor Supporter")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Contributor Supporter", 15105570, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Contributor Supporter", 15105570, false, 0, false)
	return err
}

// CreateSponsorSupporterRole function
func (h *Roles) CreateSponsorSupporterRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Sponsor Supporter")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Sponsor Supporter ", 9807270, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Sponsor Supporter ", 9807270, false, 0, false)
	return err
}

// CreatePatronSupporterRole function
func (h *Roles) CreatePatronSupporterRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Patron Supporter")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Patron Supporter", 12098102, false, 0, false)
		return err
	}

	//_, err = s.GuildRoleEdit(guildID, roleID, "Patron Supporter", 12098102, false, 0, false)
	return err
}

// CreateATVAuthorizedRole function
func (h *Roles) CreateATVAuthorizedRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "ATV Authorized")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "ATV Authorized", 3066993, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "ATV Authorized", 3066993, false, 0, false)
	return err
}

// CreatePreAlphaAuthorizedRole function
func (h *Roles) CreatePreAlphaAuthorizedRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Pre Alpha Authorized")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Pre Alpha Authorized", 9754313, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Pre Alpha Authorized", 9754313, false, 0, false)
	return err
}

// CreateForumAuthorizedRole function
func (h *Roles) CreateForumAuthorizedRole(s *discordgo.Session, guildID string) (err error) {
	_, err = getRoleIDByName(s, guildID, "Forum Authorized")
	if err != nil {
		createdrole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return err
		}
		createdrole, err = s.GuildRoleEdit(guildID, createdrole.ID, "Forum Authorized", 0, false, 0, false)
		return err
	}
	//_, err = s.GuildRoleEdit(guildID, roleID, "Forum Authorized", 0, false, 0, false)
	return err
}
