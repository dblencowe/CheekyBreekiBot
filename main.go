package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/dblencowe/CheekyBreekiBot/ammunition"
	"github.com/dblencowe/CheekyBreekiBot/helper"
	"github.com/dblencowe/CheekyBreekiBot/items"
	"github.com/dblencowe/CheekyBreekiBot/maps"
	"github.com/dblencowe/CheekyBreekiBot/quests"
	"github.com/dblencowe/CheekyBreekiBot/traders"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	discord.AddHandler(messageCreate)
	discord.Identify.Intents = discordgo.IntentsGuildMessages
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	maps.LoadMaps()
	items.LoadItems()
	quests.LoadQuests()
	ammunition.LoadAmmunition()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content[0:1] != "!" {
		return
	}

	content := strings.ToLower(message.Content)
	parts := strings.Split(content, " ")
	command, arguements := parts[0], parts[1:]

	switch command {
	case "!firecat":
		sendScaredRedGif(session, message.ChannelID)
	case "!map":
		mapSummary(session, message.ChannelID, strings.Join(arguements, " "))
	case "!quest":
		questInfo(session, message.ChannelID, strings.Join(arguements, " "))
	case "!caliber":
		caliberSummary(session, message.ChannelID, strings.Join(arguements, " "))
	default:
		fmt.Println("Unknown command", command, arguements)
	}
}

func mapSummary(session *discordgo.Session, channelId string, mapName string) {
	tarkovMap := maps.GetMap(mapName)
	session.ChannelMessageSend(channelId, tarkovMap.Summary())
}

func questInfo(session *discordgo.Session, channelId string, searchQuery string) {
	tarkovQuest := quests.GetQuest(searchQuery)
	if tarkovQuest == nil {
		session.ChannelMessageSend(channelId, fmt.Sprintf("Sorry, a search for \"%s\" did not return any quests", searchQuery))
		return
	}

	kappaValue := "Yes"
	if tarkovQuest.Nokappa {
		kappaValue = "No"
	}

	var unlocksList []string
	for i := range tarkovQuest.Unlocks {
		id := tarkovQuest.Unlocks[i]
		item := items.GetItemById(id)
		if item == nil {
			continue

		}
		unlocksList = append(unlocksList, item.Name)
	}

	embed := helper.NewEmbed().SetTitle(tarkovQuest.Title).SetURL(tarkovQuest.Wiki).AddField("Given By", traders.GetTraderById(tarkovQuest.Giver).Name).AddField("Required for Kappa?", kappaValue).AddField("Exp. Granted", fmt.Sprintf("%d exp.", tarkovQuest.Exp))
	if len(unlocksList) > 0 {
		embed = embed.AddField("Unlocks", strings.Join(unlocksList, "- \n"))
	}
	// session.ChannelMessageSend(channelId, tarkovQuest.Title)
	session.ChannelMessageSendEmbed(channelId, embed.MessageEmbed)
}

func caliberSummary(session *discordgo.Session, channelId string, term string) {
	caliber := ammunition.GetCaliber(term)
	if len(caliber) == 0 {
		session.ChannelMessageSend(channelId, fmt.Sprintf("Sorry, a search for \"%s\" returned no results.", term))
		session.ChannelMessageSend(channelId, fmt.Sprintf("Available Calibers: %s", strings.Join(ammunition.LoadedCalibers, ", ")))
		return
	}

	ammos := ammunition.GetAmmosByCaliber(caliber)
	var ammoList []string
	for i := range ammos {
		ammoList = append(ammoList, ammos[i].ShortName)
	}

	graph := ammunition.NewAmmoGraph(ammos)

	session.ChannelMessageSend(channelId, fmt.Sprintf("%s: %s", caliber, strings.Join(ammoList, ", ")))
	session.ChannelFileSend(channelId, caliber, graph)
}

func sendScaredRedGif(session *discordgo.Session, channelId string) {
	imageUrl := "https://c.tenor.com/l_LYWdB31iQAAAAC/scared-red.gif"
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error: Can't get image", imageUrl)
		return
	}

	_, err = session.ChannelFileSend(channelId, "scared-red.gif", response.Body)
}
