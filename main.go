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
	"github.com/dblencowe/CheekyBreekiBot/items"
	"github.com/dblencowe/CheekyBreekiBot/maps"
	"github.com/dblencowe/CheekyBreekiBot/quests"
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
	session.ChannelMessageSend(channelId, fmt.Sprintf("Is Kappa: %t", !tarkovQuest.Nokappa))
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
