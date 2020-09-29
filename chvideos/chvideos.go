package chvideos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dqn/pritube/api"
	"github.com/dqn/pritube/util"
)

const (
	channelVideosEndpoint  = api.BaseURL + "/channel/%s/videos"
	subsequentDataEndpoint = api.BaseURL + "/browse_ajax"

	xYoutubeClientName = "1"
)

// Chvideos is channel videos fetcher.
type Chvideos struct {
	channelID    string
	continuation string
	itct         string
	IsCompleted  bool
}

// New returns new Chvideos struct.
func New(channelID string) *Chvideos {
	return &Chvideos{
		channelID:    channelID,
		continuation: "",
		itct:         "",
		IsCompleted:  false,
	}
}

func fetchInitialData(channelID string) (*InitialDataResponse, error) {
	resp, err := api.Client.Get(fmt.Sprintf(channelVideosEndpoint, channelID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	initialData := util.GetStringInBetween(string(b), `window["ytInitialData"] = `, ";\n")
	var idr InitialDataResponse
	if err = json.Unmarshal([]byte(initialData), &idr); err != nil {
		return nil, fmt.Errorf("failed to fetch videos")
	}

	return &idr, nil
}

// retrieveInitialDataItems returns array of GridVideoRenderer, continuation and itct.
func retrieveInitialDataItems(idr *InitialDataResponse) ([]GridVideoRenderer, string, string) {
	gr := idr.Contents.TwoColumnBrowseResultsRenderer.Tabs[1].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].GridRenderer

	gvr := make([]GridVideoRenderer, 0, len(gr.Items))
	for _, item := range gr.Items {
		gvr = append(gvr, item.GridVideoRenderer)
	}

	if len(gr.Continuations) != 0 {
		data := gr.Continuations[0].NextContinuationData
		continuation := data.Continuation
		itct := data.ClickTrackingParams

		return gvr, continuation, itct
	}

	return gvr, "", ""
}

func fetchSubsequentData(continuation string, itct string) (ChannelVideosResponse, error) {
	req, err := http.NewRequest("GET", subsequentDataEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"x-youtube-client-name":    {xYoutubeClientName},
		"x-youtube-client-version": {api.ClientVersion},
	}
	req.URL.RawQuery = url.Values{
		"ctoken":       {continuation},
		"continuation": {continuation},
		"itct":         {itct},
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
		idr, err := fetchInitialData(c.channelID)
		if err != nil {
			return nil, err
		}

		var gvr []GridVideoRenderer
		gvr, c.continuation, c.itct = retrieveInitialDataItems(idr)

		if c.continuation == "" {
			c.IsCompleted = true
		}

		return gvr, nil
	}

	cvr, err := fetchSubsequentData(c.continuation, c.itct)
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
	idr, err := fetchInitialData(channelID)
	if err != nil {
		return nil, err
	}

	gvr, continuation, itct := retrieveInitialDataItems(idr)

	for continuation != "" {
		cvr, err := fetchSubsequentData(continuation, itct)
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
