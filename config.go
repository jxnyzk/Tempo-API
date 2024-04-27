package main

import (
	"os"
	"gopkg.in/yaml.v2"
)

type ConfigS struct {
	PubWebhook 	string `yaml:"public_webhook"`
	BotToken 	string `yaml:"bot_token"`
	AppID 		string `yaml:"app_id"`
	GuildID 	string `yaml:"guild_id"`
	UserID 		string `yaml:"user_id"`
}

func LoadConfig() ConfigS {
	var c ConfigS
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}