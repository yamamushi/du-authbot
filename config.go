package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

// Config struct
type Config struct {
	DiscordConfig discordConfig  `toml:"discord"`
	DBConfig      databaseConfig `toml:"database"`
	StatsConfig   statsConfig    `toml:"stats"`
}

// discordConfig struct
type discordConfig struct {
	Token      string `toml:"bot_token"`
	Playing    string `toml:"now_playing"`
	BotDevID   string `toml:"bot_dev_id"`
	BotDevMode bool   `toml:"bot_dev_mode"`
}

// databaseConfig struct
type databaseConfig struct {
	MongoHost          string `toml:"mongohost"`
	MongoDB            string `toml:"mongodb"`
	BackerRecordColumn string `toml:"backerrecordcolumn"`
	MongoUser          string `toml:"mongouser"`
	MongoPass          string `toml:"mongopass"`
}

// statsConfig struct
type statsConfig struct {
	StatGatheringTimer time.Duration `toml:"stat_gathering_timer"`
}

// ReadConfig function
func ReadConfig(path string) (config Config, err error) {

	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Println(err)
		return conf, err
	}

	return conf, nil
}
