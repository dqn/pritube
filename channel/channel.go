package channel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dqn/pritube/api"
	"github.com/dqn/pritube/util"
)

const endpoint = api.BaseURL + "/channel/%s/about"

// Channel is information of YouTube channel.
type Channel struct {
	ID              string
	Title           string
	Description     string
	Avatars         []Thumbnail
	Banners         []Thumbnail
	ChannelURL      string
	SubscriberCount int
	ViewCount       int
	Country         string
	JoinedDate      time.Time
}

func fetchChannel(channelID string) (*AboutResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpoint, channelID), nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{"accept-language": {"en"}}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rawData := util.GetStringInBetween(string(b), `window["ytInitialData"] = `, ";\n")

	var idr AboutResponse
	if err = json.Unmarshal([]byte(rawData), &idr); err != nil {
		return nil, fmt.Errorf("failed to fetch channel")
	}

	return &idr, nil
}

func parseSubscriberCount(s string) (int, error) {
	var buf bytes.Buffer
	m := 1.

	for _, r := range s {
		switch {
		case (r >= '0' && r <= '9') || r == '.':
			buf.WriteRune(r)
		case r == 'K':
			m = 1_000
		case r == 'M':
			m = 1_000_000
		}
	}

	f, err := strconv.ParseFloat(buf.String(), 64)
	if err != nil {
		return 0, err
	}

	return int(f * m), nil
}

// FetchInfo returns channel infomation.
func FetchInfo(channelID string) (*Channel, error) {
	ar, err := fetchChannel(channelID)
	if err != nil {
		return nil, err
	}

	ci := &Channel{}

	ci.ID = ar.Metadata.ChannelMetadataRenderer.ExternalID
	ci.Title = ar.Metadata.ChannelMetadataRenderer.Title
	ci.Description = ar.Metadata.ChannelMetadataRenderer.Description
	ci.Avatars = ar.Metadata.ChannelMetadataRenderer.Avatar.Thumbnails
	ci.Banners = ar.Header.C4TabbedHeaderRenderer.Banner.Thumbnails
	ci.ChannelURL = ar.Metadata.ChannelMetadataRenderer.ChannelURL

	ci.SubscriberCount, err = parseSubscriberCount(ar.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText)
	if err != nil {
		return nil, err
	}

	for _, tab := range ar.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		if !tab.TabRenderer.Selected {
			continue
		}

		slrContents := tab.TabRenderer.Content.SectionListRenderer.Contents
		if len(slrContents) == 0 {
			return nil, fmt.Errorf("contents of sectionListRenderer not found")
		}

		isrContents := slrContents[0].ItemSectionRenderer.Contents
		if len(slrContents) == 0 {
			return nil, fmt.Errorf("contents of itemSectionRenderer not found")
		}

		isrContent := isrContents[0]
		ci.Country = isrContent.ChannelAboutFullMetadataRenderer.Country.SimpleText

		ci.ViewCount, err = util.RetrieveIntFromDisplayText(isrContent.ChannelAboutFullMetadataRenderer.ViewCountText.SimpleText)
		if err != nil {
			return nil, err
		}

		runs := isrContent.ChannelAboutFullMetadataRenderer.JoinedDateText.Runs
		if len(runs) < 2 {
			return nil, fmt.Errorf("joined date not found")
		}

		ci.JoinedDate, err = time.Parse("Jan 2, 2006", runs[1].Text)
		if err != nil {
			return nil, err
		}
	}

	return ci, nil
}
