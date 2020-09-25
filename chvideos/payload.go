package chvideos

type ChannelVideosResponse []struct {
	Page      string   `json:"page"`
	XSRFToken string   `json:"xsrf_token,omitempty"`
	Response  Response `json:"response,omitempty"`
	Endpoint  Endpoint `json:"endpoint,omitempty"`
	Timing    Timing   `json:"timing,omitempty"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"`
}

type YtConfigData struct {
	Csn         string `json:"csn"`
	VisitorData string `json:"visitorData"`
}

type WebResponseContextExtensionData struct {
	YtConfigData YtConfigData `json:"ytConfigData"`
	HasDecorated bool         `json:"hasDecorated"`
}

type ResponseContext struct {
	ServiceTrackingParams           []ServiceTrackingParam          `json:"serviceTrackingParams"`
	MaxAgeSeconds                   int                             `json:"maxAgeSeconds"`
	WebResponseContextExtensionData WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type ThumbnailItem struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnail struct {
	Thumbnails []ThumbnailItem `json:"thumbnails"`
}

type AccessibilityData struct {
	Label string `json:"label"`
}

type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type Title struct {
	Accessibility Accessibility `json:"accessibility"`
	SimpleText    string        `json:"simpleText"`
}

type PublishedTimeText struct {
	SimpleText string `json:"simpleText"`
}

type ViewCountText struct {
	SimpleText string `json:"simpleText"`
}

type WebCommandMetadata struct {
	URL         string `json:"url"`
	WebPageType string `json:"webPageType"`
	RootVe      int    `json:"rootVe"`
}

type CommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}

type WatchEndpoint struct {
	VideoID string `json:"videoId"`
}

type NavigationEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	WatchEndpoint       WatchEndpoint   `json:"watchEndpoint"`
}

type ShortViewCountText struct {
	SimpleText string `json:"simpleText"`
}

type Run struct {
	Text string `json:"text"`
}

type MenuServiceItemRendererText struct {
	Runs []Run `json:"runs"`
}

type Icon struct {
	IconType string `json:"iconType"`
}

type CreatePlaylistServiceEndpoint struct {
	VideoIds []string `json:"videoIds"`
	Hack     bool     `json:"hack"`
	Params   string   `json:"params"`
}

type OnCreateListCommand struct {
	ClickTrackingParams           string                        `json:"clickTrackingParams"`
	CommandMetadata               CommandMetadata               `json:"commandMetadata"`
	CreatePlaylistServiceEndpoint CreatePlaylistServiceEndpoint `json:"createPlaylistServiceEndpoint"`
}

type AddToPlaylistCommand struct {
	OpenMiniplayer      bool                `json:"openMiniplayer"`
	OpenListPanel       bool                `json:"openListPanel"`
	VideoID             string              `json:"videoId"`
	ListType            string              `json:"listType"`
	OnCreateListCommand OnCreateListCommand `json:"onCreateListCommand"`
	VideoIds            []string            `json:"videoIds"`
}

type SignalServiceEndpointAction struct {
	AddToPlaylistCommand AddToPlaylistCommand `json:"addToPlaylistCommand"`
}

type SignalServiceEndpoint struct {
	Signal  string                        `json:"signal"`
	Actions []SignalServiceEndpointAction `json:"actions"`
}

type ServiceEndpoint struct {
	ClickTrackingParams   string                `json:"clickTrackingParams"`
	CommandMetadata       CommandMetadata       `json:"commandMetadata"`
	SignalServiceEndpoint SignalServiceEndpoint `json:"signalServiceEndpoint"`
}

type MenuServiceItemRenderer struct {
	Text            MenuServiceItemRendererText `json:"text"`
	Icon            Icon                        `json:"icon"`
	ServiceEndpoint ServiceEndpoint             `json:"serviceEndpoint"`
	TrackingParams  string                      `json:"trackingParams"`
}

type MenuRendererItem struct {
	MenuServiceItemRenderer MenuServiceItemRenderer `json:"menuServiceItemRenderer"`
}

type MenuRenderer struct {
	Items          []MenuRendererItem `json:"items"`
	TrackingParams string             `json:"trackingParams"`
	Accessibility  Accessibility      `json:"accessibility"`
}

