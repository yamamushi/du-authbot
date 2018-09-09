package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	ConfPath string
)

func init() {
	// Read our command line options
	flag.StringVar(&ConfPath, "c", "du-authbot.conf", "Path to Config File")
	flag.Parse()

	_, err := os.Stat(ConfPath)
	if err != nil {
		log.Fatal("Config file is missing: ", ConfPath)
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	fmt.Println("\n\n|| Starting du-authbot ||")
	log.SetOutput(ioutil.Discard)

	// Verify we can actually read our config file
	conf, err := ReadConfig(ConfPath)
	if err != nil {
		fmt.Println("error reading config file at: ", ConfPath)
		return
	}

	// Run a quick first time db configuration to verify that it is working properly
	fmt.Println("\nChecking Database")
	dbhandler := DBHandler{conf: &conf}
	err = dbhandler.TestConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Database Configuration Verified")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + conf.DiscordConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer dg.Close()

	fmt.Println("\n|| Initializing Main Handler ||")
	handler := PrimaryHandler{db: &dbhandler, conf: &conf, dg: dg}
	err = handler.Init()
	if err != nil {
		fmt.Println("Error in PrimaryHandler.Init(): ", err)
		return
	}
	fmt.Println("\n|| Main Handler Initialized ||")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("\nBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
