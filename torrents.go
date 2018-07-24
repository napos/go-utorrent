package utorrent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TorrentList struct {
	Build        int             `json:"build"`
	RawTorrents  [][]interface{} `json:"torrents"`
	Torrents     []Torrent
	TorrentCache string `json:"torrentc"`
}

type Torrent struct {
	Hash            string `json:"hash"`
	StatusCode      int    `json:"status_code"`
	Name            string `json:"name"`
	Size            int    `json:"size"`
	PercentProgress int    `json:"percent_progress"`
	Downloaded      int    `json:"downloaded"`
	Uploaded        int    `json:"uploaded"`
	Ratio           int    `json:"ratio"`
	UploadSpeed     int    `json:"upload_speed"`
	DownloadSpeed   int    `json:"download_speed"`
	ETA             int    `json:"eta"`
	Label           string `json:"label"`
	PeersConnected  int    `json:"peers_connected"`
	PeersTotal      int    `json:"peers_total"`
	SeedsConnected  int    `json:"seeds_connected"`
	SeedsTotal      int    `json:"seeds_total"`
	Availability    int    `json:"availability"`
	QueueOrder      int    `json:"queue_order"`
	Remaining       int    `json:"remaining"`
	Status          string `json:"status"`
	AddedOn         int    `json:"added_on"`
	CompletedOn     int    `json:"completed_on"`
	FilePath        string `json:"filepath"`
}

func (torrents *TorrentList) UnmarshalJSON(b []byte) error {
	type Alias TorrentList
	rawTorrents := &struct {
		*Alias
	}{
		Alias: (*Alias)(torrents),
	}

	err := json.Unmarshal(b, &rawTorrents)
	if err != nil {
		return err
	}

	for _, torrent := range rawTorrents.RawTorrents {
		torrents.Torrents = append(torrents.Torrents, Torrent{
			Hash:            torrent[0].(string),
			StatusCode:      int(torrent[1].(float64)),
			Name:            torrent[2].(string),
			Size:            int(torrent[3].(float64)),
			PercentProgress: int(torrent[4].(float64)),
			Downloaded:      int(torrent[5].(float64)),
			Uploaded:        int(torrent[6].(float64)),
			Ratio:           int(torrent[7].(float64)),
			UploadSpeed:     int(torrent[8].(float64)),
			DownloadSpeed:   int(torrent[9].(float64)),
			ETA:             int(torrent[10].(float64)),
			Label:           torrent[11].(string),
			PeersConnected:  int(torrent[12].(float64)),
			PeersTotal:      int(torrent[13].(float64)),
			SeedsConnected:  int(torrent[14].(float64)),
			SeedsTotal:      int(torrent[15].(float64)),
			Availability:    int(torrent[16].(float64)),
			QueueOrder:      int(torrent[17].(float64)),
			Remaining:       int(torrent[18].(float64)),
			Status:          torrent[21].(string),
			AddedOn:         int(torrent[23].(float64)),
			CompletedOn:     int(torrent[24].(float64)),
			FilePath:        torrent[26].(string),
		})
	}
	return nil
}

func (c *Client) GetTorrents() ([]Torrent, error) {
	res, err := c.get("/?list=1", nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting torrent list: %s", err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading torrent list: %s", err.Error())
	}
	var torrents TorrentList
	err = json.Unmarshal(body, &torrents)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling torrents: %s", err.Error())
	}

	return torrents.Torrents, err
}
