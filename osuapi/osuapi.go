package osuapi

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	Enabled bool
	client  *http.Client
	Key     string
)

func init() {
	client = &http.Client{}
}

func makerq(endpoint string, queryString url.Values) ([]byte, error) {
	queryString.Set("k", Key)
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/"+endpoint+"?"+queryString.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
}