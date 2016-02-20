// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the low level API functions.
package main

import (
	"fmt"
    "regexp"
    "os"
    "strings"
	"github.com/bwmarrin/discordgo"
)

func main() {

	var err error

	// Check for Username and Password CLI arguments.
	if len(os.Args) != 3 {
        fmt.Println("You must provide username and password as arguments. See below example.")
		fmt.Println(os.Args[0], " [username] [password]")
		return
	}

	// Create a new Discord Session interface and set a handler for the
	// OnMessageCreate event that happens for every new message on any channel
	dg := discordgo.Session{}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Login to the Discord server and store the authentication token
	err = dg.Login(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open websocket connection
	err = dg.Open()
	if err != nil {
		fmt.Println(err)
	}

	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}

func isMagnet(m *discordgo.MessageCreate) (bool, string){
    magRE, _ := regexp.Compile("magnet:?.+")
    hashRE, _ := regexp.Compile("([A-Z0-9a-z])+&");
    if magRE.MatchString(m.Content){
        var hash string  = hashRE.FindString(m.Content)
        return true, hash
    }    
    return false, ""
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    var torrentURL string = "http://torcache.net/torrent/HASH.torrent"
    isMagnet, hash := isMagnet(m)
        
    if isMagnet{     
        torrentURL = strings.Replace(torrentURL, "HASH", strings.Trim(strings.ToUpper(hash),"&"), 1)             
        s.ChannelMessageSend(m.ChannelID, torrentURL)
    }
    
}