package chvideos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://www.youtube.com"

	channelVideosEndpoint  = baseURL + "/channel/%s/videos"
	subsequentDataEndpoint = baseURL + "/browse_ajax"

	xYoutubeClientName    = "1"
	xYoutubeClientVersion = "2.20200617.02.00"
)

// Chvideos is channel videos fetcher.
type Chvideos struct {
	channelID    string
	client       *http.Client
	continuation string
	itct         string
	IsCompleted  bool
}

// New returns new Chvideos struct.
func New(channelID string) *Chvideos {
	return &Chvideos{
		channelID:    channelID,
		client:       &http.Client{},
		continuation: "",
		itct:         "",
		IsCompleted:  false,
	}
}

func getStringInBetween(str, start, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}

	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}

	return str[s : s+e]
}

func fetchInitialData(client *http.Client, channelID string) (*InitialDataResponse, error) {
	resp, err := client.Get(fmt.Sprintf(channelVideosEndpoint, channelID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	initialData := getStringInBetween(string(b), `window["ytInitialData"] = `, ";\n")
	var idr InitialDataResponse
	if err = json.Unmarshal([]byte(initialData), &idr); err != nil {
		return nil, err
	}

	return &idr, nil
}

// retrieveInitialDataItems returns array of GridVideoRenderer, continuation and itct.
func retrieveInitialDataItems(idr *InitialDataResponse) ([]GridVideoRenderer, string, string) {
	gr := idr.Contents.TwoColumnBrowseResultsRenderer.Tabs[1].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].GridRenderer

	data := gr.Continuations[0].NextContinuationData
	continuation := data.Continuation
	itct := data.ClickTrackingParams

	gvr := make([]GridVideoRenderer, 0, len(gr.Items))
	for _, item := range gr.Items {
		gvr = append(gvr, item.GridVideoRenderer)
	}

	return gvr, continuation, itct
}

func fetchSubsequentData(client *http.Client, continuation string, itct string) (ChannelVideosResponse, error) {
	req, err := http.NewRequest("GET", subsequentDataEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"x-youtube-client-name":    {xYoutubeClientName},
		"x-youtube-client-version": {xYoutubeClientVersion},
	}
	req.URL.RawQuery = url.Values{
		"ctoken":       {continuation},
		"continuation": {continuation},
		"itct":         {itct},
	}.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cvr ChannelVideosResponse
	if err = json.Unmarshal(b, &cvr); err != nil {
		return nil, err
	}

	return cvr, nil
}

// retrieveSubsequentDataItems returns array of GridVideoRenderer, continuation and itct.
func retrieveSubsequentDataItems(cvr ChannelVideosResponse) ([]GridVideoRenderer, string, string) {
	gc := cvr[1].Response.ContinuationContents.GridContinuation
	gvr := make([]GridVideoRenderer, 0, len(gc.Items))
	for _, item := range gc.Items {
		gvr = append(gvr, item.GridVideoRenderer)
	}

	if len(gc.Continuations) < 1 {
		return gvr, "", ""
	}

	data := gc.Continuations[0].NextContinuationData

	continuation := data.Continuation
	itct := data.ClickTrackingParams

	return gvr, continuation, itct
}

// FetchNext returns next page videos.
func (c *Chvideos) FetchNext() ([]GridVideoRenderer, error) {
	if c.IsCompleted {
		return []GridVideoRenderer{}, nil
	}

	if c.continuation == "" {
		idr, err := fetchInitialData(c.client, c.channelID)
		if err != nil {
			return nil, err
		}

		var gvr []GridVideoRenderer
		gvr, c.continuation, c.itct = retrieveInitialDataItems(idr)

		return gvr, nil
	}

	cvr, err := fetchSubsequentData(c.client, c.continuation, c.itct)
	if err != nil {
		return nil, err
	}

	var gvr []GridVideoRenderer
	gvr, c.continuation, c.itct = retrieveSubsequentDataItems(cvr)

	if c.continuation == "" {
		c.IsCompleted = true
		return gvr, nil
	}

	return gvr, nil
}

// FetchAll fetches all channel videos.
func FetchAll(channelID string) ([]GridVideoRenderer, error) {
	client := &http.Client{}

	idr, err := fetchInitialData(client, channelID)
	if err != nil {
		return nil, err
	}

	gvr, continuation, itct := retrieveInitialDataItems(idr)

	for continuation != "" {
		cvr, err := fetchSubsequentData(client, continuation, itct)
		if err != nil {
			return nil, err
		}

		var g []GridVideoRenderer
		g, continuation, itct = retrieveSubsequentDataItems(cvr)
		gvr = append(gvr, g...)

		if continuation == "" {
			break
		}
	}

	return gvr, nil
}
