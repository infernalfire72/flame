package osuapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	osuapi "github.com/thehowl/go-osuapi"
)

type GetBeatmapsOpts struct {
	BeatmapSetID int
	BeatmapID    int
	IncludeConverted bool
	BeatmapHash      string
	Limit            int
}

const (
	StatusGraveyard ApprovedStatus = iota - 2
	StatusWIP
	StatusPending
	StatusRanked
	StatusApproved
	StatusQualified
	StatusLoved
)

type ApprovedStatus int

// Beatmap is an osu! beatmap.
type Beatmap struct {
	BeatmapSetID      int            `json:"beatmapset_id,string"`
	BeatmapID         int            `json:"beatmap_id,string"`
	Approved          ApprovedStatus `json:"approved,string"`
	TotalLength       int            `json:"total_length,string"`
	HitLength         int            `json:"hit_length,string"`
	DiffName          string         `json:"version"`
	FileMD5           string         `json:"file_md5"`
	CircleSize        float64        `json:"diff_size,string"`
	OverallDifficulty float64        `json:"diff_overall,string"`
	ApproachRate      float64        `json:"diff_approach,string"`
	HPDrain           float64        `json:"diff_drain,string"`
	Mode              osuapi.Mode    `json:"mode,string"`
	ApprovedDate      osuapi.MySQLDate      `json:"approved_date"`
	LastUpdate        osuapi.MySQLDate      `json:"last_update"`
	Artist            string         `json:"artist"`
	Title             string         `json:"title"`
	Creator           string         `json:"creator"`
	BPM               float64        `json:"bpm,string"`
	Source            string         `json:"source"`
	Tags              string         `json:"tags"`
	Genre             osuapi.Genre          `json:"genre_id,string"`
	Language          osuapi.Language       `json:"language_id,string"`
	FavouriteCount    int            `json:"favourite_count,string"`
	Playcount         int            `json:"playcount,string"`
	Passcount         int            `json:"passcount,string"`
	MaxCombo          int            `json:"max_combo,string"`
	DifficultyRating  float64        `json:"difficultyrating,string"`
}

// GetBeatmaps makes a get_beatmaps request to the osu! API.
func GetBeatmaps(opts GetBeatmapsOpts) ([]Beatmap, error) {
	// setup of querystring values
	vals := url.Values{}

	if opts.BeatmapHash != "" {
		vals.Add("h", opts.BeatmapHash)
	}
	if opts.BeatmapID != 0 {
		vals.Add("b", strconv.Itoa(opts.BeatmapID))
	}
	if opts.BeatmapSetID != 0 {
		vals.Add("s", strconv.Itoa(opts.BeatmapSetID))
	}
	if opts.IncludeConverted {
		vals.Add("a", "1")
	}
	if opts.Limit != 0 {
		vals.Add("limit", strconv.Itoa(opts.Limit))
	}

	// actual request
	rawData, err := makerq("get_beatmaps", vals)
	if err != nil {
		return nil, err
	}
	beatmaps := []Beatmap{}
	err = json.Unmarshal(rawData, &beatmaps)
	if err != nil {
		return nil, err
	}
	return beatmaps, nil
}

func GetBeatmapContent(file string) []byte {
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/web/maps/"+file, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
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
