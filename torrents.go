package utorrent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// TorrentList is an interface that allows unmarshalling of the
// uTorrent/Bittorrent api into proper golang compatible Torrent structs.
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

// UnmarshallJSON is a custom unmarshaller for torrent lists. Necessary due to
// the fact uTorrent/Bittorrent does not implement a proper json api.
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

// GetTorrents returns a list of Torrent structs containing all of the torrents
// added to the uTorrent/Bittorrent server
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

// GetTorrent gets a specific torrent by info hash
func (c *Client) GetTorrent(hash string) (Torrent, error) {
	torrents, err := c.GetTorrents()
	if err != nil {
		return Torrent{}, err
	}

	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent, nil
		}
	}

	return Torrent{}, fmt.Errorf("Torrent [%s] not found", hash)
}

// PauseTorrent pauses the torrent specified by info hash
func (c *Client) PauseTorrent(hash string) error {
	err := c.action("pause", hash, nil)
	if err != nil {
		return fmt.Errorf("Error pausing torrent: %s", err.Error())
	}

	return nil
}

// UnPauseTorrent unpauses the torrent specified by info hash
func (c *Client) UnPauseTorrent(hash string) error {
	err := c.action("unpause", hash, nil)
	if err != nil {
		return fmt.Errorf("Error unpausing torrent: %s", err.Error())
	}

	return nil
}

// StartTorrent starts the torrent specified by info hash
func (c *Client) StartTorrent(hash string) error {
	err := c.action("start", hash, nil)
	if err != nil {
		return fmt.Errorf("Error starting torrent: %s", err.Error())
	}

	return nil
}

// StopTorrent stops the torrent specified by info hash
func (c *Client) StopTorrent(hash string) error {
	err := c.action("stop", hash, nil)
	if err != nil {
		return fmt.Errorf("Error stopping torrent: %s", err.Error())
	}

	return nil
}

// RecheckTorrent rechecks the torrent specified by info hash
func (c *Client) RecheckTorrent(hash string) error {
	err := c.action("recheck", hash, nil)
	if err != nil {
		return fmt.Errorf("Error rechecking torrent: %s", err.Error())
	}

	return nil
}

// RemoveTorrent removes the torrent specified by info hash
func (c *Client) RemoveTorrent(hash string) error {
	err := c.action("remove", hash, nil)
	if err != nil {
		return fmt.Errorf("Error removing torrent: %s", err.Error())
	}

	return nil
}

// RemoveTorrentAndData removes the torrent and associated data specified by info hash
func (c *Client) RemoveTorrentAndData(hash string) error {
	err := c.action("removedata", hash, nil)
	if err != nil {
		return fmt.Errorf("Error removing torrent and data: %s", err.Error())
	}

	return nil
}

// AddTorrent adds the torrent specified by url or magnet link
func (c *Client) AddTorrent(url string) error {
	res, err := c.get(fmt.Sprintf("/?action=add-url&s=%s", url), nil)
	if err != nil {
		return fmt.Errorf("Error adding torrent: %s", err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Error adding torrent: status code: %d", res.StatusCode)
	}

	return nil
}

// AddTorrentFile adds the torrent specified by a file on disk
func (c *Client) AddTorrentFile(torrentpath string) error {
	file, err := os.Open(torrentpath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("torrent_file", filepath.Base(torrentpath))
	if err != nil {
		return fmt.Errorf("Error adding torrent: %s", err)
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("Error adding torrent: %s", err)
	}

	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	res, err := c.post("/?action=add-file", body.Bytes(), &header)
	if err != nil {
		return fmt.Errorf("Error adding torrent: %s", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Error adding torrent: status code: %d", res.StatusCode)
	}

	return nil
}

// SetTorrentProperty sets a property for the given torrent.
func (c *Client) SetTorrentProperty(hash string, property string, value string) error {
	res, err := c.get(fmt.Sprintf("/?action=setprops&hash=%s&s=%s&v=%s", hash, property, value), nil)
	if err != nil {
		return fmt.Errorf("Error setting torrent (%s) '%s' to '%s': %s ", hash, property, value, err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Error setting torrent (%s) '%s' to '%s' - status code: %d ", hash, property, value, res.StatusCode)
	}

	return nil
}

// SetTorrentLabel sets the label for the given torrent
func (c *Client) SetTorrentLabel(hash string, label string) error {
	err := c.SetTorrentProperty(hash, "label", label)
	if err != nil {
		return err
	}

	return nil
}

// SetTorrentSeedRatio sets the seed ratio for the given torrent
func (c *Client) SetTorrentSeedRatio(hash string, ratio float64) error {
	err := c.SetTorrentProperty(hash, "seed_override", "1")
	if err != nil {
		return err
	}

	err = c.SetTorrentProperty(hash, "seed_ratio", strconv.FormatFloat(ratio*10, 'f', 0, 64))
	if err != nil {
		return err
	}

	return nil
}

// SetTorrentSeedTime sets the seed time for the given torrent
func (c *Client) SetTorrentSeedTime(hash string, time int) error {
	err := c.SetTorrentProperty(hash, "seed_override", "1")
	if err != nil {
		return err
	}

	err = c.SetTorrentProperty(hash, "seed_time", strconv.FormatInt(int64(time*3600), 10))
	if err != nil {
		return err
	}

	return nil
}

// QueueTop sends the torrent to the top of the download queue
func (c *Client) QueueTop(hash string) error {
	err := c.action("queuetop", hash, nil)
	if err != nil {
		return fmt.Errorf("Error setting torrent queue priority: %s", err.Error())
	}

	return nil
}

// QueueUp moves the torrent up the download queue
func (c *Client) QueueUp(hash string) error {
	err := c.action("queueup", hash, nil)
	if err != nil {
		return fmt.Errorf("Error setting torrent queue priority: %s", err.Error())
	}

	return nil
}

// QueueUp moves the torrent down the download queue
func (c *Client) QueueDown(hash string) error {
	err := c.action("queuedown", hash, nil)
	if err != nil {
		return fmt.Errorf("Error setting torrent queue priority: %s", err.Error())
	}

	return nil
}

// QueueTop sends the torrent to the bottom of the download queue
func (c *Client) QueueBottom(hash string) error {
	err := c.action("queuebottom", hash, nil)
	if err != nil {
		return fmt.Errorf("Error setting torrent queue priority: %s", err.Error())
	}

	return nil
}
