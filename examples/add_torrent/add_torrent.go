package main

import (
	"fmt"
	"os"

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

	fmt.Printf("Adding torrent via file..\n")
	err = c.AddTorrentFile("/home/iceman/Downloads/ubuntu-18.04-desktop-amd64.iso.torrent")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
