package channel

type aboutResponse struct {
	Contents contents `json:"contents"`
	Header   header   `json:"header"`
	Metadata metadata `json:"metadata"`
}

type viewCountText struct {
	SimpleText string `json:"simpleText"`
}

type run struct {
	Text string `json:"text"`
}

type joinedDateText struct {
	Runs []run `json:"runs"`
}

type country struct {
	SimpleText string `json:"simpleText"`
}

type channelAboutFullMetadataRenderer struct {
	ViewCountText  viewCountText  `json:"viewCountText"`
	JoinedDateText joinedDateText `json:"joinedDateText"`
	Country        country        `json:"country"`
}

type itemSectionRendererContent struct {
	ChannelAboutFullMetadataRenderer channelAboutFullMetadataRenderer `json:"channelAboutFullMetadataRenderer"`
}

type itemSectionRenderer struct {
	Contents []itemSectionRendererContent `json:"contents"`
}

type sectionListRendererContent struct {
	ItemSectionRenderer itemSectionRenderer `json:"itemSectionRenderer"`
}

type sectionListRenderer struct {
	Contents []sectionListRendererContent `json:"contents"`
}

type content struct {
	SectionListRenderer sectionListRenderer `json:"sectionListRenderer"`
}

type tabRenderer struct {
	Selected bool    `json:"selected"`
	Content  content `json:"content"`
}

type Tab struct {
	TabRenderer tabRenderer `json:"tabRenderer"`
}

type twoColumnBrowseResultsRenderer struct {
	Tabs []Tab `json:"tabs"`
}

type contents struct {
	TwoColumnBrowseResultsRenderer twoColumnBrowseResultsRenderer `json:"twoColumnBrowseResultsRenderer"`
}

type thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type banner struct {
	Thumbnails []thumbnail `json:"thumbnails"`
}

type subscriberCountText struct {
	SimpleText string `json:"simpleText"`
}

type c4TabbedHeaderRenderer struct {
	Banner              banner              `json:"banner"`
	SubscriberCountText subscriberCountText `json:"subscriberCountText"`
}

type header struct {
	C4TabbedHeaderRenderer c4TabbedHeaderRenderer `json:"c4TabbedHeaderRenderer"`
}

type avatar struct {
	Thumbnails []thumbnail `json:"thumbnails"`
}

type channelMetadataRenderer struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ExternalID  string `json:"externalId"`
	Avatar      avatar `json:"avatar"`
	ChannelURL  string `json:"channelUrl"`
}

type metadata struct {
	ChannelMetadataRenderer channelMetadataRenderer `json:"channelMetadataRenderer"`
}
