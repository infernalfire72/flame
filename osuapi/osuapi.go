package osuapi

import (
	"fmt"
	osuapi "github.com/thehowl/go-osuapi"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*osuapi.Client
	web *http.Client
}

func New(key string) *Client {
	return &Client{
		osuapi.NewClient(key),
		&http.Client{},
	}
}

func (c *Client) GetBeatmapContent(file string) []byte {
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/web/maps/"+file, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := c.web.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return data
}