type Menu struct {
	MenuRenderer MenuRenderer `json:"menuRenderer"`
}

type RendererText struct {
	Accessibility Accessibility `json:"accessibility"`
	SimpleText    string        `json:"simpleText"`
}

type ThumbnailOverlayTimeStatusRenderer struct {
	Text  RendererText `json:"text"`
	Style string       `json:"style"`
}

type UntoggledIcon struct {
	IconType string `json:"iconType"`
}

type ToggledIcon struct {
	IconType string `json:"iconType"`
}

type PlaylistEditEndpointAction struct {
	AddedVideoID string `json:"addedVideoId"`
	Action       string `json:"action"`
}

type PlaylistEditEndpoint struct {
	PlaylistID string                       `json:"playlistId"`
	Actions    []PlaylistEditEndpointAction `json:"actions"`
}

type UntoggledServiceEndpoint struct {
	ClickTrackingParams   string                `json:"clickTrackingParams"`
	CommandMetadata       CommandMetadata       `json:"commandMetadata"`
	PlaylistEditEndpoint  PlaylistEditEndpoint  `json:"playlistEditEndpoint"`
	SignalServiceEndpoint SignalServiceEndpoint `json:"signalServiceEndpoint"`
}

type ToggledServiceEndpoint struct {
	ClickTrackingParams  string               `json:"clickTrackingParams"`
	CommandMetadata      CommandMetadata      `json:"commandMetadata"`
	PlaylistEditEndpoint PlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type UntoggledAccessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type ToggledAccessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type ThumbnailOverlayToggleButtonRenderer struct {
	IsToggled                bool                     `json:"isToggled"`
	UntoggledIcon            UntoggledIcon            `json:"untoggledIcon"`
	ToggledIcon              ToggledIcon              `json:"toggledIcon"`
	UntoggledTooltip         string                   `json:"untoggledTooltip"`
	ToggledTooltip           string                   `json:"toggledTooltip"`
	UntoggledServiceEndpoint UntoggledServiceEndpoint `json:"untoggledServiceEndpoint"`
	ToggledServiceEndpoint   ToggledServiceEndpoint   `json:"toggledServiceEndpoint"`
	UntoggledAccessibility   UntoggledAccessibility   `json:"untoggledAccessibility"`
	ToggledAccessibility     ToggledAccessibility     `json:"toggledAccessibility"`
	TrackingParams           string                   `json:"trackingParams"`
}

type ThumbnailOverlayNowPlayingRendererText struct {
	Runs []Run `json:"runs"`
}

type ThumbnailOverlayNowPlayingRenderer struct {
	Text ThumbnailOverlayNowPlayingRendererText `json:"text"`
}

type ThumbnailOverlay struct {
	ThumbnailOverlayTimeStatusRenderer   ThumbnailOverlayTimeStatusRenderer   `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
	ThumbnailOverlayToggleButtonRenderer ThumbnailOverlayToggleButtonRenderer `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
	ThumbnailOverlayNowPlayingRenderer   ThumbnailOverlayNowPlayingRenderer   `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
}

type GridVideoRenderer struct {
	VideoID            string             `json:"videoId"`
	Thumbnail          Thumbnail          `json:"thumbnail"`
	Title              Title              `json:"title"`
	PublishedTimeText  PublishedTimeText  `json:"publishedTimeText"`
	ViewCountText      ViewCountText      `json:"viewCountText"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string             `json:"trackingParams"`
	ShortViewCountText ShortViewCountText `json:"shortViewCountText"`
	Menu               Menu               `json:"menu"`
	ThumbnailOverlays  []ThumbnailOverlay `json:"thumbnailOverlays"`
}

type GridContinuationItem struct {
	GridVideoRenderer GridVideoRenderer `json:"gridVideoRenderer"`
}

type NextContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
}

type Continuation struct {
	NextContinuationData NextContinuationData `json:"nextContinuationData"`
}

type GridContinuation struct {
	Items          []GridContinuationItem `json:"items"`
	Continuations  []Continuation         `json:"continuations"`
	TrackingParams string                 `json:"trackingParams"`
}

type ContinuationContents struct {
	GridContinuation GridContinuation `json:"gridContinuation"`
}

