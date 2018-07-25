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

	time.Sleep(5 * time.Second)

	fmt.Printf("Deleting torrent...\n")
	err = c.RemoveTorrent("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	time.Sleep(5 * time.Second)

	fmt.Printf("ReAdding torrent via URL..\n")
	err = c.AddTorrent("http://releases.ubuntu.com/18.04/ubuntu-18.04-desktop-amd64.iso.torrent")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	time.Sleep(10 * time.Second)

	fmt.Printf("Deleting torrent and data...\n")
	err = c.RemoveTorrentAndData("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
