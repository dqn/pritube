package video

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dqn/pritube/api"
)

const endpoint = api.BaseURL + "/get_video_info"

func fetchVideoInfo(videoID string) (url.Values, error) {
	req, err := http.NewRequest("GET", endpoint, nil)

	req.URL.RawQuery = url.Values{
		"video_id": {videoID},
	}.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return url.ParseQuery(string(b))
}

// FetchInfo fetches video info.
func FetchInfo(videoID string) (*PlayerResponse, error) {
	resp, err := fetchVideoInfo(videoID)
	if err != nil {
		return nil, err
	}

	if status, ok := resp["status"]; !ok {
		return nil, fmt.Errorf("failed to fetch video info")
	} else if status[0] != "ok" {
		reason, ok := resp["reason"]

		if ok {
			return nil, fmt.Errorf("fetching video error: " + reason[0])
		}

		return nil, fmt.Errorf("unknown fetch info error: %s", resp)
	}

	j, ok := resp["player_response"]
	if !ok {
		return nil, fmt.Errorf("player_response not found")
	}

	var pr PlayerResponse
	if err := json.Unmarshal([]byte(j[0]), &pr); err != nil {
		return nil, err
	}

	return &pr, nil
}