type Avatar struct {
	Thumbnails []ThumbnailItem `json:"thumbnails"`
}

type ChannelMetadataRenderer struct {
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
	PlusPageLink           string   `json:"plusPageLink"`
	ExternalID             string   `json:"externalId"`
	Keywords               string   `json:"keywords"`
	OwnerUrls              []string `json:"ownerUrls"`
	Avatar                 Avatar   `json:"avatar"`
	ChannelURL             string   `json:"channelUrl"`
	IsFamilySafe           bool     `json:"isFamilySafe"`
	AvailableCountryCodes  []string `json:"availableCountryCodes"`
	AndroidDeepLink        string   `json:"androidDeepLink"`
	AndroidAppindexingLink string   `json:"androidAppindexingLink"`
	IosAppindexingLink     string   `json:"iosAppindexingLink"`
	TabPath                string   `json:"tabPath"`
	VanityChannelURL       string   `json:"vanityChannelUrl"`
}

type Metadata struct {
	ChannelMetadataRenderer ChannelMetadataRenderer `json:"channelMetadataRenderer"`
}

type LinkAlternate struct {
	HrefURL string `json:"hrefUrl"`
}

type MicroformatDataRenderer struct {
	URLCanonical       string          `json:"urlCanonical"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Thumbnail          Thumbnail       `json:"thumbnail"`
	SiteName           string          `json:"siteName"`
	AppName            string          `json:"appName"`
	AndroidPackage     string          `json:"androidPackage"`
	IosAppStoreID      string          `json:"iosAppStoreId"`
	IosAppArguments    string          `json:"iosAppArguments"`
	OgType             string          `json:"ogType"`
	URLApplinksWeb     string          `json:"urlApplinksWeb"`
	URLApplinksIos     string          `json:"urlApplinksIos"`
	URLApplinksAndroid string          `json:"urlApplinksAndroid"`
	URLTwitterIos      string          `json:"urlTwitterIos"`
	URLTwitterAndroid  string          `json:"urlTwitterAndroid"`
	TwitterCardType    string          `json:"twitterCardType"`
	TwitterSiteHandle  string          `json:"twitterSiteHandle"`
	SchemaDotOrgType   string          `json:"schemaDotOrgType"`
	Noindex            bool            `json:"noindex"`
	Unlisted           bool            `json:"unlisted"`
	Tags               []string        `json:"tags"`
	LinkAlternates     []LinkAlternate `json:"linkAlternates"`
}

type Microformat struct {
	MicroformatDataRenderer MicroformatDataRenderer `json:"microformatDataRenderer"`
}

type Response struct {
	ResponseContext      ResponseContext      `json:"responseContext"`
	ContinuationContents ContinuationContents `json:"continuationContents"`
	Metadata             Metadata             `json:"metadata"`
	TrackingParams       string               `json:"trackingParams"`
	Microformat          Microformat          `json:"microformat"`
}

type URLEndpoint struct {
	URL string `json:"url"`
}

type Endpoint struct {
	CommandMetadata CommandMetadata `json:"commandMetadata"`
	URLEndpoint     URLEndpoint     `json:"urlEndpoint"`
}

type Info struct {
	St int `json:"st"`
}

type Timing struct {
	Info Info `json:"info"`
}

type InitialDataResponse struct {
	Contents Contents `json:"contents"`
}

type TabRenderer struct {
	Title          string             `json:"title"`
	Selected       bool               `json:"selected"`
	TrackingParams string             `json:"trackingParams"`
	Content        TabRendererContent `json:"content"`
}

type GridRenderer struct {
	Items         []GridContinuationItem `json:"items"`
	Continuations []Continuation         `json:"continuations"`
}

type ItemSectionRendererContent struct {
	GridRenderer GridRenderer `json:"gridRenderer"`
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

type TabRendererContent struct {
	SectionListRenderer SectionListRenderer `json:"sectionListRenderer"`
}

type Tab struct {
	TabRenderer TabRenderer `json:"tabRenderer,omitempty"`
}

type TwoColumnBrowseResultsRenderer struct {
	Tabs []Tab `json:"tabs"`
}

type Contents struct {
	TwoColumnBrowseResultsRenderer TwoColumnBrowseResultsRenderer `json:"twoColumnBrowseResultsRenderer"`
}
