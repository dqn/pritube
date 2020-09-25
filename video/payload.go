package video

import "time"

type PlayerResponse struct {
	ResponseContext   ResponseContext   `json:"responseContext"`
	PlayabilityStatus PlayabilityStatus `json:"playabilityStatus"`
	StreamingData     StreamingData     `json:"streamingData"`
	PlaybackTracking  PlaybackTracking  `json:"playbackTracking"`
	VideoDetails      VideoDetails      `json:"videoDetails"`
	PlayerConfig      PlayerConfig      `json:"playerConfig"`
	Storyboards       Storyboards       `json:"storyboards"`
	Microformat       Microformat       `json:"microformat"`
	TrackingParams    string            `json:"trackingParams"`
	Attestation       Attestation       `json:"attestation"`
	Endscreen         Endscreen         `json:"endscreen"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"`
}

type WebResponseContextExtensionData struct {
	HasDecorated bool `json:"hasDecorated"`
}

type ResponseContext struct {
	ServiceTrackingParams           []ServiceTrackingParam          `json:"serviceTrackingParams"`
	WebResponseContextExtensionData WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type MiniplayerRenderer struct {
	PlaybackMode string `json:"playbackMode"`
}

type Miniplayer struct {
	MiniplayerRenderer MiniplayerRenderer `json:"miniplayerRenderer"`
}

type PlayabilityStatus struct {
	Status          string     `json:"status"`
	PlayableInEmbed bool       `json:"playableInEmbed"`
	ContextParams   string     `json:"contextParams"`
	Miniplayer      Miniplayer `json:"miniplayer"`
}

type Format struct {
	URL              string `json:"url,omitempty"`
	Itag             int    `json:"itag"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength,omitempty"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate,omitempty"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
	SignatureCipher  string `json:"signatureCipher"`
}

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type ColorInfo struct {
	Primaries               string `json:"primaries,omitempty"`
	TransferCharacteristics string `json:"transferCharacteristics,omitempty"`
	MatrixCoefficients      string `json:"matrixCoefficients,omitempty"`
}

type AdaptiveFormat struct {
	Format
	InitRange       Range     `json:"initRange"`
	IndexRange      Range     `json:"indexRange"`
	ColorInfo       ColorInfo `json:"colorInfo,omitempty"`
	HighReplication bool      `json:"highReplication,omitempty"`
}

type StreamingData struct {
	ExpiresInSeconds string           `json:"expiresInSeconds"`
	Formats          []Format         `json:"formats"`
	AdaptiveFormats  []AdaptiveFormat `json:"adaptiveFormats"`
}

type URL struct {
	BaseURL                 string `json:"baseUrl"`
	ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds,omitempty"`
}

type PlaybackTracking struct {
	VideostatsPlaybackURL  URL `json:"videostatsPlaybackUrl"`
	VideostatsDelayplayURL URL `json:"videostatsDelayplayUrl"`
	VideostatsWatchtimeURL URL `json:"videostatsWatchtimeUrl"`
	PtrackingURL           URL `json:"ptrackingUrl"`
	QoeURL                 URL `json:"qoeUrl"`
	SetAwesomeURL          URL `json:"setAwesomeUrl"`
	AtrURL                 URL `json:"atrUrl"`
}

type VideoThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnail struct {
	Thumbnails []VideoThumbnail `json:"thumbnails"`
}

type VideoDetails struct {
	VideoID                string    `json:"videoId"`
	Title                  string    `json:"title"`
	LengthSeconds          string    `json:"lengthSeconds"`
	ChannelID              string    `json:"channelId"`
	IsOwnerViewing         bool      `json:"isOwnerViewing"`
	ShortDescription       string    `json:"shortDescription"`
	IsCrawlable            bool      `json:"isCrawlable"`
	Thumbnail              Thumbnail `json:"thumbnail"`
	AverageRating          float64   `json:"averageRating"`
	AllowRatings           bool      `json:"allowRatings"`
	ViewCount              string    `json:"viewCount"`
	Author                 string    `json:"author"`
	IsLowLatencyLiveStream bool      `json:"isLowLatencyLiveStream"`
	IsPrivate              bool      `json:"isPrivate"`
	IsUnpluggedCorpus      bool      `json:"isUnpluggedCorpus"`
	LatencyClass           string    `json:"latencyClass"`
	IsLiveContent          bool      `json:"isLiveContent"`
}

type AudioConfig struct {
	LoudnessDb              float64 `json:"loudnessDb"`
	PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
	EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
}

type StreamSelectionConfig struct {
	MaxBitrate string `json:"maxBitrate"`
}

type DynamicReadaheadConfig struct {
	MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
	MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
	ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
}

type MediaCommonConfig struct {
	DynamicReadaheadConfig DynamicReadaheadConfig `json:"dynamicReadaheadConfig"`
}

type WebPlayerShareEntityServiceEndpoint struct {
	SerializedShareEntity string `json:"serializedShareEntity"`
}

type WebCommandMetadata struct {
	URL      string `json:"url"`
	SendPost bool   `json:"sendPost"`
	APIURL   string `json:"apiUrl"`
}

type CommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}

