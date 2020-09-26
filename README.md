# pritube

[![test](https://github.com/dqn/pritube/workflows/test/badge.svg)](https://github.com/dqn/pritube/actions)
[![godoc](https://godoc.org/github.com/dqn/pritube?status.svg)](https://pkg.go.dev/github.com/dqn/pritube?tab=overview)

YouTube private API wrapper.

## Installation

```bash
$ go get github.com/dqn/pritube
```

## Usage Example

```go
package main

import (
  "fmt"

  "github.com/dqn/pritube/channel"
  "github.com/dqn/pritube/chvideos"
  "github.com/dqn/pritube/metadata"
  "github.com/dqn/pritube/video"
)

func main() {
  // === Channel ===
  c, err := channel.FetchInfo("CHANNEL_ID")
  if err != nil {
    // Handle error.
  }

  fmt.Println(c.Title)
  fmt.Println(c.Description)
  fmt.Println(c.SubscriberCount)
  fmt.Println(c.ViewCount)

  // === Video ===
  v, err := video.FetchInfo("VIDEO_ID")
  if err != nil {
    // Handle error.
  }

  fmt.Println(v.VideoDetails.Title)
  fmt.Println(v.VideoDetails.ViewCount)
  fmt.Println(v.VideoDetails.IsLiveContent)
  fmt.Println(v.Microformat.PlayerMicroformatRenderer.Description)

  // === Channel videos ===
  videos, err := chvideos.FetchAll("CHANNEL_ID")
  if err != nil {
    // Handle error.
  }

  for _, v := range videos {
    fmt.Println(v.VideoID)
    fmt.Println(v.Title)
    fmt.Println(v.ViewCountText.SimpleText)
  }

  // === Metadata ===
  // Only streaming live, archive and primed video are supported.
  client := metadata.New()

  m, err := client.Fetch("CHANNEL_ID")
  if err != nil {
    // Handle error.
  }

  fmt.Println(m.Title)
  fmt.Println(m.IsLive)
  fmt.Println(m.ViewCount)
  fmt.Println(m.LikeCount)
  fmt.Println(m.DislikeCount)
}
```

For more information, check out the [documentation](https://pkg.go.dev/github.com/dqn/pritube?tab=overview).

## License

MIT
