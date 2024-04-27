package main

import (
	"fmt"
	"log"	
	"time"
	"sort"
	"github.com/bwmarrin/discordgo"
)


type KeyValue struct {
	Key   string
	Value int
}

func StartBot() {

	s, err := discordgo.New("Bot " + Config.BotToken)
	if err != nil {
		log.Fatal("Failed to create bot session:", err)
	}

	_, err = s.ApplicationCommandBulkOverwrite(Config.AppID, Config.GuildID, []*discordgo.ApplicationCommand{
		{
			Name:        "leaderboard",
			Description: "make the leaderboard (Spellman only)",
		},
	})
	if err != nil {
		log.Fatal("Failed to overwrite application commands:", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()

		switch data.Name {
		case "leaderboard":
			if i.Interaction.Member.User.ID == Config.UserID {
				go func() {

				//make a varaible for the content containing the embed like this:
				/*
				
				*/

				

				res, err := s.ChannelMessageSendComplex(i.Interaction.ChannelID, &discordgo.MessageSend{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Tempo Leaderboard",
							Description: "Loading...",
							Color: 0x9f0fed,
							Footer: &discordgo.MessageEmbedFooter{
								Text: "Tempo Stats",
							},
						},
					},
				})

				if err != nil {
					log.Println("Failed to send leaderboard:", err)
				}

				//get the message id

				msg := res.ID

				for ii := 0; true; ii++{
					keys, err := GetAllClaims()
					fmt.Println(ii, "got keys")
					var keyVals []KeyValue
					for key, value := range keys {
						keyVals = append(keyVals, KeyValue{key, value})
					}

					
					sort.SliceStable(keyVals, func(i, j int) bool {
						return keyVals[i].Value > keyVals[j].Value
					})
				
					fmt.Println(ii, "sorted keys")

					var desc string
					for a, kv := range keyVals {
						desc += fmt.Sprintf("**%d.** <@%s> - Claims: **%d**\n", a+1, kv.Key, kv.Value)
						if a == 9 {
							break
						}
					}

					fmt.Println(ii, "made desc")

					Allembed := &discordgo.MessageEmbed{
						Title:       "Tempo Leaderboard",
						Description: desc,
						Color: 0x9f0fed,
						Footer: &discordgo.MessageEmbedFooter{
							Text: "Tempo Stats",
						},
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "https://cdn.discordapp.com/attachments/1140736529011069079/1152725484954722364/3.png",
						},
					}

					keys = DailyClaims 
					keyVals = nil
					for key, value := range keys {
						keyVals = append(keyVals, KeyValue{key, value})
					}

					sort.SliceStable(keyVals, func(i, j int) bool {
						return keyVals[i].Value > keyVals[j].Value
					})

					desc = ""
					for a, kv := range keyVals {
						desc += fmt.Sprintf("**%d.** <@%s> - Claims: **%d**\n", a+1, kv.Key, kv.Value)
						if a == 9 {
							break
						}
					}

					fmt.Println(ii, "made desc2")

					Dailyembed := &discordgo.MessageEmbed{
						Title:       "Tempo Daily Leaderboard",
						Description: desc,
						Color: 0x9f0fed,
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "https://cdn.discordapp.com/attachments/1140736529011069079/1152725484954722364/3.png",
						},
						Footer: &discordgo.MessageEmbedFooter{
							Text: "Tempo Stats",
						},
					}

					ccontent := new(string)
					*ccontent = "Last update ~ <t:" + fmt.Sprint(time.Now().Unix()) + ":R>"

					fmt.Println(ii, "made embeds")

					embedd := &discordgo.MessageEdit{
						Embeds: []*discordgo.MessageEmbed{Allembed, Dailyembed},
						Content: ccontent,
						ID: msg,
						Channel: i.Interaction.ChannelID,
					}

					s.ChannelMessageEditComplex(embedd)
					if err != nil {
						log.Println("Failed to send leaderboard:", err)
					}

					fmt.Println(ii, "sent embeds")

					time.Sleep(1 * time.Minute)

					fmt.Println(ii, "slept")
				}

				}()
				
			} else {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You are not authorized to use this command.",
						Flags:   1 << 6,
					},
				})
				return
			}

		}
	})

	err = s.Open()
	if err != nil {
		log.Fatal("Failed to open bot session:", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	s.Close()
}