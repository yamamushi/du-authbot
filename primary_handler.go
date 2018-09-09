package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// PrimaryHandler struct
type PrimaryHandler struct {
	db   *DBHandler
	conf *Config
	dg   *discordgo.Session
}

// Init function
func (h *PrimaryHandler) Init() error {

	// DO NOT add anything above this line!!
	// Add our main handler -
	//h.dg.AddHandler(h.Read)

	fmt.Println("\nCreating Roles Interface")
	roles := &Roles{conf: h.conf}

	fmt.Println("Creating Landing Handler")
	landingHandler := &LandingHandler{conf: h.conf, db: h.db, roles: roles}
	h.dg.AddHandler(landingHandler.Join)
	h.dg.AddHandler(landingHandler.Leave)

	fmt.Println("Creating Command Handler")
	commandHandler := &CommandHandler{conf: h.conf, db: h.db, roles: roles}
	h.dg.AddHandler(commandHandler.Read)

	fmt.Println("Creating StatsCollector Collector")
	statsCollector := &StatsCollector{conf: h.conf, db: h.db}
	go statsCollector.GatherStats(h.dg) // Spin this off into a goroutine to run every x-minutes (per config file)

	// Open a websocket connection to Discord and begin listening.
	fmt.Println("Opening Connection to Discord")
	err := h.dg.Open()
	h.dg.State.TrackMembers = true
	h.dg.State.TrackChannels = true
	h.dg.State.TrackEmojis = true
	h.dg.State.TrackPresences = true
	h.dg.State.TrackRoles = true
	h.dg.State.TrackVoice = true
	if err != nil {
		fmt.Println("Error Opening Connection: ", err)
		return err
	}
	fmt.Println("Connection Established")

	err = h.PostInit(h.dg)

	if err != nil {
		fmt.Println("Error during Post-Init")
		return err
	}

	return nil

}

// PostInit function
// Just some quick things to run after our websocket has been setup and opened
func (h *PrimaryHandler) PostInit(dg *discordgo.Session) error {
	fmt.Println("Running Post-Init")

	// Update our default playing status
	fmt.Println("Updating Discord Status")
	err := h.dg.UpdateStatus(0, h.conf.DiscordConfig.Playing)
	if err != nil {
		fmt.Println("error updating now playing,", err)
		return err
	}

	fmt.Println("Post-Init Complete")
	return nil
}
