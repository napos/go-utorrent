package main

import (
	"fmt"
	"os"

	"internal/personal_projects/uTorrentLibrary"
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

	fmt.Printf("Getting torrents..\n")
	torrents, err := c.GetTorrents()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	for _, torrent := range torrents {
		fmt.Printf("Name: %s, Added: %d, Completed: %d, Filepath: %s\n", torrent.Name, torrent.AddedOn, torrent.CompletedOn, torrent.FilePath)
	}

	os.Exit(0)
}
