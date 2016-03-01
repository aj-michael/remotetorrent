package main

import (
	"encoding/base64"
	"github.com/codegangsta/cli"
	"github.com/pusher/pusher-http-go"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "remotetorrent"
	app.Usage = "Send a torrent file to your server to download"
	var appid string
	var appkey string
	var appsecret string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "AppId",
			Value:       "pusher.com app id",
			Destination: &appid,
			EnvVar:      "RT_PUSHER_APP_ID",
		},
		cli.StringFlag{
			Name:        "Key",
			Value:       "pusher.com app key",
			Destination: &appkey,
			EnvVar:      "RT_PUSHER_APP_KEY",
		},
		cli.StringFlag{
			Name:        "Secret",
			Value:       "pusher.com app secret",
			Destination: &appsecret,
			EnvVar:      "RT_PUSHER_APP_SECRET",
		},
	}
	app.Action = func(c *cli.Context) {
		contents, err := ioutil.ReadFile(c.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
		encoded := base64.StdEncoding.EncodeToString(contents)
		client := pusher.Client{
			AppId:  appid,
			Key:    appkey,
			Secret: appsecret,
		}
		id := c.Args()[0] + strconv.FormatInt(time.Now().UnixNano(), 10)
		start := 0
		piece := 1
		chunk := 7500
		pieces := 1 + len(encoded)/chunk
		for start < len(encoded) {
			next := start + chunk
			if next > len(encoded) {
				next = len(encoded)
			}

			substring := encoded[start:next]
			data := map[string]string{
				"id":     id,
				"piece":  strconv.Itoa(piece),
				"pieces": strconv.Itoa(pieces),
				"data":   substring}
			client.Trigger("private-torrentfiles", "client-newtorrent", data)
			start = next
			piece += 1
		}
	}
	app.Run(os.Args)
}