type GetSharePanelCommand struct {
	ClickTrackingParams                 string                              `json:"clickTrackingParams"`
	CommandMetadata                     CommandMetadata                     `json:"commandMetadata"`
	WebPlayerShareEntityServiceEndpoint WebPlayerShareEntityServiceEndpoint `json:"webPlayerShareEntityServiceEndpoint"`
}

type SubscribeEndpoint struct {
	ChannelIds []string `json:"channelIds"`
	Params     string   `json:"params"`
}

type SubscribeCommand struct {
	ClickTrackingParams string            `json:"clickTrackingParams"`
	SubscribeEndpoint   SubscribeEndpoint `json:"subscribeEndpoint"`
}

type UnsubscribeCommand struct {
	ClickTrackingParams string            `json:"clickTrackingParams"`
	UnsubscribeEndpoint SubscribeEndpoint `json:"unsubscribeEndpoint"`
}

type Action struct {
	AddedVideoID   string `json:"addedVideoId"`
	RemovedVideoID string `json:"removedVideoId"`
	Action         string `json:"action"`
}

type PlaylistEditEndpoint struct {
	PlaylistID string   `json:"playlistId"`
	Actions    []Action `json:"actions"`
}

type WatchLaterCommand struct {
	ClickTrackingParams  string               `json:"clickTrackingParams"`
	PlaylistEditEndpoint PlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type WebPlayerActionsPorting struct {
	GetSharePanelCommand        GetSharePanelCommand `json:"getSharePanelCommand"`
	SubscribeCommand            SubscribeCommand     `json:"subscribeCommand"`
	UnsubscribeCommand          UnsubscribeCommand   `json:"unsubscribeCommand"`
	AddToWatchLaterCommand      WatchLaterCommand    `json:"addToWatchLaterCommand"`
	RemoveFromWatchLaterCommand WatchLaterCommand    `json:"removeFromWatchLaterCommand"`
}

type WebPlayerConfig struct {
	WebPlayerActionsPorting WebPlayerActionsPorting `json:"webPlayerActionsPorting"`
}

type PlayerConfig struct {
	AudioConfig           AudioConfig           `json:"audioConfig"`
	StreamSelectionConfig StreamSelectionConfig `json:"streamSelectionConfig"`
	MediaCommonConfig     MediaCommonConfig     `json:"mediaCommonConfig"`
	WebPlayerConfig       WebPlayerConfig       `json:"webPlayerConfig"`
}

type PlayerStoryboardSpecRenderer struct {
	Spec string `json:"spec"`
}

type Storyboards struct {
	PlayerStoryboardSpecRenderer PlayerStoryboardSpecRenderer `json:"playerStoryboardSpecRenderer"`
}

type Embed struct {
	IframeURL      string `json:"iframeUrl"`
	FlashURL       string `json:"flashUrl"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	FlashSecureURL string `json:"flashSecureUrl"`
}

type Text struct {
	SimpleText string `json:"simpleText"`
}

type LiveBroadcastDetails struct {
	IsLiveNow      bool      `json:"isLiveNow"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}

type PlayerMicroformatRenderer struct {
	Thumbnail            Thumbnail            `json:"thumbnail"`
	Embed                Embed                `json:"embed"`
	Title                Text                 `json:"title"`
	Description          Text                 `json:"description"`
	LengthSeconds        string               `json:"lengthSeconds"`
	OwnerProfileURL      string               `json:"ownerProfileUrl"`
	ExternalChannelID    string               `json:"externalChannelId"`
	AvailableCountries   []string             `json:"availableCountries"`
	IsUnlisted           bool                 `json:"isUnlisted"`
	HasYpcMetadata       bool                 `json:"hasYpcMetadata"`
	ViewCount            string               `json:"viewCount"`
	Category             string               `json:"category"`
	PublishDate          string               `json:"publishDate"`
	OwnerChannelName     string               `json:"ownerChannelName"`
	LiveBroadcastDetails LiveBroadcastDetails `json:"liveBroadcastDetails"`
	UploadDate           string               `json:"uploadDate"`
}

type Microformat struct {
	PlayerMicroformatRenderer PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type BotguardData struct {
	Program        string `json:"program"`
	InterpreterURL string `json:"interpreterUrl"`
}

type PlayerAttestationRenderer struct {
	Challenge    string       `json:"challenge"`
	BotguardData BotguardData `json:"botguardData"`
}

type Attestation struct {
	PlayerAttestationRenderer PlayerAttestationRenderer `json:"playerAttestationRenderer"`
}

type Image struct {
	Thumbnails []VideoThumbnail `json:"thumbnails"`
}

type IconThumbnail struct {
	URL string `json:"url"`
}

type Icon struct {
	Thumbnails []IconThumbnail `json:"thumbnails"`
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

type BrowseEndpoint struct {
	BrowseID string `json:"browseId"`
}

type Endpoint struct {
	ClickTrackingParams string         `json:"clickTrackingParams"`
	BrowseEndpoint      BrowseEndpoint `json:"browseEndpoint,omitempty"`
	WatchEndpoint       WatchEndpoint  `json:"watchEndpoint,omitempty"`
}

type Run struct {
	Text string `json:"text"`
}

type TextRun struct {
	Runs []Run `json:"runs"`
}

type ServiceEndpoint struct {
	ClickTrackingParams string            `json:"clickTrackingParams"`
	SubscribeEndpoint   SubscribeEndpoint `json:"subscribeEndpoint,omitempty"`
	UnsubscribeEndpoint SubscribeEndpoint `json:"unsubscribeEndpoint,omitempty"`
}

type SubscribeButtonRenderer struct {
	ButtonText             TextRun           `json:"buttonText"`
	Subscribed             bool              `json:"subscribed"`
	Enabled                bool              `json:"enabled"`
	Type                   string            `json:"type"`
	ChannelID              string            `json:"channelId"`
	ShowPreferences        bool              `json:"showPreferences"`
	SubscribedButtonText   TextRun           `json:"subscribedButtonText"`
	UnsubscribedButtonText TextRun           `json:"unsubscribedButtonText"`
	TrackingParams         string            `json:"trackingParams"`
	UnsubscribeButtonText  TextRun           `json:"unsubscribeButtonText"`
	ServiceEndpoints       []ServiceEndpoint `json:"serviceEndpoints"`
}

type HovercardButton struct {
	SubscribeButtonRenderer SubscribeButtonRenderer `json:"subscribeButtonRenderer"`
}

type SigninEndpoint struct {
	ClickTrackingParams string `json:"clickTrackingParams"`
}

type EndscreenElementRenderer struct {
	Style                     string          `json:"style"`
	Image                     Image           `json:"image"`
	VideoDuration             Text            `json:"videoDuration,omitempty"`
	Icon                      Icon            `json:"icon,omitempty"`
	Left                      float64         `json:"left"`
	Width                     float64         `json:"width"`
	Top                       float64         `json:"top"`
	AspectRatio               float64         `json:"aspectRatio"`
	StartMs                   string          `json:"startMs"`
	EndMs                     string          `json:"endMs"`
	Title                     Title           `json:"title"`
	Metadata                  Text            `json:"metadata"`
	CallToAction              Text            `json:"callToAction,omitempty"`
	Dismiss                   Text            `json:"dismiss,omitempty"`
	Endpoint                  Endpoint        `json:"endpoint"`
	HovercardButton           HovercardButton `json:"hovercardButton,omitempty"`
	TrackingParams            string          `json:"trackingParams"`
	IsSubscribe               bool            `json:"isSubscribe,omitempty"`
	SigninEndpoint            SigninEndpoint  `json:"signinEndpoint,omitempty"`
	UseClassicSubscribeButton bool            `json:"useClassicSubscribeButton,omitempty"`
	ID                        string          `json:"id"`
}

type WatchEndpoint struct {
	VideoID string `json:"videoId"`
}

type Element struct {
	EndscreenElementRenderer EndscreenElementRenderer `json:"endscreenElementRenderer,omitempty"`
}

type EndscreenRenderer struct {
	Elements       []Element `json:"elements"`
	StartMs        string    `json:"startMs"`
	TrackingParams string    `json:"trackingParams"`
}

type Endscreen struct {
	EndscreenRenderer EndscreenRenderer `json:"endscreenRenderer"`
}
