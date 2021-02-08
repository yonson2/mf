package torrent

import (
	"fmt"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"github.com/yonson2/mf/config"
	"github.com/yonson2/mf/open"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func newClient() (*torrent.Client, error) {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = os.TempDir()
	cfg.Debug = false
	cfg.Logger = log.Discard
	c, err := torrent.NewClient(cfg)
	return c, err
}

func getTorrent(c *torrent.Client, url string) (*torrent.Torrent, error) {
	//Check for magnet or if we should download the torrent file.
	if strings.HasPrefix(url, "magnet:") {
		return c.AddMagnet(url)
	}
	//Download the file
	file, err := ioutil.TempFile("", "leetflix-torrent")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	contents, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer contents.Body.Close()
	_, err = io.Copy(file, contents.Body)
	if err != nil {
		return nil, err
	}
	return c.AddTorrentFromFile(file.Name())
}

func getLargestFile(t *torrent.Torrent) *torrent.File {
	var largest *torrent.File
	var maxSize int64

	for _, file := range t.Files() {
		if maxSize < file.Length() {
			maxSize = file.Length()
			largest = file
		}
	}
	return largest
}

func StreamTorrent(tURL string) error {
	// Try to get the player first to avoid performing other operations on error
	player, err := getPlayer()
	if err != nil {
		return err
	}
	//Generate client.
	client, err := newClient()
	if err != nil {
		return err
	}
	//Generate torrent download, depending on magnet or http link
	torrent, err := getTorrent(client, tURL)
	if err != nil {
		return err
	}
	//Wait for torrent info to be available
	<-torrent.GotInfo()
	//Get Biggest file.
	file := getLargestFile(torrent)
	fmt.Println("About to stream file", file.DisplayPath())
	//Get file reader from file.
	fileReader := file.NewReader()
	defer fileReader.Close()
	//Spin up http server and redirect to file reader on request
	videoHost := "127.0.0.1:" + config.HttpPort
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=\""+file.DisplayPath()+"\"")
		http.ServeContent(w, r, file.DisplayPath(), time.Now(), fileReader)
	})
	listener, err := net.Listen("tcp", videoHost)
	if err != nil {
		return err
	}
	go http.Serve(listener, nil)
	//Open player pointing to url
	err = open.RunWith("http://"+videoHost+"/"+file.DisplayPath(), player)
	//Player was closed, delete file
	err = os.Remove(filepath.Join(os.TempDir(), file.Path()))
	err = listener.Close()
	return err
}

func getPlayer() (string, error) {
	path, err := exec.LookPath("mpv")
	if err != nil {
		path, err = exec.LookPath("mplayer")
		if err != nil {
			path, err = exec.LookPath("vlc")
		}
	}
	return path, err
}
