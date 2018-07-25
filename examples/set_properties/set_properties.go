package main

import (
	"fmt"
	"os"
	"time"

	"github.com/naposproject/go-utorrent"
)

func main() {
	c, err := utorrent.NewClient(&utorrent.Client{
		API:      "http://192.168.1.163:8085/gui",
		Username: "admin",
		Password: os.Getenv("TORRENT_PASSWORD"),
	})

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Printf("Adding torrent via URL..\n")
	err = c.AddTorrent("http://releases.ubuntu.com/18.04/ubuntu-18.04-desktop-amd64.iso.torrent")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	// Wait for torrent to be added/started
	time.Sleep(5 * time.Second)

	fmt.Printf("Setting torrent label..\n")
	err = c.SetTorrentLabel("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE", "OS")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent Seed Time..\n")
	err = c.SetTorrentSeedTime("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE", 5)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent Seed Ratio..\n")
	err = c.SetTorrentSeedRatio("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE", 5.2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent queue priority..\n")
	err = c.QueueTop("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent queue priority..\n")
	err = c.QueueDown("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent queue priority..\n")
	err = c.QueueUp("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Setting torrent queue priority..\n")
	err = c.QueueBottom("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
