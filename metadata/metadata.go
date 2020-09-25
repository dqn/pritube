package metadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dqn/pritube/api"
	"github.com/dqn/pritube/util"
)

const endpoint = api.BaseURL + "/youtubei/v1/updated_metadata"

// Client for fetching metadata.
type Client struct {
	key      string
	Language string
}

// Metadata is fetched metadata struct.
type Metadata struct {
	ViewCount      string
	ShortViewCount string
	ViewCountInt   int
	IsLive         bool
	LikeCount      string
	DislikeCount   string
	Date           string
	Title          string
	Description    string
}

// New returns new metadata client.
func New() *Client {
	return &Client{
		Language: "en",
	}
}

func (c *Client) updateKey() error {
	resp, err := api.Client.Get(api.BaseURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.key = util.GetStringInBetween(string(b), `"innertubeApiKey":"`, `"`)

	if c.key == "" {
		return fmt.Errorf("failed to update key")
	}

	return nil
}

func (c *Client) fetchMetadata(videoID string) (*metadataResponse, error) {
	b, err := json.Marshal(&metadataRequest{
		Context: context{
			Client: client{
				Hl:            c.Language,
				ClientName:    "WEB",
				ClientVersion: api.ClientVersion,
			},
		},
		VideoID: videoID,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = url.Values{"key": {c.key}}.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r metadataResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	if len(r.Actions) == 0 {
		err = fmt.Errorf("this video id is not available")
		return nil, err
	}

	return &r, err
}
func parseViewerCountString(s string) (int, error) {

	var buf bytes.Buffer

	for _, r := range s {
		if r >= '0' && r <= '9' {
			buf.WriteRune(r)
		}
	}

	return strconv.Atoi(buf.String())
}

// Fetch metadata.
func (c *Client) Fetch(videoID string) (*Metadata, error) {
	if c.key == "" {
		if err := c.updateKey(); err != nil {
			return nil, err
		}
	}

	resp, err := c.fetchMetadata(videoID)
	if err != nil {
		return nil, err
	}

	meta := &Metadata{}

	for _, action := range resp.Actions {
		viewCount := action.UpdateViewershipAction.ViewCount.VideoViewCountRenderer
		toggleButton := action.UpdateToggleButtonTextAction
		date := action.UpdateDateTextAction
		title := action.UpdateTitleAction
		description := action.UpdateDescriptionAction.Description

		if runs := viewCount.ViewCount.Runs; len(runs) != 0 {
			meta.ViewCount = runs[0].Text
			meta.ShortViewCount = viewCount.ExtraShortViewCount.SimpleText
			meta.IsLive = viewCount.IsLive

			count, err := parseViewerCountString(meta.ViewCount)
			if err != nil {
				return nil, err
			}
			meta.ViewCountInt = count

			continue
		}
		if toggleButton.ButtonID == "TOGGLE_BUTTON_ID_TYPE_LIKE" {
			meta.LikeCount = toggleButton.DefaultText.SimpleText
			continue
		}
		if toggleButton.ButtonID == "TOGGLE_BUTTON_ID_TYPE_DISLIKE" {
			meta.DislikeCount = toggleButton.DefaultText.SimpleText
			continue
		}
		if date.DateText.SimpleText != "" {
			meta.Date = date.DateText.SimpleText
			continue
		}
		if title.Title.SimpleText != "" {
			meta.Title = title.Title.SimpleText
			continue
		}
		if runs := description.Runs; len(runs) != 0 {
			var buf bytes.Buffer
			for _, descriptionRun := range runs {
				buf.WriteString(descriptionRun.Text)
			}
			meta.Description += buf.String()
			continue
		}
	}

	return meta, nil
}
