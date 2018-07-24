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

	fmt.Printf("Stopping torrent..\n")
	err = c.StopTorrent("001938A83994ACDC9CA57BB44D3B539FDAC90175")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	time.Sleep(10 * time.Second)

	fmt.Printf("Starting torrent..\n")
	err = c.StartTorrent("001938A83994ACDC9CA57BB44D3B539FDAC90175")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	os.Exit(0)
}
