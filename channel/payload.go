package channel

type AboutResponse struct {
	Contents Contents `json:"contents"`
	Header   Header   `json:"header"`
	Metadata Metadata `json:"metadata"`
}

type ViewCountText struct {
	SimpleText string `json:"simpleText"`
}

type Run struct {
	Text string `json:"text"`
}

type JoinedDateText struct {
	Runs []Run `json:"runs"`
}

type Country struct {
	SimpleText string `json:"simpleText"`
}

type ChannelAboutFullMetadataRenderer struct {
	ViewCountText  ViewCountText  `json:"viewCountText"`
	JoinedDateText JoinedDateText `json:"joinedDateText"`
	Country        Country        `json:"country"`
}

type ItemSectionRendererContent struct {
	ChannelAboutFullMetadataRenderer ChannelAboutFullMetadataRenderer `json:"channelAboutFullMetadataRenderer"`
}

type ItemSectionRenderer struct {
	Contents []ItemSectionRendererContent `json:"contents"`
}

type SectionListRendererContent struct {
	ItemSectionRenderer ItemSectionRenderer `json:"itemSectionRenderer"`
}

type SectionListRenderer struct {
	Contents []SectionListRendererContent `json:"contents"`
}

type Content struct {
	SectionListRenderer SectionListRenderer `json:"sectionListRenderer"`
}

type TabRenderer struct {
	Selected bool    `json:"selected"`
	Content  Content `json:"content"`
}

type Tab struct {
	TabRenderer TabRenderer `json:"tabRenderer"`
}

type TwoColumnBrowseResultsRenderer struct {
	Tabs []Tab `json:"tabs"`
}

type Contents struct {
	TwoColumnBrowseResultsRenderer TwoColumnBrowseResultsRenderer `json:"twoColumnBrowseResultsRenderer"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Banner struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type SubscriberCountText struct {
	SimpleText string `json:"simpleText"`
}

type C4TabbedHeaderRenderer struct {
	Banner              Banner              `json:"banner"`
	SubscriberCountText SubscriberCountText `json:"subscriberCountText"`
}

type Header struct {
	C4TabbedHeaderRenderer C4TabbedHeaderRenderer `json:"c4TabbedHeaderRenderer"`
}

type Avatar struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type ChannelMetadataRenderer struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ExternalID  string `json:"externalId"`
	Avatar      Avatar `json:"avatar"`
	ChannelURL  string `json:"channelUrl"`
}

type Metadata struct {
	ChannelMetadataRenderer ChannelMetadataRenderer `json:"channelMetadataRenderer"`
}
